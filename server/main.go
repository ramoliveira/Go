package main

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", ReturnError)
	http.ListenAndServe(":8080", nil)
}

func ReturnError(w http.ResponseWriter, r *http.Request) {
	errorResponse := ErrorResponse{500, "Damn"}

	js, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	w.Write(js)
}
