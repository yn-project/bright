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
	topicID := "f2f2fs"
	go func() {
		producer, err := dataFinProducer(topicID)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		for i := 0; i < 100; i++ {
			_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
				Key:     fmt.Sprintf("%v-%v", utils.RandomBase58(5), i),
				Payload: []byte("ssdf"),
			})
		}
	}()
	consum := func(name string) {
		consummer, err := dataFinConsummer(topicID, name)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		for {
			select {
			case item := <-consummer.Chan():
				fmt.Printf("%v,%v,%v\n", name, item.Key(), string(item.Payload()))
				item.AckID(item.ID())
				// time.Sleep(time.Second + time.Microsecond*300)
			case <-time.NewTicker(time.Second * 10).C:
				return
			}

		}

	}

	time.Sleep(time.Second)

	go consum("c1")
	time.Sleep(time.Second)
	// go consum("c2")
	time.Sleep(time.Second * 2)
	consum("c3")
	// time.Sleep(time.Second * 33)

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
		Topic: topic,
		// SubscriptionInitialPosition: pulsar.SubscriptionPositionEarliest,
		SubscriptionName: "name",
		// Name:                        name,
		Type: pulsar.Shared,
	})

	return consumer, err
}
