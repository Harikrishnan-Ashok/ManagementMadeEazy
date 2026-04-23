package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/dto"
	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/models"
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

func GetTransactionList(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var req dto.TransactionFiltersDto

		q := r.URL.Query()

		if v := q.Get("start_date"); v != "" {
			req.StartDate = &v
		}
		if v := q.Get("end_date"); v != "" {
			req.EndDate = &v
		}
		if v := q.Get("trans_type"); v != "" {
			t := models.TransactionType(v)
			req.TransType = &t
		}
		if v := q.Get("page"); v != "" {
			n, err := strconv.Atoi(v)
			if err != nil {
				http.Error(w, "invalid next_page", http.StatusBadRequest)
				return
			}
			req.Page = &n
		}

		data, err := services.ListTransactionService(db, req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(dto.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(dto.SuccessResponse{
			Result:  "success",
			Message: "Accounts Successfully fetched",
			Data:    data,
		})
	}
}
