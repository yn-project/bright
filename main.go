package main

import (
	"context"
	"fmt"

	"yun.tea/block/bright/account/pkg/db"
	"yun.tea/block/bright/common/logger"
	contractdb "yun.tea/block/bright/contract/pkg/db"

	"yun.tea/block/bright/account/pkg/mgr"
)

func main() {
	db.Init()
	contractdb.Init()
	logger.Init(logger.DebugLevel, "/var/log/a.log")
	mgr.CheckAllAccountState(context.Background())
	fmt.Println(mgr.GetAccountReport(context.Background(), "0xbE9Fdc66cB7c462354E95C99534fC6e0eDFeA0dc"))
}
