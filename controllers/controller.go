package controllers

import (
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// No book with the isdn in the url has been found
	w.WriteHeader(http.StatusOK)
}
