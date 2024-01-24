//nolint:nolintlint,dupl
package account

import (
	"context"

	converter "yun.tea/block/bright/account/pkg/converter/account"
	crud "yun.tea/block/bright/account/pkg/crud/account"
	"yun.tea/block/bright/account/pkg/mgr"
	"yun.tea/block/bright/account/pkg/sign"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"yun.tea/block/bright/common/cruder"
	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/proto/bright"
	proto "yun.tea/block/bright/proto/bright/account"

	"github.com/google/uuid"
)

func (s *Server) CreateAccount(ctx context.Context, in *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	priKey, pubKey, err := sign.GenAccount()
	if err != nil {
		logger.Sugar().Errorw("CreateAccount", "error", err)
		return &proto.CreateAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	balance, isRoot, enable, err := mgr.CheckStateAndBalance(ctx, pubKey)
	if err != nil {
		logger.Sugar().Warnw("CreateAccount", "error", err)
	}

	info := &proto.AccountReq{
		Address: &pubKey,
		PriKey:  &priKey,
		Balance: &balance,
		IsRoot:  &isRoot,
		Enable:  &enable,
		Remark:  &in.Remark,
	}

	crudInfo, err := crud.Create(ctx, info)
	if err != nil {
		logger.Sugar().Errorw("CreateAccount", "error", err)
		return &proto.CreateAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateAccountResponse{
		Info: converter.Ent2Grpc(crudInfo),
	}, nil
}

func (s *Server) ImportAccount(ctx context.Context, in *proto.ImportAccountRequest) (*proto.ImportAccountResponse, error) {
	pubKey, err := sign.GetPubKey(in.PriKey)
	if err != nil {
		logger.Sugar().Errorw("ImportAccount", "error", err)
		return &proto.ImportAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	isRoot := false
	enable := false
	balance, isRoot, enable, err := mgr.CheckStateAndBalance(ctx, pubKey)
	if err != nil {
		logger.Sugar().Warnw("ImportAccount", "error", err)
	}

	info := &proto.AccountReq{
		Address: &pubKey,
		PriKey:  &in.PriKey,
		Balance: &balance,
		IsRoot:  &isRoot,
		Enable:  &enable,
		Remark:  &in.Remark,
	}
	crudInfo, err := crud.Create(ctx, info)
	if err != nil {
		logger.Sugar().Errorw("ImportAccount", "error", err)
		return &proto.ImportAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.ImportAccountResponse{
		Info: converter.Ent2Grpc(crudInfo),
	}, nil
}

func (s *Server) GetAccount(ctx context.Context, in *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetAccount", "ID", in.GetID(), "error", err)
		return &proto.GetAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetAccount", "ID", in.GetID(), "error", err)
		return &proto.GetAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetAccountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAccountPriKey(ctx context.Context, in *proto.GetAccountPriKeyRequest) (*proto.GetAccountPriKeyResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetAccount", "ID", in.GetID(), "error", err)
		return &proto.GetAccountPriKeyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetAccount", "ID", in.GetID(), "error", err)
		return &proto.GetAccountPriKeyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetAccountPriKeyResponse{
		PriKey: info.PriKey,
	}, nil
}

func (s *Server) GetAccounts(ctx context.Context, in *proto.GetAccountsRequest) (*proto.GetAccountsResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetAccounts", "error", err)
		return &proto.GetAccountsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetAccountsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) SetRootAccount(ctx context.Context, in *proto.SetRootAccountRequest) (*proto.SetRootAccountResponse, error) {
	var err error
	conds := &proto.Conds{
		IsRoot: &bright.BoolVal{
			Op:    cruder.EQ,
			Value: true,
		},
	}
	_, total, err := crud.Rows(ctx, conds, 0, 0)
	if err != nil {
		logger.Sugar().Errorw("SetRootAccount", "error", err)
		return &proto.SetRootAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	if total == 0 {
		// 检查是否有合约且是否在合约内是root
	} else {
		// 请求交换root账户，并更新redis中和数据库中的数据
	}

	isRoot := true
	info, err := crud.Update(ctx, &proto.AccountReq{
		ID:     &in.ID,
		IsRoot: &isRoot,
	})

	if err != nil {
		logger.Sugar().Errorw("SetRootAccount", "error", err)
		return &proto.SetRootAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.SetRootAccountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) DeleteAccount(ctx context.Context, in *proto.DeleteAccountRequest) (*proto.DeleteAccountResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteAccount", "ID", in.GetID(), "error", err)
		return &proto.DeleteAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteAccount", "ID", in.GetID(), "error", err)
		return &proto.DeleteAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.DeleteAccountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
