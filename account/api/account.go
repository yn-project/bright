//nolint:nolintlint,dupl
package account

import (
	"context"

	converter "yun.tea/block/bright/account/pkg/converter/account"
	crud "yun.tea/block/bright/account/pkg/crud/account"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"yun.tea/block/bright/common/logger"
	proto "yun.tea/block/bright/proto/bright/account"

	"github.com/google/uuid"
)

func (s *Server) CreateAccount(ctx context.Context, in *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	var err error
	isRoot := false
	enable := false
	balance := "0"
	in.GetInfo().Balance = &balance
	in.GetInfo().IsRoot = &isRoot
	in.GetInfo().Enable = &enable
	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAccount", "error", err)
		return &proto.CreateAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateAccountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAccounts(ctx context.Context, in *proto.CreateAccountsRequest) (*proto.CreateAccountsResponse, error) {
	var err error

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateAccounts", "error", "Infos is empty")
		return &proto.CreateAccountsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateAccounts", "error", err)
		return &proto.CreateAccountsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateAccountsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAccount(ctx context.Context, in *proto.UpdateAccountRequest) (*proto.UpdateAccountResponse, error) {
	var err error

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateAccount", "ID", in.GetInfo().GetID(), "error", err)
		return &proto.UpdateAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateAccount", "ID", in.GetInfo().GetID(), "error", err)
		return &proto.UpdateAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.UpdateAccountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) UpdateAccounts(ctx context.Context, in *proto.UpdateAccountsRequest) (*proto.UpdateAccountsResponse, error) {
	failedInfos := []*proto.FailedInfo{}
	for _, v := range in.GetInfos() {
		if _, err := uuid.Parse(v.GetID()); err != nil {
			failedInfos = append(failedInfos, &proto.FailedInfo{
				ID:  *v.ID,
				MSG: err.Error(),
			})
			continue
		}

		_, err := crud.Update(ctx, v)
		if err != nil {
			failedInfos = append(failedInfos, &proto.FailedInfo{
				ID:  *v.ID,
				MSG: err.Error(),
			})
		}
	}

	return &proto.UpdateAccountsResponse{
		Infos: failedInfos,
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
