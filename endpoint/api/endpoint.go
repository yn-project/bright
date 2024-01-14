//nolint:nolintlint,dupl
package endpoint

import (
	"context"

	converter "yun.tea/block/bright/endpoint/pkg/converter/endpoint"
	crud "yun.tea/block/bright/endpoint/pkg/crud/endpoint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/proto/bright/basetype"
	proto "yun.tea/block/bright/proto/bright/endpoint"

	"github.com/google/uuid"
)

func (s *Server) CreateEndpoint(ctx context.Context, in *proto.CreateEndpointRequest) (*proto.CreateEndpointResponse, error) {
	var err error
	in.GetInfo().State = basetype.EndpointState_EndpointDefault.Enum()
	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateEndpoint", "error", err)
		return &proto.CreateEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateEndpointResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateEndpoints(ctx context.Context, in *proto.CreateEndpointsRequest) (*proto.CreateEndpointsResponse, error) {
	var err error

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateEndpoints", "error", "Infos is empty")
		return &proto.CreateEndpointsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateEndpoints", "error", err)
		return &proto.CreateEndpointsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateEndpointsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateEndpoint(ctx context.Context, in *proto.UpdateEndpointRequest) (*proto.UpdateEndpointResponse, error) {
	var err error

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateEndpoint", "ID", in.GetInfo().GetID(), "error", err)
		return &proto.UpdateEndpointResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateEndpoint", "ID", in.GetInfo().GetID(), "error", err)
		return &proto.UpdateEndpointResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.UpdateEndpointResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) UpdateEndpoints(ctx context.Context, in *proto.UpdateEndpointsRequest) (*proto.UpdateEndpointsResponse, error) {
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

	return &proto.UpdateEndpointsResponse{
		Infos: failedInfos,
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

	return &proto.DeleteEndpointResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
