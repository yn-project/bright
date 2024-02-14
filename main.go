package main

import (
	"context"
	"fmt"

	"yun.tea/block/bright/common/logger"
	proto "yun.tea/block/bright/proto/bright/contract"

	contract "yun.tea/block/bright/contract/api"
	"yun.tea/block/bright/contract/pkg/solcode"
)

func main() {
	logger.Init(logger.DebugLevel, "./a.log")
	cs := &contract.Server{}
	resp, _ := cs.CompileContractCode(context.Background(), &proto.CompileContractCodeRequest{
		Code:            solcode.SOL_CODE,
		Pkg:             solcode.PKG,
		AimContractName: solcode.SOL_CONTRACT,
	})
	fmt.Println(resp.Info.API)
}
