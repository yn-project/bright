//nolint:nolintlint,dupl
package datafin

import (
	"context"
	"fmt"

	converter "yun.tea/block/bright/datafin/pkg/converter/topic"
	crud "yun.tea/block/bright/datafin/pkg/crud/topic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	accountmgr "yun.tea/block/bright/account/pkg/mgr"
	data_fin "yun.tea/block/bright/common/chains/eth/datafin"
	"yun.tea/block/bright/common/utils"

	"yun.tea/block/bright/common/logger"
	proto "yun.tea/block/bright/proto/bright/topic"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/portto/solana-go-sdk/types"
)

func (s *TopicServer) CreateTopic(ctx context.Context, in *proto.CreateTopicRequest) (*proto.CreateTopicResponse, error) {
	err := accountmgr.WithWriteContract(ctx, false, func(ctx context.Context, acc *accountmgr.AccountKey, contract *data_fin.DataFin) error {
		privateKey, err := crypto.HexToECDSA(acc.Pri)
		if err != nil {
			return err
		}
		var tx *types.Transaction
		topicID := utils.RandomBase58(8)

		switch in.Type {
		case proto.TopicType_IdType:
			topicID = fmt.Sprintf("id-%v", topicID)
			tx, err = contract.CreateIDTopic(privateKey, topicID, in.Name, in.Remark, in.ChangeAble)
		case proto.TopicType_OriginalType:
			topicID = fmt.Sprintf("or-%v", topicID)
			tx, err = contract.CreateTopic(privateKey, topicID, in.Name, in.Remark)
		default:
			return fmt.Errorf("please select a exact topic type")
		}
		return nil
	})
	if err != nil {
		logger.Sugar().Errorw("CreateTopic", "error", err)
		return &proto.CreateTopicResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	logger.Sugar().Infof("success to create topic,name: %v,address: %v", info.Name, info.Address)
	return &proto.CreateTopicResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *TopicServer) GetTopic(ctx context.Context, in *proto.GetTopicRequest) (*proto.GetTopicResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetTopic", "ID", in.GetID(), "error", err)
		return &proto.GetTopicResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetTopic", "ID", in.GetID(), "error", err)
		return &proto.GetTopicResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to get topic,name: %v,address: %v", info.Name, info.Address)
	return &proto.GetTopicResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *TopicServer) GetTopics(ctx context.Context, in *proto.GetTopicsRequest) (*proto.GetTopicsResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetTopics", "error", err)
		return &proto.GetTopicsResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to get topics,offset-limit: %v-%v,total: %v", in.GetOffset(), in.GetLimit(), total)
	return &proto.GetTopicsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *TopicServer) DeleteTopic(ctx context.Context, in *proto.DeleteTopicRequest) (*proto.DeleteTopicResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteTopic", "ID", in.GetID(), "error", err)
		return &proto.DeleteTopicResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteTopic", "ID", in.GetID(), "error", err)
		return &proto.DeleteTopicResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to delete topic,name: %v,address: %v", info.Name, info.Address)
	return &proto.DeleteTopicResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
