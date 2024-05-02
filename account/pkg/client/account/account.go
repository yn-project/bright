//nolint:nolintlint,dupl
package account

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"yun.tea/block/bright/common/cruder"
	"yun.tea/block/bright/config"
	proto "yun.tea/block/bright/proto/bright/account"
)

var timeout = 10 * time.Second

type handlerFunc func(context.Context, proto.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handlerFunc) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v",
			config.GetConfig().Account.Domain,
			config.GetConfig().Account.GrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := proto.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateAccount(ctx context.Context, in *proto.CreateAccountRequest) (resp *proto.CreateAccountResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.CreateAccount(ctx, in)
		return resp, err
	})
	return resp, err
}

func ImportAccount(ctx context.Context, in *proto.ImportAccountRequest) (resp *proto.ImportAccountResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.ImportAccount(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetAccount(ctx context.Context, in *proto.GetAccountRequest) (resp *proto.GetAccountResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetAccount(ctx, in)
		return resp, err
	})
	return resp, err
}

func GetAccounts(ctx context.Context, in *proto.GetAccountsRequest) (resp *proto.GetAccountsResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.GetAccounts(ctx, in)
		return resp, err
	})
	return resp, err
}

func SetRootAccount(ctx context.Context, in *proto.SetRootAccountRequest) (resp *proto.SetRootAccountResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.SetRootAccount(ctx, in)
		return resp, err
	})
	return resp, err
}

func SetAdminAccount(ctx context.Context, in *proto.SetAdminAccountRequest) (resp *proto.SetAdminAccountResponse, err error) {
	_, err = withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.SetAdminAccount(ctx, in)
		return resp, err
	})
	return resp, err
}

func DeleteAccount(ctx context.Context, in *proto.DeleteAccountRequest) (resp *proto.DeleteAccountResponse, err error) {
	ret, err := withCRUD(ctx, func(_ctx context.Context, cli proto.ManagerClient) (cruder.Any, error) {
		resp, err = cli.DeleteAccount(ctx, in)
		return resp, err
	})
	return ret.(*proto.DeleteAccountResponse), err
}
