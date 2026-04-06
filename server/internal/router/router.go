package router

import (
	"database/sql"
	"net/http"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/handlers"
)

func SetupRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", handlers.TestHandler)
	mux.HandleFunc("/create_account", handlers.CreateAccountHandler(db))
	mux.HandleFunc("/accounts", handlers.GetAllAccountsHandler(db))
	mux.HandleFunc("/account/", handlers.GetAccountByID(db))
	mux.HandleFunc("/update_account/", handlers.UpdateAccountDetailsByID(db))
	return mux
}
