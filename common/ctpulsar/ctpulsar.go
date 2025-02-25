package ctpulsar

import (
	"fmt"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"yun.tea/block/bright/config"
)

func Client() (pulsar.Client, error) {
	pulsarConfig := config.GetConfig().Pulsar
	pulsarURL := fmt.Sprintf(
		"pulsar://%v:%v",
		pulsarConfig.Domain,
		pulsarConfig.Port,
	)

	return pulsar.NewClient(pulsar.ClientOptions{
		URL:               pulsarURL,
		OperationTimeout:  time.Duration(pulsarConfig.OperationTimeout) * time.Second,
		ConnectionTimeout: time.Duration(pulsarConfig.ConnectionTimeout) * time.Second,
	})
}
