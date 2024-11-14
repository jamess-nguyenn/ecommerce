package jobs

import (
	"ecommerce/database/connection"
	"ecommerce/helpers"
	"encoding/json"
	"github.com/hibiken/asynq"
	"time"
)

// InitializeQueueClient initializes and starts queue client
func InitializeQueueClient() *asynq.Client {
	client := asynq.NewClient(
		connection.BuildRedisClientOptionsAsynq(),
	)

	return client
}

// InitializeQueueServer initializes and starts queue server
func InitializeQueueServer() *asynq.Server {
	server := asynq.NewServer(
		connection.BuildRedisClientOptionsAsynq(),
		asynq.Config{Concurrency: 10},
	)

	return server
}

func NewTask(queueName string, payload map[string]any) *asynq.Task {
	// serialize payload to JSON
	payloadBytes, _ := json.Marshal(payload)

	return asynq.NewTask(queueName, payloadBytes)
}

func BuildTaskOption() asynq.Option {
	return asynq.ProcessAt(helpers.GetTimeNow().Add(3 * time.Second))
}
