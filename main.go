package main

import (
	"context"
	"fmt"

	"github.com/Vigo-Tea/go-ethereum-ant/common"
	"yun.tea/block/bright/account/pkg/mgr"
	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/common/utils"
	contractmgr "yun.tea/block/bright/contract/pkg/mgr"
)

func main() {
	logger.Init(logger.DebugLevel, "./a.log")
	contractAddr, err := contractmgr.GetContract()
	if err != nil {
		logger.Sugar().Errorw("Maintain", "Msg", "failed to check state of accounts", "Err", err)
		return
	}
	rootAddr, treeAccounts, err := mgr.GetAllAdmin(context.Background(), contractAddr, common.HexToAddress("0xbE9Fdc66cB7c462354E95C99534fC6e0eDFeA0dc"))
	fmt.Println(err)
	fmt.Println(rootAddr)
	fmt.Println(utils.PrettyStruct(treeAccounts))
}
