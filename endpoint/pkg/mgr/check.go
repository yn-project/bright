package mgr

import (
	"context"
	"fmt"

	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
)

const (
	RightChainID = 16
)

func CheckStateAndChainID(ctx context.Context, url string) error {
	cli, err := ethclient.Dial(url)
	if err != nil {
		return err
	}

	chainID, err := cli.ChainID(ctx)
	if err != nil {
		return err
	}

	if chainID.Int64() != RightChainID {
		return fmt.Errorf("wrong chainid: %v , want: %v", chainID.String(), RightChainID)
	}
	return nil
}
