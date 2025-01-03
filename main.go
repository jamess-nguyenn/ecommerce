package main

import (
	"ecommerce/helpers"
	"ecommerce/routes"
	"fmt"
	"net/http"
)

func main() {
	router := routes.BuildApiRouter()

	if err := http.ListenAndServe(helpers.GetServerAddress(), router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
