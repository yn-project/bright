package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	cli "github.com/urfave/cli/v2"
	"yun.tea/block/bright/config"

	"yun.tea/block/bright/common/logger"
	api "yun.tea/block/bright/contract/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"yun.tea/block/bright/contract/pkg/db"
	"yun.tea/block/bright/contract/pkg/servicename"
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
		err := logger.Init(logger.DebugLevel, config.GetConfig().Contract.LogFile)
		if err != nil {
			return err
		}
		return db.Init()
	},
	Action: func(c *cli.Context) error {
		go runGRPCServer(config.GetConfig().Contract.GrpcPort)
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
