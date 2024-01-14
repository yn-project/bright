package contract

import (
	"google.golang.org/grpc"
	proto "yun.tea/block/bright/proto/bright/contract"
)

type Server struct {
	proto.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	proto.RegisterManagerServer(server, &Server{})
}
