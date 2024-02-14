//nolint:nolintlint,dupl
package account

import (
	"context"
	"time"

	converter "yun.tea/block/bright/account/pkg/converter/account"
	crud "yun.tea/block/bright/account/pkg/crud/account"
	"yun.tea/block/bright/account/pkg/mgr"
	"yun.tea/block/bright/account/pkg/sign"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	data_fin "yun.tea/block/bright/common/chains/eth/datafin"
	"yun.tea/block/bright/common/logger"
	proto "yun.tea/block/bright/proto/bright/account"
	"yun.tea/block/bright/proto/bright/basetype"

	"github.com/Vigo-Tea/go-ethereum-ant/accounts/abi/bind"
	"github.com/Vigo-Tea/go-ethereum-ant/common"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	"github.com/google/uuid"
)

const (
	LongRequestTimeout = time.Second * 24
)

func (s *Server) CreateAccount(ctx context.Context, in *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	priKey, pubKey, err := sign.GenAccount()
	if err != nil {
		logger.Sugar().Errorw("CreateAccount", "error", err)
		return &proto.CreateAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	accReport, err := mgr.GetAccountReport(ctx, pubKey)
	if err != nil {
		logger.Sugar().Warnw("CreateAccount", "error", err)
	}

	info := &proto.AccountReq{
		Address: &pubKey,
		PriKey:  &priKey,
		Balance: &accReport.Balance,
		IsRoot:  &accReport.IsRoot,
		State:   &accReport.State,
		Nonce:   &accReport.Nonce,
		Remark:  &accReport.Remark,
	}

	crudInfo, err := crud.Create(ctx, info)
	if err != nil {
		logger.Sugar().Errorw("CreateAccount", "error", err)
		return &proto.CreateAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to create account,address: %v", info.Address)
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

	accReport, err := mgr.GetAccountReport(ctx, pubKey)
	if err != nil {
		logger.Sugar().Warnw("ImportAccount", "error", err)
	}

	info := &proto.AccountReq{
		Address: &pubKey,
		PriKey:  &in.PriKey,
		Balance: &accReport.Balance,
		IsRoot:  &accReport.IsRoot,
		State:   &accReport.State,
		Nonce:   &accReport.Nonce,
		Remark:  &accReport.Remark,
	}
	crudInfo, err := crud.Create(ctx, info)
	if err != nil {
		logger.Sugar().Errorw("ImportAccount", "error", err)
		return &proto.ImportAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to import account,address: %v", info.Address)
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

	logger.Sugar().Infof("success to get account,address: %v", info.Address)
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

	logger.Sugar().Infof("success to get account pri key,address: %v", info.Address)
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

	logger.Sugar().Infof("success to get accounts,offset-limit: %v-%v,total: %v", in.GetOffset(), in.GetLimit(), total)
	return &proto.GetAccountsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) SetRootAccount(ctx context.Context, in *proto.SetRootAccountRequest) (*proto.SetRootAccountResponse, error) {
	logger.Sugar().Infow("SetRootAccount", "start to set admin account", in.ID)
	ctx, cancel := context.WithTimeout(ctx, LongRequestTimeout)
	defer cancel()

	var err error
	defer func() {
		if err != nil {
			logger.Sugar().Errorw("SetRootAccount", "Err", err)
		}
		logger.Sugar().Infow("SetRootAccount", "end to set root account", in.ID)
	}()

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("SetRootAccount", "ID", in.GetID(), "error", err)
		return &proto.SetRootAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("SetRootAccount", "ID", in.GetID(), "error", err)
		return &proto.SetRootAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	if info.IsRoot {
		return &proto.SetRootAccountResponse{
			Info: converter.Ent2Grpc(info),
		}, nil
	}

	err = mgr.WithWriteContract(ctx, true, func(ctx context.Context, txOpts *bind.TransactOpts, contract *data_fin.DataFin, cli *ethclient.Client) error {
		_, err = contract.ChangeOwner(txOpts, common.HexToAddress(info.Address))
		return err
	})

	if err != nil {
		return &proto.SetRootAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	if err = mgr.GetAccountMGR().SetRootAccount(&mgr.AccountKey{Pri: info.PriKey, Pub: info.Address}); err != nil {
		return &proto.SetRootAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	go mgr.CheckAllAccountState(context.Background())
	info.State = basetype.AccountState_AccountAvailable.String()
	info.IsRoot = true

	logger.Sugar().Infof("success to set root for account,address: %v", info.Address)
	return &proto.SetRootAccountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) SetAdminAccount(ctx context.Context, in *proto.SetAdminAccountRequest) (*proto.SetAdminAccountResponse, error) {
	logger.Sugar().Infow("SetAdminAccount", "start to set admin account", in.ID)
	ctx, cancel := context.WithTimeout(ctx, LongRequestTimeout)
	defer cancel()

	var err error
	defer func() {
		if err != nil {
			logger.Sugar().Errorw("SetAdminAccount", "Err", err)
		}
		logger.Sugar().Infow("SetAdminAccount", "end to set admin account", in.ID)
	}()

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("SetAdminAccount", "ID", in.GetID(), "error", err)
		return &proto.SetAdminAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("SetAdminAccount", "ID", in.GetID(), "error", err)
		return &proto.SetAdminAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	if info.State == basetype.AccountState_AccountAvailable.String() {
		return &proto.SetAdminAccountResponse{
			Info: converter.Ent2Grpc(info),
		}, nil
	}

	err = mgr.WithWriteContract(ctx, true, func(ctx context.Context, txOpts *bind.TransactOpts, contract *data_fin.DataFin, cli *ethclient.Client) error {
		_, err = contract.AddAdmin(txOpts, common.HexToAddress(info.Address), info.Remark)
		return err
	})

	if err != nil {
		return &proto.SetAdminAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	go mgr.CheckAllAccountState(context.Background())
	info.State = basetype.AccountState_AccountAvailable.String()
	logger.Sugar().Infof("success to set admin for account ,address: %v", info.Address)
	return &proto.SetAdminAccountResponse{
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

	logger.Sugar().Infof("success to delete account,address: %v", info.Address)
	return &proto.DeleteAccountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
