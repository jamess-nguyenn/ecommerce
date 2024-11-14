package jobs

import (
	"context"
	"ecommerce/constants"
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
)

func NewWeeklyReport(weekName string) *asynq.Task {
	payload := map[string]any{"weekName": weekName}

	return NewTask(constants.WeeklyReportQueueName, payload)
}

func EnqueueWeeklyReport(client *asynq.Client, weekName string) error {
	task := NewWeeklyReport(weekName)

	if _, err := client.Enqueue(task, BuildTaskOption()); err != nil {
		return err
	}

	return nil
}

func HandleWeeklyReport(ctx context.Context, task *asynq.Task) error {
	var mapPayload map[string]any
	json.Unmarshal(task.Payload(), &mapPayload)
	weekName := mapPayload["weekName"]

	log.Printf("Week name: %s\n", weekName)

	return nil
}
