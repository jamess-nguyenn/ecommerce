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
}
