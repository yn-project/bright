package datafin

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	datafinproto "yun.tea/block/bright/proto/bright/datafin"
	topicproto "yun.tea/block/bright/proto/bright/topic"
)

type TopicServer struct {
	topicproto.UnimplementedManagerServer
}

type DataFinServer struct {
	datafinproto.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	topicproto.RegisterManagerServer(server, &TopicServer{})
	datafinproto.RegisterManagerServer(server, &DataFinServer{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	err := topicproto.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	if err != nil {
		return err
	}
	return topicproto.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
