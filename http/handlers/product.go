package handlers

import (
	"ecommerce/helpers"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateProduct "+helpers.GetDatetime())
}

func GetListProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetListProduct "+helpers.GetDatetime())
}

func ShowProduct(w http.ResponseWriter, r *http.Request) {
	id := GetId(r)

	fmt.Fprintf(w, "ShowProduct Id: %d %s", id, helpers.GetDatetime())
}

func UpdateProductById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateProductById "+helpers.GetDatetime())
}

func DeleteProductById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeleteProductById "+helpers.GetDatetime())
}
