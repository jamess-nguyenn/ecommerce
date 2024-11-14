package main

import (
	"ecommerce/database/connection"
	"ecommerce/database/factories"
	"ecommerce/helpers"
	mysqlmodels "ecommerce/models/mysql"
	mysqlrepositories "ecommerce/repositories/mysql"
	"fmt"
	"sync"
	"time"
)

func main() {
	db, err := connection.ConnectMysql()
	if err != nil {
		fmt.Printf("Could not connect to the database: %v\n", err)
		return
	}

	defer func() {
		if err = db.Close(); err != nil {
			fmt.Printf("Error closing the database connections: %v\n", err)
		}
	}()

	//recordNumber := 10
	recordNumber := 20
	//recordNumber := 200000
	records := factories.SeedProduct(recordNumber)

	beginningTime := helpers.GetTimeNow()

	batchSize := 2
	workerNumber := 3

	//batchSize := 500
	//workerNumber := 5

	var wg sync.WaitGroup
	jobs := make(chan []*mysqlmodels.Product, workerNumber)

	repo := mysqlrepositories.NewProductRepository(db)

	// beginning ==============================
	//if err = repo.CreateMany(records); err != nil {
	//	fmt.Printf("Error inserting records: %+v\n", err)
	//} else {
	//	fmt.Printf("Inserted records successfully\n")
	//}
	//return
	// ending ==============================

	// start the worker pool
	for w := 1; w <= workerNumber; w++ {
		go func(worker int) {
			for batch := range jobs {
				mysqlrepositories.BatchInsert(&wg, repo, batch, worker)
			}
		}(w)
	}

	// send batches to the worker pool
	actualRecordNumber := len(records)
	for i := 0; i < actualRecordNumber; i += batchSize {
		wg.Add(1)

		jobs <- records[i:mysqlrepositories.CalculateEndingIndex(i, batchSize, actualRecordNumber)]
	}

	close(jobs)
	wg.Wait()

	endingTime := helpers.GetTimeNow()

	duration := endingTime.Sub(beginningTime)

	formatingLayout := time.DateTime + ".000000"
	fmt.Printf("Duration: %v\n%s - Beginning datetime\n%s - Ending datetime\n",
		duration, beginningTime.Format(formatingLayout), endingTime.Format(formatingLayout),
	)

}
