package mgr

import (
	"context"

	"github.com/Vigo-Tea/go-ethereum-ant/common"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	contractmgr "yun.tea/block/bright/contract/pkg/mgr"
	endpointmgr "yun.tea/block/bright/endpoint/pkg/mgr"
	"yun.tea/block/bright/proto/bright/basetype"
)

const (
	RightChainID = 16
)

func CheckStateAndBalance(ctx context.Context, address string) (balance string, isRoot bool, state basetype.AccountState, err error) {
	pubAddr := common.HexToAddress(address)
	balance = "0"

	err = endpointmgr.WithClient(ctx, func(ctx context.Context, cli *ethclient.Client) error {
		_balance, err := cli.BalanceAt(ctx, pubAddr, nil)
		if err == nil {
			balance = _balance.String()
		}
		return err
	})

	if err != nil {
		return balance, false, basetype.AccountState_AccountUnkonwn, err
	}

	contractAddr, err := contractmgr.GetContract()
	if err != nil {
		return balance, false, basetype.AccountState_AccountUnkonwn, err
	}

	fromAddr, err := getFromAccount(ctx)
	if err != nil {
		return balance, false, basetype.AccountState_AccountUnkonwn, err
	}

	rootAcc, treeAccs, err := GetAllEnableAdmin(ctx, contractAddr, fromAddr)
	if err != nil {
		return balance, false, basetype.AccountState_AccountUnkonwn, err
	}

	if rootAcc == address {
		return balance, true, basetype.AccountState_AccountAvailable, nil
	}

	if _, ok := treeAccs[address]; ok {
		return balance, false, basetype.AccountState_AccountAvailable, nil
	}

	return balance, false, basetype.AccountState_AccountAvailable, nil
}
