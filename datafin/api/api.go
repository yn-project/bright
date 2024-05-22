package datafin

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	datafinproto "yun.tea/block/bright/proto/bright/datafin"
	filerecordproto "yun.tea/block/bright/proto/bright/filerecord"
	mqueueproto "yun.tea/block/bright/proto/bright/mqueue"
	overviewproto "yun.tea/block/bright/proto/bright/overview"
	topicproto "yun.tea/block/bright/proto/bright/topic"
)

type TopicServer struct {
	topicproto.UnimplementedManagerServer
}

type DataFinServer struct {
	datafinproto.UnimplementedManagerServer
}

type FileRecordServer struct {
	filerecordproto.UnimplementedManagerServer
}

type MqueueServer struct {
	mqueueproto.UnimplementedManagerServer
}

type OverviewServer struct {
	overviewproto.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	topicproto.RegisterManagerServer(server, &TopicServer{})
	datafinproto.RegisterManagerServer(server, &DataFinServer{})
	filerecordproto.RegisterManagerServer(server, &FileRecordServer{})
	mqueueproto.RegisterManagerServer(server, &MqueueServer{})
	overviewproto.RegisterManagerServer(server, &OverviewServer{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	err := topicproto.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	if err != nil {
		return err
	}
	err = filerecordproto.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	if err != nil {
		return err
	}
	err = datafinproto.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	if err != nil {
		return err
	}
	err = mqueueproto.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	if err != nil {
		return err
	}
	err = overviewproto.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	if err != nil {
		return err
	}
	return nil
}
