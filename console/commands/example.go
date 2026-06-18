package commands

import (
	"ecommerce/queues/jobs"
	"log"
)

func exampleCommand() {
	client := jobs.InitializeQueueClient()

	email := "example_user@example.com"
	log.Printf("Sending to email: %s\n", email)
	jobs.EnqueueSendingEmail(client, email)

	//func produceMessage() {
	//	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	//	if err != nil {
	//		log.Fatalf("Failed to create producer: %s", err)
	//	}
	//	defer p.Close()
	//
	//	message := fmt.Sprintf("Cron job executed at %s", time.Now().Format(time.RFC3339))
	//
	//	err = p.Produce(&kafka.Message{
	//		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	//		Value:          []byte(message),
	//	}, nil)
	//	if err != nil {
	//		log.Printf("Failed to produce message: %v", err)
	//	} else {
	//		log.Printf("Produced message: %s", message)
	//	}
	//}

}
