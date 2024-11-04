package main

import (
	"context"
	"fmt"
	"time"

	"yun.tea/block/bright/account/pkg/client/account"
	"yun.tea/block/bright/common/utils"
	proto "yun.tea/block/bright/proto/bright/account"
)

func main() {
	fmt.Println(time.Now())
	for i := 0; i < 1000; i++ {
		resp, err := account.CreateAccount(context.Background(), &proto.CreateAccountRequest{
			Remark: "ssss",
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(utils.PrettyStruct(resp))
	}
	fmt.Println(time.Now())
}
