package main

import (
	"ecommerce/constants"
	"ecommerce/database/connection"
	"ecommerce/helpers"
	"ecommerce/lib"
	"fmt"
	"log"
	"time"
)

func main() {
	kc, err := connection.ConnectConsumer()
	if err != nil {
		helpers.LogFatal(fmt.Sprintf("Could not connect to the consumer: %v", err))
	}

	defer func() {
		if err = kc.Close(); err != nil {
			helpers.LogError(fmt.Sprintf("Error closing the consumer connection: %v", err))
		}
	}()

	topic := constants.SendingEmailQueueName
	lib.SubscribeTopics(kc.Consumer, []string{topic}, nil)

	log.Println("The consumer started. Waiting for messages ...")

	//for {
	//	message, err := lib.ConsumeMessage(kc.Consumer, -1)
	//	if err == nil {
	//		var item connection.SendingEmail
	//
	//		if err = helpers.UnmarshalValue(message.Value, &item); err != nil {
	//			continue
	//		}
	//		log.Printf("Received a message on %v: value - %+v\n", message.TopicPartition, item)
	//	} else {
	//		log.Printf("Consumer error: %v (%v)\n", err, message)
	//	}
	//}

	batchSize := 2                    // Number of messages to process in one batch
	timeout := 100 * time.Millisecond // Polling timeout

	for {
		messages := lib.ConsumeBatchMessage(kc.Consumer, batchSize, timeout)

		// Process the batch
		log.Printf("Processing batch of %d messages ...\n\n", len(messages))
		for _, message := range messages {
			var item connection.SendingEmail

			if err = helpers.UnmarshalValue(message.Value, &item); err != nil {
				continue
			}
			log.Printf("Received a message on %v: value - %+v\n", message.TopicPartition, item)
		}

		// Commit offsets after processing the batch
		lib.CommitOffset(kc.Consumer)
	}

}
