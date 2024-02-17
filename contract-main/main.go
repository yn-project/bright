package main

import (
	"fmt"

	"yun.tea/block/bright/common/solc"
	"yun.tea/block/bright/contract/pkg/solcode"
)

func main() {
	a, b, err := solc.GetABIAndBIN(solcode.SOL_FILENAME, solcode.SOL_CODE, solcode.SOL_CONTRACT)
	fmt.Println(err)
	p, err := solc.GenAPICode(a, b, solcode.PKG)
	fmt.Println(p)
	fmt.Println(err)
}
