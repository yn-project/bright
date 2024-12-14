package main

import (
	"fmt"
	"log"
	"os"

	"net/http"
	_ "net/http/pprof"

	banner "github.com/common-nighthawk/go-figure"
	cli "github.com/urfave/cli/v2"
	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/common/version"
	"yun.tea/block/bright/endpoint/pkg/servicename"
)

const (
	serviceName = "Endpoint"
	usageText   = "Endpoint Service"
)

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:6063", nil)
	}()
	commands := cli.Commands{runCmd}

	description := fmt.Sprintf(
		"%v service cli\nFor help on any individual command run <%v COMMAND -h>\n",
		serviceName,
		serviceName,
	)
	banner.NewColorFigure(serviceName, "", "green", true).Print()
	vesion, err := version.GetVersion()
	if err != nil {
		log.Fatalf("fail to get version, %v", err)
	}

	app := &cli.App{
		Name:        serviceName,
		Version:     vesion,
		Description: description,
		Usage:       usageText,
		Commands:    commands,
	}

	if err != nil {
		logger.Sugar().Errorf("fail to create %v: %v", servicename.ServiceName, err)
		return
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatalf("fail to run %v: %v", serviceName, err)
	}
}
