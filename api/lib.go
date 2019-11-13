package main

import (
	"net/http"
	"encoding/json"
)

func ResponseError(w http.ResponseWriter, statusCode int, errMsg []string) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	response := map[string][]string{"message": errMsg}
	json.NewEncoder(w).Encode(response)
}
