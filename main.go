package main

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	kafka2 "github.com/localthreader/codelivery/application/kafka"
	"github.com/localthreader/codelivery/infra/kafka"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading .env file")
	}

}

func main() {
	consumerChannel := make(chan *ckafka.Message)
	consumer := kafka.NewConsumer(consumerChannel)
	go consumer.Consume()

	for msg := range consumerChannel {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}
