//nolint:nolintlint,dupl
package mqueue

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"yun.tea/block/bright/common/cruder"
	"yun.tea/block/bright/config"
	proto "yun.tea/block/bright/proto/bright/mqueue"
)

var timeout = 10 * time.Second

type handlerFunc func(context.Context, proto.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handlerFunc) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v",
			config.GetConfig().DataFin.Domain,
			config.GetConfig().DataFin.GrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := proto.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateMqueue(ctx context.Context, in *proto.CreateMqueueRequest) (resp *proto.CreateMqueueResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateMqueue(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetMqueues(ctx context.Context, in *proto.GetMqueuesRequest) (resp *proto.GetMqueuesResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetMqueues(ctx, in)
		return resp, err
	})
	return resp, err
}
