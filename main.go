package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Vigo-Tea/go-ethereum-ant/accounts/abi/bind"
	"github.com/Vigo-Tea/go-ethereum-ant/common"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	data_fin "yun.tea/block/bright/common/chains/eth/datafin"
	"yun.tea/block/bright/common/logger"
)

func main() {
	logger.Init(logger.DebugLevel, "./a.log")
	logger.Sugar().Info("start")

	cli, err := ethclient.Dial("https://rest.baas.alipay.com/w3/api/a00e36c5/35N604248fA9u3IfW8BeR2RcQ4ZbMfXb")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer cli.Close()

	df, err := data_fin.NewDataFin(common.HexToAddress("0x32eef15f7dd340e20Be199d3D3476A0BB83209a3"), cli)
	if err != nil {
		fmt.Println(err)
		return
	}

	// from := common.HexToAddress("0xD4cce71928bbb36A07d32B0926eA71fb5F5Aeb1B")
	// from := common.HexToAddress("0xbE9Fdc66cB7c462354E95C99534fC6e0eDFeA0dc")
	// from := common.HexToAddress("0x97A6cE565FC4F12dd1Bb3819487fA2d278DB0eCD")
	from := common.HexToAddress("0xD7C5475046948efB8b17Ec75258eA28B6C77A230")
	// from := common.HexToAddress("0x71cbF588c93aEF9B304c2be0918Bdc9AcBFE3Fc6")

	nonce, err := cli.NonceAt(context.Background(), from, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nonce)
	time.Sleep(time.Second)
	addr, err := df.GetOwner(&bind.CallOpts{
		From: from,

		Pending: true,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(addr)

	logger.Sugar().Info("end")

}
