//nolint:nolintlint,dupl
package datafin

import (
	"context"

	converter "yun.tea/block/bright/datafin/pkg/converter/mqueue"
	crud "yun.tea/block/bright/datafin/pkg/crud/mqueue"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"yun.tea/block/bright/common/logger"
	proto "yun.tea/block/bright/proto/bright/mqueue"

	"github.com/google/uuid"
)

func (s *MqueueServer) CreateMqueue(ctx context.Context, in *proto.CreateMqueueRequest) (*proto.CreateMqueueResponse, error) {
	if in.Info == nil || in.Info.Name == nil {
		return &proto.CreateMqueueResponse{}, status.Error(codes.Internal, "not allow nil params")
	}
	topicName := uuid.NewString()
	in.Info.TopicName = &topicName
	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateMqueue", "error", err)
		return &proto.CreateMqueueResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to create mqueue,name: %v", info.Name)
	return &proto.CreateMqueueResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *MqueueServer) GetMqueue(ctx context.Context, in *proto.GetMqueueRequest) (*proto.GetMqueueResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetMqueue", "ID", in.GetID(), "error", err)
		return &proto.GetMqueueResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetMqueue", "ID", in.GetID(), "error", err)
		return &proto.GetMqueueResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to get mqueue,name: %v", info.Name)
	return &proto.GetMqueueResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *MqueueServer) GetMqueues(ctx context.Context, in *proto.GetMqueuesRequest) (*proto.GetMqueuesResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetMqueues", "error", err)
		return &proto.GetMqueuesResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to get mqueues,offset-limit: %v-%v,total: %v", in.GetOffset(), in.GetLimit(), total)
	return &proto.GetMqueuesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *MqueueServer) DeleteMqueue(ctx context.Context, in *proto.DeleteMqueueRequest) (*proto.DeleteMqueueResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteMqueue", "ID", in.GetID(), "error", err)
		return &proto.DeleteMqueueResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteMqueue", "ID", in.GetID(), "error", err)
		return &proto.DeleteMqueueResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to delete mqueue,name: %v", info.Name)
	return &proto.DeleteMqueueResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
