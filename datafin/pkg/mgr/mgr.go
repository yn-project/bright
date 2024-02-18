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
	"yun.tea/block/bright/common/ctpulsar"
	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/common/utils"
	topicclient "yun.tea/block/bright/datafin/pkg/client/topic"
	"yun.tea/block/bright/datafin/pkg/crud/datafin"
	"yun.tea/block/bright/datafin/pkg/db"
	proto "yun.tea/block/bright/proto/bright/datafin"
	"yun.tea/block/bright/proto/bright/topic"
)

const (
	maxTxPackNum     = 1000
	maxTxPackTimeout = time.Second * 10
	maxTopicNum      = 1000
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
			fmt.Println(utils.PrettyStruct(item))
			DataFinTask(ctx, item.TopicID, item.Type == topic.TopicType_IDType)
		}(item)
	}
}

func DataFinTask(ctx context.Context, topicID string, isIDTopic bool) error {
	consummerName := fmt.Sprintf("consummer-%v", topicID)
	consummer, err := dataFinConsummer(topicID, consummerName)
	if err != nil {
		return err
	}

	updateTimeout := time.After(maxTxPackTimeout)
	items := []*proto.DataFinInfo{}
	fulled := false

	for {
		select {
		case msg := <-consummer.Chan():
			item := &proto.DataFinInfo{}
			err := json.Unmarshal(msg.Payload(), item)
			fmt.Println(utils.PrettyStruct(item))
			dataFinID := msg.Key()
			state := proto.DataFinState_DataFinStateProcessing
			remark := ""
			if err != nil {
				dataFinID = msg.Key()
				state = proto.DataFinState_DataFinStateFailed
				remark = err.Error()
			}

			_ = updateStateAndAck(ctx, dataFinID, state, remark, &msg)
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
			fmt.Println(err)
			_ = updateStateAndAck(ctx, item.DataFinID, proto.DataFinState_DataFinStateFailed, err.Error(), nil)
			continue
		}
		val := _val.ToBigInt()
		vals = append(vals, val)
		ids = append(ids, item.DataID)
	}

	var tx *types.Transaction
	txTime := uint32(time.Now().Unix())
	err = mgr.WithWriteContract(ctx, false, func(ctx context.Context, txOpts *bind.TransactOpts, contract *data_fin.DataFin, cli *ethclient.Client) error {
		fmt.Println(ids)
		fmt.Println(utils.PrettyStruct(vals))
		if isIDTopic {
			fmt.Println(txOpts.From)
			tx, err = contract.AddIDsItems(txOpts, topicID, txTime, ids, vals)
		} else {
			tx, err = contract.AddItems(txOpts, topicID, txTime, vals)
		}
		if err != nil {
			return err
		}

		time.Sleep(3 * time.Second)
		_, _, err = cli.TransactionByHash(ctx, tx.Hash())
		return err
	})

	for _, item := range items {
		id := item.DataFinID
		state := proto.DataFinState_DataFinStateSeccess

		if err != nil {
			_ = updateStateAndAck(ctx, item.DataFinID, proto.DataFinState_DataFinStateDefault, err.Error(), nil)
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

func updateStateAndAck(ctx context.Context, id string, state proto.DataFinState, remark string, msg *pulsar.ConsumerMessage) error {
	_, err := datafin.Update(ctx, &proto.DataFinReq{
		DataFinID: &id,
		State:     &state,
		Remark:    &remark,
	})
	if err == nil && msg != nil {
		err = msg.AckID(msg.ID())
	}
	if err != nil {
		logger.Sugar().Errorf("failed to update datafin or ack it to pulsar,err: %v", err)
	}
	return err
}
