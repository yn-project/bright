package datafin

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	topicproto "yun.tea/block/bright/proto/bright/topic"
)

type TopicServer struct {
	topicproto.UnimplementedManagerServer
}

type DataFinServer struct {
	topicproto.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	topicproto.RegisterManagerServer(server, &TopicServer{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return topicproto.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
