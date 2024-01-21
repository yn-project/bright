package main

import (
	"fmt"

	"yun.tea/block/bright/account/pkg/sign"
)

func main() {
	fmt.Println(sign.GenAccount())
	fmt.Println(sign.GetPubKey("aaef9013937aec948a3966cdc88ea11e7a0f773adf19031f6425c4d61d3a23fd"))
}
