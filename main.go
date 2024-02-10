package main

import (
	"context"
	"fmt"

	"yun.tea/block/bright/contract/pkg/db"

	"yun.tea/block/bright/account/pkg/mgr"
)

func main() {
	db.Init()
	fmt.Println(mgr.GetAccountReport(context.Background(), "0xbE9Fdc66cB7c462354E95C99534fC6e0eDFeA0dc"))
}
