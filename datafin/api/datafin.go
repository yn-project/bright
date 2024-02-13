//nolint:nolintlint,dupl
package datafin

import (
	"context"

	converter "yun.tea/block/bright/datafin/pkg/converter/datafin"
	crud "yun.tea/block/bright/datafin/pkg/crud/datafin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"yun.tea/block/bright/common/logger"
	proto "yun.tea/block/bright/proto/bright/datafin"
	"yun.tea/block/bright/proto/bright/topic"

	"github.com/google/uuid"
)

func (s *DataFinServer) CreateDataFin(ctx context.Context, in *proto.CreateDataFinRequest) (*proto.CreateDataFinResponse, error) {
	topicServer := &TopicServer{}
	_, err := topicServer.GetTopic(ctx, &topic.GetTopicRequest{TopicID: in.TopicID})
	if err != nil {
		return &proto.CreateDataFinResponse{}, err
	}

	reqList := []*proto.DataFinReq{}
	retries := uint32(0)
	for _, v := range in.Infos {
		reqList = append(reqList, &proto.DataFinReq{
			DataID:  &v.DataID,
			TopicID: &in.TopicID,
			Retries: &retries,
			State:   proto.DataFinState_DataFinStateDefault.Enum(),
		})
	}

	crud.CreateBulk(ctx, reqList)
	logger.Sugar().Infof("success to create datafin,name: %v,address: %v", info.Name, info.Address)
	return &proto.CreateDataFinResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *DataFinServer) GetDataFin(ctx context.Context, in *proto.GetDataFinRequest) (*proto.GetDataFinResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetDataFin", "ID", in.GetID(), "error", err)
		return &proto.GetDataFinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetDataFin", "ID", in.GetID(), "error", err)
		return &proto.GetDataFinResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to get datafin,name: %v,address: %v", info.Name, info.Address)
	return &proto.GetDataFinResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *DataFinServer) GetDataFins(ctx context.Context, in *proto.GetDataFinsRequest) (*proto.GetDataFinsResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetDataFins", "error", err)
		return &proto.GetDataFinsResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to get datafins,offset-limit: %v-%v,total: %v", in.GetOffset(), in.GetLimit(), total)
	return &proto.GetDataFinsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *DataFinServer) DeleteDataFin(ctx context.Context, in *proto.DeleteDataFinRequest) (*proto.DeleteDataFinResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteDataFin", "ID", in.GetID(), "error", err)
		return &proto.DeleteDataFinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteDataFin", "ID", in.GetID(), "error", err)
		return &proto.DeleteDataFinResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to delete datafin,name: %v,address: %v", info.Name, info.Address)
	return &proto.DeleteDataFinResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
