package connection

import (
	"ecommerce/helpers"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"strings"
)

const (
	groupIDConsumerGroup1 = "consumer-group1"
	groupIDConsumerGroup2 = "consumer-group2"

	autoOffsetResetEarliest = "earliest"
	autoOffsetResetLatest   = "latest"
	autoOffsetResetNone     = "none"
)

type SendingEmail struct {
	Index   int    `json:"index"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type KafkaProducer struct {
	Producer *kafka.Producer
}

type KafkaConsumer struct {
	Consumer *kafka.Consumer
}

func buildBootstrapServer(index int) string {
	bootstrapServer := fmt.Sprintf("%s:%s",
		os.Getenv(fmt.Sprintf("KAFKA%d_HOST", index)),
		os.Getenv(fmt.Sprintf("KAFKA%d_PORT", index)),
	)

	return bootstrapServer
}

func buildBootstrapServers() string {
	var bootstrapServers []string

	kafkaNumber := helpers.ToInt(os.Getenv("KAFKA_NUMBER"))
	for index := 1; index <= kafkaNumber; index++ {
		bootstrapServers = append(bootstrapServers, buildBootstrapServer(index))
	}

	return strings.Join(bootstrapServers, ",")
}

func buildConfigMapProducer() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": buildBootstrapServers(),
	}
}

func buildConfigMapConsumer() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": buildBootstrapServers(),
		"group.id":          groupIDConsumerGroup1,
		"auto.offset.reset": autoOffsetResetEarliest,
	}
}

func openConnectionProducer(configMap *kafka.ConfigMap) (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the producer: %w", err)
	}

	return producer, nil
}

func openConnectionConsumer(configMap *kafka.ConfigMap) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the consumer: %w", err)
	}

	return consumer, nil
}

func closeConnectionProducer(producer *kafka.Producer) error {
	producer.Close()

	return nil
}

func closeConnectionConsumer(consumer *kafka.Consumer) error {
	if err := consumer.Close(); err != nil {
		return fmt.Errorf("failed to close the consumer: %w", err)
	}

	return nil
}

func ConnectProducer() (*KafkaProducer, error) {
	// open a connection to the producer
	producer, err := openConnectionProducer(buildConfigMapProducer())
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{Producer: producer}, nil
}

func ConnectConsumer() (*KafkaConsumer, error) {
	// open a connection to the consumer
	consumer, err := openConnectionConsumer(buildConfigMapConsumer())
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{Consumer: consumer}, nil
}

func (p *KafkaProducer) Close() error {
	// close a connection to the producer
	if err := closeConnectionProducer(p.Producer); err != nil {
		return fmt.Errorf("errors occurred while closing the producer: %v", err)
	}

	return nil
}

func (c *KafkaConsumer) Close() error {
	// close a connection to the consumer
	if err := closeConnectionConsumer(c.Consumer); err != nil {
		return fmt.Errorf("errors occurred while closing the consumer: %v", err)
	}

	return nil
}
