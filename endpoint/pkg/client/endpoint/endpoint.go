//nolint:nolintlint,dupl
package endpoint

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"yun.tea/block/bright/common/cruder"
	"yun.tea/block/bright/config"
	proto "yun.tea/block/bright/proto/bright/endpoint"
)

var timeout = 10 * time.Second

type handlerFunc func(context.Context, proto.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handlerFunc) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v",
			config.GetConfig().Endpoint.Domain,
			config.GetConfig().Endpoint.GrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := proto.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateEndpoint(ctx context.Context, in *proto.CreateEndpointRequest) (resp *proto.CreateEndpointResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateEndpoint(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetEndpoint(ctx context.Context, in *proto.GetEndpointRequest) (resp *proto.GetEndpointResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetEndpoint(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetEndpoints(ctx context.Context, in *proto.GetEndpointsRequest) (resp *proto.GetEndpointsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetEndpoints(ctx, in)
		return resp, err
	})
	return resp, err
}

func DeleteEndpoint(ctx context.Context, in *proto.DeleteEndpointRequest) (resp *proto.DeleteEndpointResponse, err error) {
	ret, err := withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.DeleteEndpoint(ctx, in)
		return resp, err
	})
	return ret.(*proto.DeleteEndpointResponse), err
}
