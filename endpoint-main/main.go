package main

import (
	"context"
	"time"

	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/endpoint/pkg/mgr"
)

func main() {
	logger.Init(logger.DebugLevel, "./a.log")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	go mgr.Maintain(ctx)
	<-ctx.Done()
}
