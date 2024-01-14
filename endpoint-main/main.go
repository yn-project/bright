package main

import (
	"context"
	"time"

	"yun.tea/block/bright/endpoint/pkg/mgr"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	go mgr.Maintain(ctx)
	<-ctx.Done()
}
