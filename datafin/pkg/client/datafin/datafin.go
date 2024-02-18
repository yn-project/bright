//nolint:nolintlint,dupl
package datafin

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"yun.tea/block/bright/common/cruder"
	"yun.tea/block/bright/config"
	proto "yun.tea/block/bright/proto/bright/datafin"
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

func CreateDataFin(ctx context.Context, in *proto.CreateDataFinRequest) (resp *proto.CreateDataFinResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateDataFin(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetDataFins(ctx context.Context, in *proto.GetDataFinsRequest) (resp *proto.GetDataFinsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetDataFins(ctx, in)
		return resp, err
	})
	return resp, err
}
