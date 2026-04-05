package handlers

import (
	"encoding/json"
	"net/http"
)

type TestResponse struct {
	Message string `json:"message"`
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := TestResponse{
		Message: "test successful",
	}

	json.NewEncoder(w).Encode(response)
}
