package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/Vigo-Tea/go-ethereum-ant/accounts/abi/bind"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"

	accountdb "yun.tea/block/bright/account/pkg/db"
	endpointdb "yun.tea/block/bright/account/pkg/db"
	"yun.tea/block/bright/account/pkg/mgr"
	data_fin "yun.tea/block/bright/common/chains/eth/datafin"
	"yun.tea/block/bright/common/utils"
	contractdb "yun.tea/block/bright/contract/pkg/db"
)

func main() {
	contractdb.Init()
	accountdb.Init()
	endpointdb.Init()

	err := mgr.WithWriteContract(context.Background(), false, func(ctx context.Context, txOpts *bind.TransactOpts, contract *data_fin.DataFin, cli *ethclient.Client) error {
		tx, err := contract.AddItems(txOpts, "1", uint32(time.Now().Unix()), []*big.Int{big.NewInt(23)})
		if err != nil {
			return err
		}
		fmt.Println(utils.PrettyStruct(tx))
		time.Sleep(time.Second * 3)

		tx, _, err = cli.TransactionByHash(ctx, tx.Hash())
		if err != nil {
			return err
		}
		fmt.Println(utils.PrettyStruct(tx))
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
