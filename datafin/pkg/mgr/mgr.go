package mgr

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/Vigo-Tea/go-ethereum-ant/accounts/abi/bind"
	"github.com/Vigo-Tea/go-ethereum-ant/core/types"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	"yun.tea/block/bright/account/pkg/mgr"
	data_fin "yun.tea/block/bright/common/chains/eth/datafin"
	"yun.tea/block/bright/common/cruder"
	"yun.tea/block/bright/common/ctredis"
	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/common/utils"
	topicclient "yun.tea/block/bright/datafin/pkg/client/topic"
	converter "yun.tea/block/bright/datafin/pkg/converter/datafin"
	"yun.tea/block/bright/datafin/pkg/crud/datafin"
	"yun.tea/block/bright/datafin/pkg/db"
	"yun.tea/block/bright/datafin/pkg/db/ent"
	"yun.tea/block/bright/proto/bright"
	proto "yun.tea/block/bright/proto/bright/datafin"
	"yun.tea/block/bright/proto/bright/topic"
)

const (
	maxTxPackNum         = 1000
	maxTxPackTimeout     = time.Second * 30
	maxTaskInterval      = time.Second * 30
	taskGapTime          = time.Second
	minRetryInterval     = time.Second * 10
	maxTopicNum          = 1000
	maxRetries           = 5
	onQueueLockKey       = "lock-state-on_queue"
	onProccessingLockKey = "lock-state-proccesing"
	lockTimeout          = time.Second * 10
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

func PutDataFinInfos(ctx context.Context, topicID string, infos []*proto.DataFinInfo) error {
	state := proto.DataFinState_DataFinStateOnQueue
	for _, info := range infos {
		_, err := datafin.Update(ctx, &proto.DataFinReq{
			DataFinID: &info.DataFinID,
			DataID:    &info.DataID,
			TopicID:   &topicID,
			DataFin:   &info.DataFin,
			Retries:   &info.Retries,
			State:     &state,
			Remark:    &info.Remark,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func Maintain(ctx context.Context) {
	for i := uint32(0); i <= maxRetries; i++ {
		go retry(ctx, i, minRetryInterval)
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
	updateTimeout := time.After(maxTxPackTimeout)
	items := []*ent.DataFin{}
	fulled := false

	defer func() {
		taskMap[topicID] = false
		logger.Sugar().Infof("end %v", topicID)
	}()
	for {
		select {
		case <-time.NewTicker(time.Second * 2).C:
			locked, err := ctredis.TryPubLock(onProccessingLockKey, lockTimeout)
			if err != nil {
				logger.Sugar().Errorf("failed to lock on redis,err: %v", err)
				continue
			}
			if !locked {
				continue
			}

			infos, _, err := datafin.Rows(ctx, &proto.Conds{
				State: &bright.StringVal{
					Op:    cruder.EQ,
					Value: proto.DataFinState_DataFinStateOnQueue.String(),
				},
				TopicID: &bright.StringVal{
					Op:    cruder.EQ,
					Value: topicID,
				},
			}, 0, maxTxPackNum)
			if err != nil {
				logger.Sugar().Errorf("failed get infos from db,err: %v", err)
				continue
			}

			state := proto.DataFinState_DataFinStateProcessing
			for _, info := range infos {
				id := info.ID.String()
				_, err = datafin.Update(ctx, &proto.DataFinReq{
					DataFinID: &id,
					State:     &state,
				})
				if err != nil {
					logger.Sugar().Errorf("failed to ack datafin id to pulsar,err: %v", err)
					continue
				} else {
					items = append(items, info)
				}
			}
			ctredis.UnPubLock(onProccessingLockKey)
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
		_val, err := utils.FromHexString(item.Datafin)
		if err != nil {
			remark := err.Error()
			id := item.ID.String()
			_, _ = datafin.Update(ctx, &proto.DataFinReq{
				DataFinID: &id,
				State:     proto.DataFinState_DataFinStateFailed.Enum(),
				Remark:    &remark,
			})
			continue
		}
		vals = append(vals, _val.ToBigInt())
		ids = append(ids, item.DataID)
		fmt.Println(_val.ToString())
		fmt.Println(item.DataID)
	}

	var tx *types.Transaction
	var err error
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

	state := proto.DataFinState_DataFinStateSeccess

	var txHash string
	var remark string
	if tx != nil {
		txHash = tx.Hash().Hex()
	}

	if err != nil {
		logger.Sugar().Error(err)
		txTime = 0
		state = proto.DataFinState_DataFinStateDefault
		txHash = ""
		remark = err.Error()
	}

	for _, item := range items {
		item.Retries++
		id := item.ID.String()
		_, err = datafin.Update(ctx, &proto.DataFinReq{
			DataFinID: &id,
			TxTime:    &txTime,
			TxHash:    &txHash,
			Remark:    &remark,
			Retries:   &item.Retries,
			State:     &state,
		})
		if err != nil {
			logger.Sugar().Errorf("failed to update datafin,err: %v", err)
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
		case <-time.NewTicker(minInterval << time.Duration(retries)).C:
			locked := false
			for i := 0; i < 3; i++ {
				locked, _ = ctredis.TryPubLock(onQueueLockKey, lockTimeout)
				if locked {
					break
				}
				time.Sleep(lockTimeout)
			}
			if !locked {
				continue
			}

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
			ctredis.UnPubLock(onQueueLockKey)
		case <-ctx.Done():
			return
		}
	}
}
