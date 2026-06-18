package lib

import (
	"ecommerce/helpers"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"time"
)

func SubscribeTopics(consumer *kafka.Consumer, topics []string, rebalanceCb kafka.RebalanceCb) {
	if err := consumer.SubscribeTopics(topics, rebalanceCb); err != nil {
		helpers.LogFatal(fmt.Sprintf("Failed to subscribe to topics: %v", err))
	}
}

func ConsumeMessage(consumer *kafka.Consumer, timeout time.Duration) (*kafka.Message, error) {
	return consumer.ReadMessage(timeout)
}

func ConsumeBatchMessage(consumer *kafka.Consumer, batchSize int, timeout time.Duration) []*kafka.Message {
	messages := make([]*kafka.Message, 0, batchSize)

	for len(messages) < batchSize {
		ev := consumer.Poll(int(timeout.Milliseconds()))
		if ev == nil {
			continue // No message within the timeout
		}

		switch message := ev.(type) {
		case *kafka.Message:
			messages = append(messages, message)
		case kafka.Error:
			log.Printf("Consumer error: %v\n", message)
		default:
			log.Printf("Consumer ignored the event: %s\n", message)
		}
	}

	return messages
}

func CommitOffset(consumer *kafka.Consumer) {
	if _, err := consumer.Commit(); err != nil {
		helpers.LogError(fmt.Sprintf("Failed to commit offset: %v", err))
	}
}
