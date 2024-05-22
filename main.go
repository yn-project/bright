package main

import (
	"context"
	"fmt"
	"time"

	"yun.tea/block/bright/account/pkg/crud/txnum"
	"yun.tea/block/bright/account/pkg/db"
)

func main() {
	fmt.Println(db.Init())
	rows, err := txnum.Rows(context.Background(), 5)
	fmt.Println(err)
	for _, row := range rows {
		fmt.Println(row)
		fmt.Println(time.Unix(int64(row.TimeAt), 0))
	}
}
