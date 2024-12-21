//nolint:nolintlint,dupl
package datafin

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	accountmgr "yun.tea/block/bright/account/pkg/mgr"
	data_fin "yun.tea/block/bright/common/chains/eth/datafin"
	"yun.tea/block/bright/common/utils"
	contractmgr "yun.tea/block/bright/contract/pkg/mgr"
	topicconverter "yun.tea/block/bright/datafin/pkg/converter/topic"
	crud "yun.tea/block/bright/datafin/pkg/crud/topic"

	"yun.tea/block/bright/common/logger"
	proto "yun.tea/block/bright/proto/bright/topic"

	"github.com/Vigo-Tea/go-ethereum-ant/accounts/abi/bind"
	"github.com/Vigo-Tea/go-ethereum-ant/core/types"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
)

func (s *TopicServer) CreateTopic(ctx context.Context, in *proto.CreateTopicRequest) (*proto.CreateTopicResponse, error) {
	topicID := ""
	contractAddr := ""
	onChain := false
	err := accountmgr.WithWriteContract(ctx, false, func(ctx context.Context, txOpts *bind.TransactOpts, contract *data_fin.DataFin, cli *ethclient.Client) error {
		var tx *types.Transaction
		var err error
		topicID = utils.RandomBase58(8)

		switch in.Type {
		case proto.TopicType_IDType:
			topicID = fmt.Sprintf("id-%v", topicID)
			tx, err = contract.CreateIDTopic(txOpts, topicID, in.Name, in.Remark, in.ChangeAble)
		case proto.TopicType_OriginalType:
			in.ChangeAble = false
			topicID = fmt.Sprintf("or-%v", topicID)
			tx, err = contract.CreateTopic(txOpts, topicID, in.Name, in.Remark)
		default:
			return fmt.Errorf("please select a exact topic type")
		}
		if err != nil {
			return err
		}

		isPending := true
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second * 2)
			_, isPending, err = cli.TransactionByHash(ctx, tx.Hash())
			if isPending || err != nil {
				continue
			}
			break
		}
		if isPending || err != nil {
			return fmt.Errorf("create topic failed,please retry")
		}

		contractAddr = tx.To().Hex()
		onChain = true
		return nil
	})
	if err != nil {
		logger.Sugar().Errorw("CreateTopic", "error", err)
		return &proto.CreateTopicResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Create(ctx, &proto.TopicReq{
		TopicID:    &topicID,
		Name:       &in.Name,
		Type:       &in.Type,
		Contract:   &contractAddr,
		OnChain:    &onChain,
		ChangeAble: &in.ChangeAble,
		Remark:     &in.Remark,
	})
	if err != nil {
		logger.Sugar().Errorw("CreateTopic", "error", err)
		return &proto.CreateTopicResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.Sugar().Infof("success to create topic,name: %v,topicID: %v,contract: %v", info.Name, info.TopicID, info.Contract)
	return &proto.CreateTopicResponse{
		TopicID:    info.TopicID,
		Name:       info.Name,
		Type:       proto.TopicType(proto.TopicType_value[info.Type]),
		ChangeAble: info.ChangeAble,
		OnChain:    info.OnChain,
		Remark:     info.Remark,
		CreatedAt:  info.CreatedAt,
	}, nil
}

func (s *TopicServer) GetTopic(ctx context.Context, in *proto.GetTopicRequest) (*proto.GetTopicResponse, error) {
	contractAddr, err := contractmgr.GetContract(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetTopic", "TopicID", in.TopicID, "error", err)
		return &proto.GetTopicResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := crud.Row(ctx, in.TopicID, contractAddr.Hex())
	if err != nil {
		logger.Sugar().Errorw("GetTopic", "TopicID", in.TopicID, "error", err)
		return &proto.GetTopicResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to get topic,name: %v,topicID: %v,contract: %v", info.Name, info.TopicID, contractAddr)
	return &proto.GetTopicResponse{
		TopicID:    info.TopicID,
		Name:       info.Name,
		Type:       proto.TopicType(proto.TopicType_value[info.Type]),
		ChangeAble: info.ChangeAble,
		OnChain:    info.OnChain,
		Remark:     info.Remark,
		CreatedAt:  info.CreatedAt,
	}, nil
}

func (s *TopicServer) GetTopics(ctx context.Context, in *proto.GetTopicsRequest) (*proto.GetTopicsResponse, error) {
	var err error
	contractAddr, err := contractmgr.GetContract(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetTopics", "error", err)
		return &proto.GetTopicsResponse{}, status.Error(codes.Internal, err.Error())
	}

	rows, total, err := crud.Rows(ctx, int(in.GetOffset()), int(in.GetLimit()), contractAddr.Hex())
	if err != nil {
		logger.Sugar().Errorw("GetTopics", "error", err)
		return &proto.GetTopicsResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to get topics,offset-limit: %v-%v,total: %v", in.GetOffset(), in.GetLimit(), total)
	return &proto.GetTopicsResponse{
		Infos: topicconverter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}
