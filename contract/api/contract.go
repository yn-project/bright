//nolint:nolintlint,dupl
package contract

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Vigo-Tea/go-ethereum-ant/accounts/abi/bind"
	"github.com/Vigo-Tea/go-ethereum-ant/crypto"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	acconutcrud "yun.tea/block/bright/account/pkg/crud/account"
	accountmgr "yun.tea/block/bright/account/pkg/mgr"
	data_fin "yun.tea/block/bright/common/chains/eth/datafin"
	"yun.tea/block/bright/common/constant"
	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/common/solc"
	converter "yun.tea/block/bright/contract/pkg/converter/contract"
	crud "yun.tea/block/bright/contract/pkg/crud/contract"
	"yun.tea/block/bright/contract/pkg/solcode"
	endpointmgr "yun.tea/block/bright/endpoint/pkg/mgr"
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
	accID, err := uuid.Parse(in.AccountID)
	if err != nil {
		err = fmt.Errorf("wrong account id,cannot parse it to uuid")
		return &proto.CreateContractWithAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	acc, err := acconutcrud.Row(ctx, accID)
	if err != nil {
		err = fmt.Errorf("failed to get account info,err: %v", err)
		return &proto.CreateContractWithAccountResponse{}, status.Error(codes.Internal, err.Error())
	}
	aMGR := accountmgr.GetAccountMGR()
	contractAddr := ""
	err = endpointmgr.WithClient(ctx, func(ctx context.Context, cli *ethclient.Client) error {
		unlockFunc, err := aMGR.GetAccount(ctx, &accountmgr.AccountKey{
			Pri: acc.PriKey,
			Pub: acc.Address,
		}, 3)
		if err != nil {
			return err
		}
		defer unlockFunc()
		privateKey, err := crypto.HexToECDSA(acc.PriKey)
		if err != nil {
			return err
		}

		keyedTransctor, err := bind.NewKeyedTransactorWithChainID(privateKey, constant.ChainID)
		if err != nil {
			return fmt.Errorf("get eth chainID err: %v", err)
		}

		_, tx, _, err := data_fin.DeployDataFin(keyedTransctor, cli)
		if err != nil {
			return fmt.Errorf("failed to deploy contract, err: %v", err)
		}

		time.Sleep(time.Second * 3)
		recipt, err := cli.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return fmt.Errorf("failed to deploy contract, err: %v", err)
		}
		contractAddr = recipt.ContractAddress.Hex()
		return nil
	})

	if err != nil {
		err = fmt.Errorf("failed to create contract,err: %v", err)
		return &proto.CreateContractWithAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	version := solcode.VERSION
	info, err := crud.Create(ctx, &proto.ContractReq{
		Name:    &in.Name,
		Address: &contractAddr,
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
	info, err := crud.Row(ctx)
	if err != nil && strings.Contains(err.Error(), "ent: contract not found") {
		return &proto.GetContractResponse{
			Info: &proto.Contract{
				ID:        "",
				Name:      "无合约，请创建！",
				Address:   "",
				CreatedAt: uint64(time.Now().Unix()),
				UpdatedAt: uint64(time.Now().Unix()),
			},
		}, nil
	}
	if err != nil {
		logger.Sugar().Errorw("GetContract", "error", err)
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
