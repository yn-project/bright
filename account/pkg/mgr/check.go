package mgr

import (
	"context"

	"github.com/Vigo-Tea/go-ethereum-ant/common"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	contractmgr "yun.tea/block/bright/contract/pkg/mgr"
	endpointmgr "yun.tea/block/bright/endpoint/pkg/mgr"
)

const (
	RightChainID = 16
)

func CheckStateAndBalance(ctx context.Context, address string) (balance string, isRoot, enable bool, err error) {
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
		return balance, false, false, err
	}

	contractAddr, err := contractmgr.GetContract()
	if err != nil {
		return balance, false, false, err
	}

	fromAddr, err := getFromAccount(ctx)
	if err != nil {
		return balance, false, false, err
	}

	rootAcc, treeAccs, err := GetAllEnableAdmin(ctx, contractAddr, fromAddr)
	if err != nil {
		return balance, false, false, err
	}

	if rootAcc == address {
		return balance, true, true, nil
	}

	if _, ok := treeAccs[address]; ok {
		return balance, false, true, nil
	}

	return balance, false, false, nil
}
