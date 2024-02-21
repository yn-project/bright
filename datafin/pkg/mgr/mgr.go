package mgr

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/Vigo-Tea/go-ethereum-ant/accounts/abi/bind"
	"github.com/Vigo-Tea/go-ethereum-ant/core/types"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	"github.com/apache/pulsar-client-go/pulsar"
	"yun.tea/block/bright/account/pkg/mgr"
	data_fin "yun.tea/block/bright/common/chains/eth/datafin"
	"yun.tea/block/bright/common/cruder"
	"yun.tea/block/bright/common/ctpulsar"
	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/common/utils"
	topicclient "yun.tea/block/bright/datafin/pkg/client/topic"
	converter "yun.tea/block/bright/datafin/pkg/converter/datafin"
	"yun.tea/block/bright/datafin/pkg/crud/datafin"
	"yun.tea/block/bright/datafin/pkg/db"
	"yun.tea/block/bright/proto/bright"
	proto "yun.tea/block/bright/proto/bright/datafin"
	"yun.tea/block/bright/proto/bright/topic"
)

const (
	maxTxPackNum     = 1000
	maxTxPackTimeout = time.Minute
	maxTaskInterval  = time.Minute
	taskGapTime      = time.Second * 8
	maxTopicNum      = 1000
	maxRetries       = 5
)

var (
	taskMap = make(map[string]bool)
)

func init() {
	err := db.Init()
	if err != nil {
		logger.Sugar().Error(err)
	}
}

func dataFinProducer(topicID string) (pulsar.Producer, error) {
	cli, err := ctpulsar.Client()
	if err != nil {
		return nil, err
	}
	producer, err := cli.CreateProducer(pulsar.ProducerOptions{
		Topic: topicID,
		Name:  "datafin-producer",
	})
	return producer, err
}

func dataFinConsummer(topic string, name string) (pulsar.Consumer, error) {
	cli, err := ctpulsar.Client()
	if err != nil {
		return nil, err
	}

	consumer, err := cli.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: name,
		Type:             pulsar.Shared,
		RetryEnable:      true,
	})

	return consumer, err
}

