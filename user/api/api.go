package user

import (
	"google.golang.org/grpc"
	proto "yun.tea/block/bright/proto/bright/user"
)

type Server struct {
	proto.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	proto.RegisterManagerServer(server, &Server{})
}