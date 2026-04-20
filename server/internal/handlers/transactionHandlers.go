package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/dto"
	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/services"
)

func CreateTransactionHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var req dto.CreateTransactionDTO

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, "Invalid Request Body", http.StatusBadRequest)
			return
		}

		err = services.CreateTransactionService(db, req)
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(dto.SuccessResponse{Message: "Transcation Created Successfully"})
	}
}
