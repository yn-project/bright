package main

import (
	"fmt"

	"github.com/Vigo-Tea/go-ethereum-ant/common/hexutil"
	"yun.tea/block/bright/common/utils"
)

func main() {
	fmt.Println(hexutil.Decode("0xd32050b9b7f1b18f7dd8639bba9284b8344b003f8be47daad6d53459622ca68f"))
	rr, err := utils.FromHexString("0xd32050b9b7f1b18f7dd8639bba9284b8344b003f8be47daad6d53459622ca68f")
	fmt.Println(err)
	fmt.Println(rr.ToHexString())
}
