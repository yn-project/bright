package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Vigo-Tea/go-ethereum-ant/common"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	"yun.tea/block/bright/common/utils"
)

func main() {
	cli, err := ethclient.Dial("https://rest.baas.alipay.com/w3/api/a00e36c5/35N604248fA9u3IfW8BeR2RcQ4ZbMfXb")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	privateKeyStr := "9138747718925d94fb6f3ee732bb387dd779375119ce501e95c478c2ff0eeb2e"

	err = TestTx(cli, privateKeyStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func TestTx(backend *ethclient.Client, priKey string) error {
	tx, ispending, err := backend.TransactionByHash(context.Background(), common.HexToHash("0xef1d804b2ca3251e433ae6c442fe65f326df4812180784c881c0ab7fb18bf7c3"))
	if err != nil {
		return err
	}
	fmt.Println(utils.PrettyStruct(tx))
	fmt.Println(utils.PrettyStruct(ispending))
	return nil
}
