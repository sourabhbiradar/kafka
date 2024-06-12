package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/IBM/sarama"

	"github.com/gofiber/fiber/v2"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func connectProducer(brokerURL []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 3

	conn, err := sarama.NewSyncProducer(brokerURL, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func PuchCommentToQueue(topic string, message []byte) error {
	brokerURL := []string{"localhost:29987"}
	producer, err := connectProducer(brokerURL)
	if err != nil {
		return errors.New("failed to create producer")
	}
	defer producer.Close()

	// topic = "comment"
	// message = cmtInBytes

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%v)/partition(%v)/offset(%v)\n", topic, partition, offset)
	return nil
}

func createComment(c *fiber.Ctx) error {
	cmt := new(Comment) // new instance of Comment

	if err := c.BodyParser(cmt); err != nil { // req body into Commnet struct
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return err
	}
	cmtInBytes, err := json.Marshal(cmt)
	PuchCommentToQueue("comment", cmtInBytes)

	err = c.JSON(&fiber.Map{
		"success": true,
		"message": "comment pushed successfully",
		"comment": cmt,
	})

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating product",
		})
		return err
	}
	return err
}
