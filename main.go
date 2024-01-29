package main

import (
	"fmt"

	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/common/utils"
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
	for i := 0; i < 10; i++ {
		fmt.Println(utils.RandomBase58(64))
	}
}
