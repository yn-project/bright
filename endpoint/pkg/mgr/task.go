package mgr

import (
	"context"
	"fmt"
	"time"
)

const (
	RefreshTime = time.Second * 5
)

func Maintain(ctx context.Context) {
	for {
		select {
		case <-time.NewTicker(RefreshTime).C:
			fmt.Println("sss")
		case <-ctx.Done():
		}
	}
}
