package main

import (
	"context"
	"fmt"
	"time"

	"yun.tea/block/bright/account/pkg/mgr"
)

func main() {
	fmt.Println(time.Now().String())
	fmt.Println(mgr.CheckStateAndBalance(context.Background(), "0x7243176257d634A65Ce737c2cd5FAb6B3f65154B"))
	fmt.Println(time.Now().String())
}
