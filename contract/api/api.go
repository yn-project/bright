package contract

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	proto "yun.tea/block/bright/proto/bright/contract"
)

type Server struct {
	proto.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	proto.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return proto.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
