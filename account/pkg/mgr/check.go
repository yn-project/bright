package mgr

import (
	"context"

	"github.com/Vigo-Tea/go-ethereum-ant/common"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	endpointmgr "yun.tea/block/bright/endpoint/pkg/mgr"
)

const (
	RightChainID = 16
)

func CheckStateAndBalance(ctx context.Context, address string) (string, error) {
	pubAddr := common.HexToAddress(address)
	balance := "0"
	err := endpointmgr.WithClient(ctx, func(ctx context.Context, cli *ethclient.Client) error {
		_balance, err := cli.BalanceAt(ctx, pubAddr, nil)
		if err == nil {
			balance = _balance.String()
		}
		return err
	})
	return balance, err
}
