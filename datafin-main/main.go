package main

import (
	"context"

	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/datafin/pkg/mgr"
)

func main() {
	logger.Init(logger.DebugLevel, "./a.log")
	mgr.ParseFileTask(context.Background())
}
