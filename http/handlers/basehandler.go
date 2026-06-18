package handlers

import (
	"ecommerce/helpers"
	"net/http"
)

func GetId(r *http.Request) uint64 {
	pathParams := helpers.GetPathParams(r)
	return helpers.ToUint64(pathParams["id"])
}
