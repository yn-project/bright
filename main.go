package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Vigo-Tea/go-ethereum-ant/accounts/abi/bind"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"

	accountdb "yun.tea/block/bright/account/pkg/db"
	endpointdb "yun.tea/block/bright/account/pkg/db"
	"yun.tea/block/bright/account/pkg/mgr"
	data_fin "yun.tea/block/bright/common/chains/eth/datafin"
	contractdb "yun.tea/block/bright/contract/pkg/db"
)

func main() {
	contractdb.Init()
	accountdb.Init()
	endpointdb.Init()

	err := mgr.WithWriteContract(context.Background(), false, func(ctx context.Context, txOpts *bind.TransactOpts, contract *data_fin.DataFin, cli *ethclient.Client) error {
		contract.AddItems(txOpts, "")
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

}
