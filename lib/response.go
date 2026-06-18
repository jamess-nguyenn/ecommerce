package lib

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewResponse(w http.ResponseWriter, status string, code int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if data == nil {
		data = map[string]any{}
	}

	json.NewEncoder(w).Encode(&Response{Status: status, Code: code, Message: message, Data: data})
}

func ResponseSuccess(w http.ResponseWriter, data any, message string, code int) {
	NewResponse(w, "success", code, message, data)
}

func ResponseOK(w http.ResponseWriter, data any, message string) {
	ResponseSuccess(w, data, message, http.StatusOK)
}

func ResponseCreated(w http.ResponseWriter, data any, message string) {
	ResponseSuccess(w, data, message, http.StatusCreated)
}

func ResponseError(w http.ResponseWriter, message string, data any, code int) {
	NewResponse(w, "error", code, message, data)
}

func ResponseInvalidRequest(w http.ResponseWriter, data any) {
	ResponseError(w, "Invalid Request", data, http.StatusBadRequest)
}

func ResponseInternalServerError(w http.ResponseWriter, data any) {
	ResponseError(w, "Internal Server Error", data, http.StatusInternalServerError)
}
