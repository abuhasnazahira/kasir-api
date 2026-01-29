package helpers

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status  int         `json:"responseCode"`
	Message string      `json:"responseMessage,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, statusCode int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, APIResponse{
		Status:  statusCode,
		Message: message,
	})
}

func Success(w http.ResponseWriter, message string, data interface{}) {
	JSON(w, http.StatusOK, APIResponse{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func SuccessMessage(w http.ResponseWriter, message string) {
	JSON(w, http.StatusOK, APIResponse{
		Status:  http.StatusOK,
		Message: message,
	})
}
