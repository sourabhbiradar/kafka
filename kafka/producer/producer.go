package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var addr = "localhost:9092"

func main() {
	// creating & configuring producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": addr,
	})

	if err != nil {
		log.Panic(err)
	}

	// close producer
	defer p.Close()

	// topic creation
	topic := "myTopic"
	message := []string{"m1", "m2", "m3", "m4"}

	for _, msg := range message {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,             // into myTopic
				Partition: kafka.PartitionAny, // no specific partition
			},
			Value: []byte(msg),
		}, nil)

	}

	// waiting for messages to be delivered
	p.Flush(15 * 1000)
}
