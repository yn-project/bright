//nolint:nolintlint,dupl
package contract

import (
	"context"

	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/common/solc"
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
