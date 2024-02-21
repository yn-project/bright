package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"yun.tea/block/bright/common/ctpulsar"
	"yun.tea/block/bright/common/utils"
)

func main() {
	topicID := "fffs"
	producer, err := dataFinProducer(topicID)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	randkey := utils.RandomBase58(5)
	for i := 0; i < 10; i++ {
		_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
			Key:     fmt.Sprintf("%v-%v", randkey, i),
			Payload: []byte("ssdf"),
		})
	}

	fmt.Println(randkey)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	time.Sleep(time.Minute * 10)

	consummer, err := dataFinConsummer(topicID, "sssss")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	for i := 0; i < 10; i++ {
		item := <-consummer.Chan()
		fmt.Println(item.Key())
		fmt.Println(string(item.Payload()))
	}

}
func dataFinProducer(topicID string) (pulsar.Producer, error) {
	cli, err := ctpulsar.Client()
	if err != nil {
		return nil, err
	}
	producer, err := cli.CreateProducer(pulsar.ProducerOptions{
		Topic: topicID,
		Name:  "sss",
	})
	return producer, err
}

func dataFinConsummer(topic string, name string) (pulsar.Consumer, error) {
	cli, err := ctpulsar.Client()
	if err != nil {
		return nil, err
	}

	consumer, err := cli.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: name,
		Type:             pulsar.Shared,
		RetryEnable:      true,
	})

	return consumer, err
}
