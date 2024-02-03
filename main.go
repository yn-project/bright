package main

import (
	"fmt"

	"yun.tea/block/bright/common/utils"
	"yun.tea/block/bright/proto/bright/datafin"
)

func main() {
	DataID := "a123"
	DataFin := "0xajiemv"
	fmt.Println(utils.PrettyStruct(datafin.CheckDataFinRequest{
		TopicID: "id",
		DataID:  &DataID,
		DataFin: &DataFin,
	}))
}
