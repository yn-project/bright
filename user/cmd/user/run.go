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
	"yun.tea/block/bright/user/pkg/db"

	"yun.tea/block/bright/common/logger"
	api "yun.tea/block/bright/user/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"yun.tea/block/bright/user/pkg/servicename"
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
		err := logger.Init(logger.DebugLevel, config.GetConfig().User.LogFile)
		if err != nil {
			return err
		}
		return db.Init()
	},
	Action: func(c *cli.Context) error {
		go runGRPCServer(config.GetConfig().User.GrpcPort)
		go runHTTPServer(config.GetConfig().User.HTTPPort, config.GetConfig().User.GrpcPort)
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
	httpAddr := fmt.Sprintf(":%v", httpPort)
	grpcAddr := fmt.Sprintf(":%v", grpcPort)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := api.RegisterGateway(mux, grpcAddr, opts)
	if err != nil {
		log.Fatalf("Fail to register gRPC gateway service user: %v", err)
	}

	err = http.ListenAndServe(httpAddr, mux)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
