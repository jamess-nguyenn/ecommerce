package handlers

import (
	"ecommerce/helpers"
	"net/http"
	"strconv"
)

func GetId(r *http.Request) int {
	params := helpers.GetParams(r)
	id, _ := strconv.Atoi(params["id"])

	return id
}
