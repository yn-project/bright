//nolint:nolintlint,dupl
package contract

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/common/solc"
	converter "yun.tea/block/bright/contract/pkg/converter/contract"
	crud "yun.tea/block/bright/contract/pkg/crud/contract"
	"yun.tea/block/bright/contract/pkg/solcode"
	proto "yun.tea/block/bright/proto/bright/contract"
)

func (s *Server) GetContractCode(ctx context.Context, in *proto.GetContractCodeRequest) (*proto.GetContractCodeResponse, error) {
	return &proto.GetContractCodeResponse{
		Info: &proto.ContractCode{
			Data:            solcode.SOL_CODE,
			AimContractName: solcode.SOL_CONTRACT,
			Pkg:             solcode.PKG,
			Remark:          solcode.DESCRIPTION,
		},
	}, nil
}

func (s *Server) CompileContractCode(ctx context.Context, in *proto.CompileContractCodeRequest) (*proto.CompileContractCodeResponse, error) {
	abiCode, binCode, err := solc.GetABIAndBIN(solcode.SOL_FILENAME, in.Code, in.AimContractName)
	if err != nil {
		logger.Sugar().Errorf("failed to complie contract code(%v),err %v", in.AimContractName, err)
		return &proto.CompileContractCodeResponse{}, err
	}

	apiCode, err := solc.GenAPICode(abiCode, binCode, in.Pkg)
	if err != nil {
		logger.Sugar().Errorf("failed to complie contract code(%v),err %v", in.AimContractName, err)
		return &proto.CompileContractCodeResponse{}, err
	}

	logger.Sugar().Infof("success to complie contract code(%v)", in.AimContractName)
	return &proto.CompileContractCodeResponse{
		Info: &proto.ContractGEN{
			ABI: abiCode,
			API: apiCode,
			BIN: binCode,
		},
	}, nil
}
func (s *Server) CreateContractWithAccount(ctx context.Context, in *proto.CreateContractWithAccountRequest) (*proto.CreateContractWithAccountResponse, error) {
	version := solcode.VERSION
	info, err := crud.Create(ctx, &proto.ContractReq{
		Name:    &in.Name,
		Address: &in.Name,
		Version: &version,
		Remark:  &in.Remark,
	})
	if err != nil {
		return &proto.CreateContractWithAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to create contract,name: %v,address: %v", info.Name, info.Address)
	return &proto.CreateContractWithAccountResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetContract(ctx context.Context, in *proto.GetContractRequest) (*proto.GetContractResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetContract", "ID", in.GetID(), "error", err)
		return &proto.GetContractResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetContract", "ID", in.GetID(), "error", err)
		return &proto.GetContractResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to get contract,name: %v,address: %v", info.Name, info.Address)
	return &proto.GetContractResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) DeleteContract(ctx context.Context, in *proto.DeleteContractRequest) (*proto.DeleteContractResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteContract", "ID", in.GetID(), "error", err)
		return &proto.DeleteContractResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteContract", "ID", in.GetID(), "error", err)
		return &proto.DeleteContractResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to delete contract,name: %v,address: %v", info.Name, info.Address)
	return &proto.DeleteContractResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
