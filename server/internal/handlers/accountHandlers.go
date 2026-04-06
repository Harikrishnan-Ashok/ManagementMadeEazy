package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/dto"
	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/services"

	"github.com/google/uuid"
)

func CreateAccountHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var req dto.CreateAccountDTO

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, "Invalid Requst Body", http.StatusBadRequest)
			return
		}

		err = services.CreateAccountService(db, req)
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(dto.SuccessResponse{Message: "account Creation Successfull"})
	}
}

func GetAllAccountsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		archivedParam := r.URL.Query().Get("archived")
		archived := false
		if archivedParam != "" {
			val, err := strconv.ParseBool(archivedParam)
			if err != nil {
				http.Error(w, "Invaild `Archived` query param", http.StatusBadRequest)
				return
			}
			archived = val
		}

		data, err := services.GetAllAccountService(db, archived)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(dto.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(dto.SuccessResponse{
			Result:  "success",
			Message: "Accounts Successfully fetched",
			Data:    data,
		})
	}
}

func GetAccountByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		data := &dto.SingleAccountDetails{}
		idStr := strings.TrimPrefix(r.URL.Path, "/account/")
		if idStr != "" {
			parsedID, err := uuid.Parse(idStr)
			if err != nil {
				http.Error(w, "Unable to Parse UUID", http.StatusBadRequest)
				return
			}
			data, err = services.GetSingleAccountByID(db, parsedID)
			if err != nil {
				http.Error(w, "Unable to Get Account", http.StatusInternalServerError)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(dto.SuccessResponse{
			Result:  "success",
			Message: "Account Found",
			Data:    data,
		})
	}
}

func UpdateAccountDetailsByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		var body dto.UpdateAccountDetails
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&body); err != nil {
			http.Error(w, "Invalid Body", http.StatusBadRequest)
			return
		}

		idStr := strings.TrimPrefix(r.URL.Path, "/update_account/")
		if idStr == "" {
			http.Error(w, "Missing account id", http.StatusBadRequest)
			return
		}
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "Unable to Parse UUID", http.StatusBadRequest)
			return
		}
		err = services.UpdateAccountDetailsByID(db, id, &body)
		if err != nil {
			http.Error(w, "Update Failed", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(dto.SuccessResponse{
			Result:  "success",
			Message: "Account Updated",
		})
	}
}
