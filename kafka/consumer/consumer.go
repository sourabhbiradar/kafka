package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var addr = "localhost:9092"

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": addr,
		"group.id":          "mygroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Panic(err)
	}

	defer c.Close()

	// consumer subscribing to "myTopic" topic
	c.SubscribeTopics([]string{"myTopic"}, nil)

	// reading messages from myTopic
	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Mesaage on %s : %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			log.Printf("consumer error : %v (%v)\n", err, msg)
		}
	}
}
