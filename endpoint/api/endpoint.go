//nolint:nolintlint,dupl
package endpoint

import (
	"context"

	converter "yun.tea/block/bright/endpoint/pkg/converter/endpoint"
	crud "yun.tea/block/bright/endpoint/pkg/crud/endpoint"
	"yun.tea/block/bright/endpoint/pkg/mgr"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/proto/bright/basetype"
	proto "yun.tea/block/bright/proto/bright/endpoint"

	"github.com/google/uuid"
)

func (s *Server) CreateEndpoint(ctx context.Context, in *proto.CreateEndpointRequest) (*proto.CreateEndpointResponse, error) {
	if in.Info == nil {
		return &proto.CreateEndpointResponse{}, status.Error(codes.Internal, "not allow nil params")
	}
	if *in.Info.RPS > 10000 || *in.Info.RPS < 1 {
		return &proto.CreateEndpointResponse{}, status.Error(codes.Internal, "wrong rps range,allow [1,10000]")
	}
	var err error
	in.GetInfo().State = basetype.EndpointState_EndpointDefault.Enum()
	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateEndpoint", "error", err)
		return &proto.CreateEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err = mgr.CheckAndUpdateEndpoint(ctx, info)
	if err != nil {
		logger.Sugar().Warnw("CreateEndpoint", "warning", err)
	}

	logger.Sugar().Infof("success to create endpoint,name: %v,address: %v", info.Name, info.Address)
	return &proto.CreateEndpointResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetEndpoint(ctx context.Context, in *proto.GetEndpointRequest) (*proto.GetEndpointResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetEndpoint", "ID", in.GetID(), "error", err)
		return &proto.GetEndpointResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetEndpoint", "ID", in.GetID(), "error", err)
		return &proto.GetEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to get endpoint,name: %v,address: %v", info.Name, info.Address)
	return &proto.GetEndpointResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetEndpoints(ctx context.Context, in *proto.GetEndpointsRequest) (*proto.GetEndpointsResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetEndpoints", "error", err)
		return &proto.GetEndpointsResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to get endpoints,offset-limit: %v-%v,total: %v", in.GetOffset(), in.GetLimit(), total)
	return &proto.GetEndpointsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) DeleteEndpoint(ctx context.Context, in *proto.DeleteEndpointRequest) (*proto.DeleteEndpointResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteEndpoint", "ID", in.GetID(), "error", err)
		return &proto.DeleteEndpointResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteEndpoint", "ID", in.GetID(), "error", err)
		return &proto.DeleteEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	logger.Sugar().Infof("success to delete endpoint,name: %v,address: %v", info.Name, info.Address)
	return &proto.DeleteEndpointResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
