package mgr

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
	"yun.tea/block/bright/common/ctpulsar"
)

const (
	datafinTopic  = "datafin-task-topic"
	consummerName = "datafin-task-topic-dafault-consummer"
)

func dataFinProducer() (pulsar.Producer, error) {
	cli, err := ctpulsar.Client()
	if err != nil {
		return nil, err
	}
	producer, err := cli.CreateProducer(pulsar.ProducerOptions{
		Topic: datafinTopic,
	})
	return producer, err
}

func dataFinConsummer() (pulsar.Consumer, error) {
	cli, err := ctpulsar.Client()
	if err != nil {
		return nil, err
	}

	consumer, err := cli.Subscribe(pulsar.ConsumerOptions{
		Topic:            datafinTopic,
		SubscriptionName: consummerName,
		Type:             pulsar.Shared,
	})

	return consumer, err
}

func DataFinTask(ctx context.Context) {

}
