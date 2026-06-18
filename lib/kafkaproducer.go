package lib

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func BuildTopicPartition(topic *string, partition int32) kafka.TopicPartition {
	//kafka.PartitionAny
	return kafka.TopicPartition{
		Topic:     topic,
		Partition: partition,
	}
}

func BuildMessage(topicPartition kafka.TopicPartition, value []byte) kafka.Message {
	return kafka.Message{
		TopicPartition: topicPartition,
		Value:          value,
	}
}

func ProduceMessage(producer *kafka.Producer, message *kafka.Message, deliveryChan chan kafka.Event) error {
	return producer.Produce(message, deliveryChan)
}

func FlushMessage(producer *kafka.Producer, timeoutMs int) int {
	return producer.Flush(timeoutMs)
}

func DeliverMessage(producer *kafka.Producer) {
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v with message: %v\n", ev.TopicPartition, string(ev.Value))
				} else {
					log.Printf("Delivered message to %v with message: %v\n", ev.TopicPartition, string(ev.Value))
				}
			case kafka.Error:
				log.Printf("Producer error: %v\n", ev)
			default:
				log.Printf("Producer ignored the event: %s\n", ev)
			}
		}
	}()
}
