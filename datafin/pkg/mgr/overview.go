package mgr

import (
	"context"
	"time"

	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	"yun.tea/block/bright/account/pkg/client/account"
	"yun.tea/block/bright/account/pkg/crud/blocknum"
	"yun.tea/block/bright/account/pkg/crud/txnum"
	"yun.tea/block/bright/common/cruder"
	"yun.tea/block/bright/config"
	"yun.tea/block/bright/datafin/pkg/client/topic"
	"yun.tea/block/bright/endpoint/pkg/client/endpoint"
	"yun.tea/block/bright/endpoint/pkg/mgr"
	"yun.tea/block/bright/proto/bright"
	accountproto "yun.tea/block/bright/proto/bright/account"
	"yun.tea/block/bright/proto/bright/basetype"
	endpointproto "yun.tea/block/bright/proto/bright/endpoint"
	"yun.tea/block/bright/proto/bright/overview"
	topicproto "yun.tea/block/bright/proto/bright/topic"
)

const (
	refreshInterval = time.Minute * 5
)

var overviewData *overview.Overview

func GetOverviewData() *overview.Overview {
	return overviewData
}

func OverviewRun(ctx context.Context) {
	for {
		overviewData = &overview.Overview{
			OverviewAt:   uint32(time.Now().Unix()),
			ChainName:    config.GetConfig().Chain.Name,
			ChainID:      config.GetConfig().Chain.ID,
			ChainExplore: config.GetConfig().Chain.Explore,
			ContractLang: config.GetConfig().Chain.Lang,
		}

		mgr.WithClient(ctx, func(ctx context.Context, cli *ethclient.Client) error {
			height, err := cli.BlockNumber(ctx)
			if err != nil {
				return err
			}

			_, err = blocknum.UpsertLatestHeight(ctx, height)
			return err
		})

		for _, v := range basetype.EndpointState_name {
			resp, err := endpoint.GetEndpoints(ctx, &endpointproto.GetEndpointsRequest{Conds: &endpointproto.Conds{
				State: &bright.StringVal{
					Op:    cruder.EQ,
					Value: v,
				},
			}})
			if err == nil && resp != nil {
				overviewData.EndpointStatesNum[v] = resp.Total
				overviewData.EndpointNum += resp.Total
			} else {
				overviewData.EndpointStatesNum[v] = 0
			}
		}

		for _, v := range basetype.AccountState_name {
			resp, err := account.GetAccounts(ctx, &accountproto.GetAccountsRequest{Conds: &accountproto.Conds{
				State: &bright.StringVal{
					Op:    cruder.EQ,
					Value: v,
				},
			}})
			if err == nil && resp != nil {
				overviewData.AccountStatesNum[v] = resp.Total
				overviewData.AccountNum += resp.Total
			} else {
				overviewData.AccountStatesNum[v] = 0
			}
		}

		resp, err := topic.GetTopics(ctx, &topicproto.GetTopicsRequest{})
		if err == nil && resp != nil {
			overviewData.ContractTopicNum = resp.Total
		}

		nums := 20
		overviewData.TxNums = []*overview.TimeNum{}
		txnums, err := txnum.Rows(ctx, nums)
		if err == nil {
			for i := 0; i < nums; i++ {
				if i >= len(txnums) {
					break
				}
				overviewData.TxNums[i] = &overview.TimeNum{
					TimeAt: txnums[i].TimeAt,
					Num:    uint64(txnums[i].Num),
				}
			}
		}

		overviewData.BlockNums = []*overview.TimeNum{}
		blocknums, err := blocknum.Rows(ctx, nums)
		if err == nil {
			for i := 0; i < nums; i++ {
				if i >= len(blocknums) {
					break
				}
				overviewData.BlockNums[i] = &overview.TimeNum{
					TimeAt: blocknums[i].TimeAt,
					Num:    blocknums[i].Height,
				}
			}
		}

		<-time.NewTicker(refreshInterval).C
	}
}