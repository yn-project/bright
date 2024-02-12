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
	"yun.tea/block/bright/account/pkg/db"
	"yun.tea/block/bright/account/pkg/mgr"
	"yun.tea/block/bright/config"
	contractdb "yun.tea/block/bright/contract/pkg/db"

	api "yun.tea/block/bright/account/api"
	"yun.tea/block/bright/common/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"yun.tea/block/bright/account/pkg/servicename"
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
		err := db.Init()
		if err != nil {
			return err
		}
		err = contractdb.Init()
		if err != nil {
			return err
		}
		return logger.Init(logger.DebugLevel, config.GetConfig().Account.LogFile)
	},
	Action: func(c *cli.Context) error {
		go mgr.Maintain(c.Context)
		go runGRPCServer(config.GetConfig().Account.GrpcPort)
		go runHTTPServer(config.GetConfig().Account.HTTPPort, config.GetConfig().Account.GrpcPort)
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

		<-sigchan
		os.Exit(1)
		return nil
	},
}

func runGRPCServer(grpcPort int) {
	account := fmt.Sprintf(":%v", grpcPort)
	lis, err := net.Listen("tcp", account)
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
	httpAccount := fmt.Sprintf(":%v", httpPort)
	grpcAccount := fmt.Sprintf(":%v", grpcPort)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := api.RegisterGateway(mux, grpcAccount, opts)
	if err != nil {
		log.Fatalf("Fail to register gRPC gateway service account: %v", err)
	}

	err = http.ListenAndServe(httpAccount, mux)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
