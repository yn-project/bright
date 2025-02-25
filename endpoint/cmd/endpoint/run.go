package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	cli "github.com/urfave/cli/v2"
	"yun.tea/block/bright/config"
	"yun.tea/block/bright/endpoint/pkg/db"
	"yun.tea/block/bright/endpoint/pkg/mgr"

	"yun.tea/block/bright/common/logger"
	api "yun.tea/block/bright/endpoint/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"yun.tea/block/bright/endpoint/pkg/servicename"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"r"},
	Usage:   fmt.Sprintf("Run %v daemon", servicename.ServiceName),
	After: func(c *cli.Context) error {
		return logger.Sync()
	},
	Before: func(ctx *cli.Context) error {
		err := logger.Init(logger.DebugLevel, config.GetConfig().Endpoint.LogFile)
		if err != nil {
			return err
		}
		return db.Init()
	},
	Action: func(c *cli.Context) error {
		go mgr.Maintain(c.Context)
		go runGRPCServer(config.GetConfig().Endpoint.GrpcPort)
		go runHTTPServer(config.GetConfig().Endpoint.HTTPPort, config.GetConfig().Endpoint.GrpcPort)
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

		<-sigchan
		os.Exit(1)
		return nil
	},
}

func runGRPCServer(grpcPort int) {
	endpoint := fmt.Sprintf(":%v", grpcPort)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	api.Register(server)
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runHTTPServer(httpPort, grpcPort int) {
	httpEndpoint := fmt.Sprintf(":%v", httpPort)
	grpcEndpoint := fmt.Sprintf(":%v", grpcPort)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := api.RegisterGateway(mux, grpcEndpoint, opts)
	if err != nil {
		log.Fatalf("Fail to register gRPC gateway service endpoint: %v", err)
	}

	err = http.ListenAndServe(httpEndpoint, mux)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
