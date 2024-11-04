package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"yun.tea/block/bright/common/utils"
	"yun.tea/block/bright/endpoint/pkg/client/endpoint"
	"yun.tea/block/bright/proto/bright/basetype"
	proto "yun.tea/block/bright/proto/bright/endpoint"
)

func main() {
	fmt.Println(time.Now())
	wg := sync.WaitGroup{}
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 20; i++ {
				name := uuid.NewString()
				address := "test address"
				state := basetype.EndpointState_EndpointDefault
				rps := uint32(1)
				remark := "ssss"
				resp, err := endpoint.CreateEndpoint(context.Background(), &proto.CreateEndpointRequest{
					Info: &proto.EndpointReq{
						Name:    &name,
						Address: &address,
						RPS:     &rps,
						State:   &state,
						Remark:  &remark,
					},
				})
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(utils.PrettyStruct(resp))
			}
		}()
	}
	wg.Wait()
	fmt.Println(time.Now())
}
