package controllers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendJSONResponse(w http.ResponseWriter, data interface{}) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := Response{Code: http.StatusOK, Message: "success", Data: data}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
