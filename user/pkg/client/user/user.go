package user

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"yun.tea/block/bright/common/cruder"
	"yun.tea/block/bright/config"
	proto "yun.tea/block/bright/proto/bright/user"
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

func CreateUser(ctx context.Context, in *proto.CreateUserRequest) (resp *proto.CreateUserResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateUser(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetUser(ctx context.Context, in *proto.GetUserRequest) (resp *proto.GetUserResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetUser(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetUsers(ctx context.Context, in *proto.GetUsersRequest) (resp *proto.GetUsersResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetUsers(ctx, in)
		return resp, err
	})
	return resp, err
}

func DeleteUser(ctx context.Context, in *proto.DeleteUserRequest) (resp *proto.DeleteUserResponse, err error) {
	ret, err := withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.DeleteUser(ctx, in)
		return resp, err
	})
	return ret.(*proto.DeleteUserResponse), err
}
