package main

import (
	"ecommerce/constants"
	"ecommerce/database/connection"
	"ecommerce/helpers"
	"ecommerce/lib"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	kp, err := connection.ConnectProducer()
	if err != nil {
		helpers.LogFatal(fmt.Sprintf("Could not connect to the producer: %v", err))
	}

	defer func() {
		if err = kp.Close(); err != nil {
			helpers.LogError(fmt.Sprintf("Error closing the producer connection: %v", err))
		}
	}()

	topic := constants.SendingEmailQueueName
	topicPartition := lib.BuildTopicPartition(&topic, kafka.PartitionAny)
	datetime := helpers.GetDatetime()

	const totalMessages = 6
	//deliveryChan := make(chan kafka.Event, totalMessages)

	for i := 1; i <= totalMessages; i++ {
		index := i + 50

		value := connection.SendingEmail{
			Index:   index,
			Email:   "email_" + helpers.GenerateString(8) + "@example.com",
			Name:    "name " + helpers.GenerateLetter(8),
			Title:   "this is the title at " + datetime,
			Content: "it is the content at " + datetime,
		}

		jsonBytes := helpers.MarshalValue(value)
		if jsonBytes == nil {
			continue
		}

		message := lib.BuildMessage(topicPartition, jsonBytes)
		if err = lib.ProduceMessage(kp.Producer, &message, nil); err != nil {
			log.Printf("Failed to send message to %v with Message %d: %v\n", topic, index, err)
		} else {
			log.Printf("Delivered a message on topic %s: value - %+v\n", topic, value)
		}
	}

	// Wait for delivery reports for all messages
	//for i := 0; i < totalMessages; i++ {
	//	event := <-deliveryChan
	//	msg := event.(*kafka.Message)
	//
	//	if msg.TopicPartition.Error != nil {
	//		log.Printf("Delivery failed: %v\n", msg.TopicPartition.Error)
	//	} else {
	//		log.Printf("Message delivered to topic %s [%d] at offset %v\n",
	//			*msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset)
	//	}
	//}
	//
	//close(deliveryChan)

	// Optional but safe fallback
	lib.FlushMessage(kp.Producer, 1*1000)
}
