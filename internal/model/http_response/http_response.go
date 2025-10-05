package httpresponse

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func JSONError(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(Response{
		Status: statusCode,
		Msg:    msg,
	})
}

func JSONSuccess(w http.ResponseWriter, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(payload)
}
