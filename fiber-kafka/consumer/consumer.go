package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {
	topic := "comments"
	worker, err := connectConsumer([]string{"localhost:29987"})

	if err != nil {
		log.Panic(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Consumer started ...")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// signal INTerrupted & signal TERMinated ; if any of these is received it will b sent to sigChan

	msgCount := 0 // received messages count

	doneCh := make(chan struct{}) // to notify when done

	// receive msgs
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)

			case msg := <-consumer.Messages():
				fmt.Printf("Received message : %s , Topic : %s , Count : %d", string(msg.Value), string(msg.Topic), msgCount)

			case <-sigChan:
				fmt.Println("intrruption deteced")

				doneCh <- struct{}{}
			}
		}

	}()
	<-doneCh
	fmt.Println("Processed ", msgCount, " messages")

	if err := worker.Close(); err != nil {
		panic(err)
	}
}

func connectConsumer(brokerURL []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokerURL, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
