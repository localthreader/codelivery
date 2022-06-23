package kafka

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

type Consumer struct {
	MessageChannel chan *ckafka.Message
}

func NewConsumer(channel chan *ckafka.Message) *Consumer {
	return &Consumer{
		MessageChannel: channel,
	}
}

func (k *Consumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error consuming kafka message: " + err.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopic")}
	c.SubscribeTopics(topics, nil)
	fmt.Println("Kafka consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)
		fmt.Printf("leu: %s", msg.Value)
		if err == nil {
			k.MessageChannel <- msg
		} else {
			fmt.Println(err.Error())
		}
	}
}
