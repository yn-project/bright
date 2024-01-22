package main

import (
	"fmt"

	"yun.tea/block/bright/endpoint/pkg/mgr"
)

func main() {
	val := []string{"sss", "zzz"}
	fmt.Println(mgr.GetEndpintIntervalMGR().SetEndpoinsList(val))
	val2 := []string{}
	fmt.Println(mgr.GetEndpintIntervalMGR().GetEndpoinsList())
	fmt.Println(val2)
	// fmt.Println(mgr.CheckStateAndBalance(context.Background(), ""))
}
