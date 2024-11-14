package handlers

import (
	"ecommerce/helpers"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateUser "+helpers.GetDatetime())
}

func GetListUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetListUser "+helpers.GetDatetime())
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	id := GetId(r)

	fmt.Fprintf(w, "ShowUser Id: %d %s", id, helpers.GetDatetime())
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateUserById "+helpers.GetDatetime())
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeleteUserById "+helpers.GetDatetime())
}