func PutDataFinInfos(ctx context.Context, topicID string, infos []*proto.DataFinInfo) error {
	producer, err := dataFinProducer(topicID)
	if err != nil {
		return err
	}
	defer producer.Close()

	for _, info := range infos {
		payload, err := json.Marshal(info)
		if err != nil {
			return err
		}
		_, err = producer.Send(ctx, &pulsar.ProducerMessage{
			Key:     info.DataFinID,
			Payload: payload,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func Maintain(ctx context.Context) {
	for i := uint32(1); i <= maxRetries; i++ {
		go retry(ctx, i, maxTxPackTimeout)
	}

	for {
		select {
		case <-time.NewTicker(maxTaskInterval).C:
			resp, err := topicclient.GetTopics(ctx, &topic.GetTopicsRequest{
				Offset: 0,
				Limit:  maxTopicNum,
			})
			if err != nil {
				logger.Sugar().Errorf("failed to get topics,err %v", err)
				return
			}
			for _, item := range resp.Infos {
				func(item *topic.TopicInfo) {
					if _, ok := taskMap[item.TopicID]; !ok {
						taskMap[item.TopicID] = false
					}
					if taskMap[item.TopicID] {
						return
					}
					taskMap[item.TopicID] = true
					go dataFinTask(ctx, item.TopicID, item.Type == topic.TopicType_IDType)
				}(item)
				time.Sleep(taskGapTime)
			}
		case <-ctx.Done():
			return
		}
	}
}

func dataFinTask(ctx context.Context, topicID string, isIDTopic bool) error {
	consummerName := fmt.Sprintf("consummer-%v", topicID)
	consummer, err := dataFinConsummer(topicID, consummerName)
	if err != nil {
		return err
	}
	updateTimeout := time.After(maxTxPackTimeout)
	items := []*proto.DataFinInfo{}
	fulled := false

	logger.Sugar().Infof("start %v %v", topicID, time.Now().String())
	defer func() {
		taskMap[topicID] = false
		logger.Sugar().Infof("end %v %v", topicID, time.Now().String())
	}()
	for {
		select {
		case msg := <-consummer.Chan():
			item := &proto.DataFinInfo{}
			err := json.Unmarshal(msg.Payload(), item)
			dataFinID := msg.Key()
			state := proto.DataFinState_DataFinStateProcessing
			remark := ""
			if err != nil {
				dataFinID = msg.Key()
				state = proto.DataFinState_DataFinStateFailed
				remark = err.Error()
			}

			_ = updateStateAndAck(ctx, &proto.DataFinReq{
				DataFinID: &dataFinID,
				State:     &state,
				Remark:    &remark,
			}, &msg)
			if err = msg.AckID(msg.ID()); err != nil {
				logger.Sugar().Errorf("failed to ack datafin id to pulsar,err: %v", err)
			}

			items = append(items, item)
			if len(items) >= maxTxPackNum {
				fulled = true
			}
		case <-updateTimeout:
			fulled = true
		}
		if fulled {
			break
		}
	}

	if len(items) == 0 {
		return nil
	}

	vals := []*big.Int{}
	ids := []string{}
	for _, item := range items {
		_val, err := utils.FromHexString(item.DataFin)
		if err != nil {
			remark := err.Error()
			_ = updateStateAndAck(ctx, &proto.DataFinReq{
				DataFinID: &item.DataFinID,
				State:     proto.DataFinState_DataFinStateFailed.Enum(),
				Remark:    &remark,
			}, nil)
			continue
		}
		val := _val.ToBigInt()
		vals = append(vals, val)
		ids = append(ids, item.DataID)
	}

	var tx *types.Transaction
	txTime := uint32(time.Now().Unix())
	err = mgr.WithWriteContract(ctx, false, func(ctx context.Context, txOpts *bind.TransactOpts, contract *data_fin.DataFin, cli *ethclient.Client) error {
		if isIDTopic {
			tx, err = contract.AddIDsItems(txOpts, topicID, txTime, ids, vals)
		} else {
			tx, err = contract.AddItems(txOpts, topicID, txTime, vals)
		}
		if err != nil {
			return err
		}

		time.Sleep(5 * time.Second)
		_, _, err = cli.TransactionByHash(ctx, tx.Hash())
		return err
	})

	if err != nil {
		logger.Sugar().Error(err)
	}

	for _, item := range items {
		id := item.DataFinID
		state := proto.DataFinState_DataFinStateSeccess

		if err != nil {
			item.Retries++
			remark := err.Error()
			_ = updateStateAndAck(ctx, &proto.DataFinReq{
				DataFinID: &id,
				State:     proto.DataFinState_DataFinStateDefault.Enum(),
				Retries:   &item.Retries,
				Remark:    &remark,
			}, nil)
		} else {
			_, err = datafin.Update(ctx, &proto.DataFinReq{
				DataFinID: &id,
				TxTime:    &txTime,
				TxHash:    &item.TxHash,
				State:     &state,
			})
			if err != nil {
				logger.Sugar().Errorf("failed to update datafin,err: %v", err)
			}
		}
	}

	return nil
}

func retry(ctx context.Context, retries uint32, minInterval time.Duration) {
	conds := &proto.Conds{
		Retries: &bright.Uint32Val{
			Op:    cruder.EQ,
			Value: retries,
		},
		State: &bright.StringVal{
			Op:    cruder.EQ,
			Value: proto.DataFinState_DataFinStateDefault.String(),
		},
	}
	for {
		select {
		case <-time.NewTicker(minInterval << time.Duration(retries-1)).C:
			resp, err := topicclient.GetTopics(ctx, &topic.GetTopicsRequest{
				Offset: 0,
				Limit:  maxTopicNum,
			})
			if err != nil {
				logger.Sugar().Errorf("failed to get topics,err %v", err)
				return
			}

			for _, item := range resp.Infos {
				conds.TopicID = &bright.StringVal{
					Op:    cruder.EQ,
					Value: item.TopicID,
				}

				rows, _, err := datafin.Rows(ctx, conds, 0, maxTxPackNum)
				if err != nil {
					logger.Sugar().Error(err)
				}
				if len(rows) == 0 {
					continue
				}
				if retries >= maxRetries {
					for _, row := range rows {
						id := row.ID.String()
						remark := "exhausted the maximum number of retries"
						datafin.Update(ctx, &proto.DataFinReq{
							DataFinID: &id,
							State:     proto.DataFinState_DataFinStateFailed.Enum(),
							Remark:    &remark,
						})
					}
					continue
				}
				_ = PutDataFinInfos(ctx, item.TopicID, converter.Ent2GrpcMany(rows))
			}
		case <-ctx.Done():
			return
		}
	}
}

func updateStateAndAck(ctx context.Context, in *proto.DataFinReq, msg *pulsar.ConsumerMessage) error {
	_, err := datafin.Update(ctx, in)
	if err == nil && msg != nil {
		err = msg.AckID(msg.ID())
	}
	if err != nil {
		logger.Sugar().Errorf("failed to update datafin or ack it to pulsar,err: %v", err)
	}
	return err
}
