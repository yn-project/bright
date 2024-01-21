package main

import (
	"context"
	"fmt"

	"yun.tea/block/bright/account/pkg/mgr"
)

func main() {
	fmt.Println(mgr.CheckStateAndBalance(context.Background(), ""))
}
