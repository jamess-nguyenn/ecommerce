package main

import (
	"ecommerce/database/connection"
	"ecommerce/helpers"
	"ecommerce/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Hello")
	return

	db, err := connection.ConnectMysql()
	if err != nil {
		helpers.LogError("Could not connect to the mysql database", map[string]any{"error": err.Error()})
		return
	}

	defer func() {
		if err = db.Close(); err != nil {
			helpers.LogError("Error closing the mysql database connections", map[string]any{"error": err.Error()})
		}
	}()

	router := routes.BuildApiRouter(db)

	if err := http.ListenAndServe(helpers.GetServerAddress(), router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
