package account

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	proto "yun.tea/block/bright/proto/bright/account"
)

type Server struct {
	proto.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	proto.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, account string, opts []grpc.DialOption) error {
	return proto.RegisterManagerHandlerFromEndpoint(context.Background(), mux, account, opts)
}
