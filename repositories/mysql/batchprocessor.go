package mysql

import (
	"fmt"
	"sync"
)

func BatchInsert[model any](wg *sync.WaitGroup, repo BatchCreator[model], records []*model, worker int) {
	defer wg.Done()

	if err := repo.CreateMany(records); err != nil {
		fmt.Printf("Worker %d: Error inserting records: %+v\n", worker, err)
	} else {
		fmt.Printf("Worker %d: Inserted records successfully\n", worker)
	}
}

func BatchUpdate[model any](wg *sync.WaitGroup, repo BatchUpdater[model], records []*model, worker int) {
	defer wg.Done()

	if err := repo.UpdateMany(records); err != nil {
		fmt.Printf("Worker %d: Error updating records: %+v\n", worker, err)
	} else {
		fmt.Printf("Worker %d: Updated records successfully\n", worker)
	}
}

func CalculateEndingIndex(startingIndex, batchSize, recordNumber int) int {
	endingIndex := startingIndex + batchSize
	if endingIndex > recordNumber {
		endingIndex = recordNumber
	}
	return endingIndex
}
