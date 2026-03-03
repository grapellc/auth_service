package messaging

import (
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	"github.com/spf13/viper"
)

var (
	Subscriber *amqp.Subscriber
	Publisher  *amqp.Publisher

	Logger watermill.LoggerAdapter
)

func NewPubSub() {
	Logger = watermill.NewStdLogger(false, false)

	uri := viper.GetString("amqp.uri")
	if uri == "" {
		log.Println("AMQP URI is empty, skipping PubSub initialization")
		return
	}

	amqpConfig := amqp.NewDurablePubSubConfig(uri, amqp.GenerateExchangeNameTopicName)

	var sub *amqp.Subscriber
	var pub *amqp.Publisher
	var err error

	// Retry loop for AMQP connection
	for i := 0; i < 10; i++ {
		sub, err = amqp.NewSubscriber(amqpConfig, Logger)
		if err == nil {
			pub, err = amqp.NewPublisher(amqpConfig, Logger)
			if err == nil {
				break
			}
		}
		log.Printf("Failed to connect to AMQP (attempt %d/10): %v. Retrying in 5s...", i+1, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to AMQP after 10 attempts: %v", err)
	}

	Subscriber = sub
	Publisher = pub
}

func Close() {
	if Subscriber != nil {
		Subscriber.Close()
	}
	if Publisher != nil {
		Publisher.Close()
	}
}
