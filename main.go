package main

import (
	"context"
	"fmt"

	"yun.tea/block/bright/account/pkg/client/account"
	"yun.tea/block/bright/common/utils"
	proto "yun.tea/block/bright/proto/bright/account"
)

func main() {
	resp, err := account.CreateAccount(context.Background(), &proto.CreateAccountRequest{
		Remark: "ssss",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(utils.PrettyStruct(resp))
}
