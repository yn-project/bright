package mgr

import (
	"context"
	"fmt"

	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	"yun.tea/block/bright/common/constant"
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

	if chainID.Int64() != constant.ChainID.Int64() {
		return fmt.Errorf("wrong chainid: %v , want: %v", chainID.String(), constant.ChainID)
	}
	return nil
}
