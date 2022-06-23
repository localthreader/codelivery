package kafka

import (
	"encoding/json"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	route2 "github.com/localthreader/codelivery/application/route"
	"github.com/localthreader/codelivery/infra/kafka"
	"os"
	"time"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewProducer()
	route := route2.Route{}
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportToJson()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}

}
