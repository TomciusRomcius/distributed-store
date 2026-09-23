package controllers

import (
	"encoding/json"
	"net/http"
)

type QueryRequest struct {
	Query string `json:"query"`
}

func QueryController(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	queryRequest := &QueryRequest{}
	err := json.NewDecoder(req.Body).Decode(queryRequest)
	if err != nil {
		http.Error(w, "invalid params", http.StatusBadRequest)
		return
	}
	if queryRequest.Query == "" {
		http.Error(w, "query is required", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
