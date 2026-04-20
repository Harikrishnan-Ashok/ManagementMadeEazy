package router

import (
	"database/sql"
	"net/http"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/handlers"
)

func SetupRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	// account api routes
	mux.HandleFunc("/create_account", handlers.CreateAccountHandler(db))
	mux.HandleFunc("/accounts", handlers.GetAllAccountsHandler(db))
	mux.HandleFunc("/account/", handlers.GetAccountByID(db))
	mux.HandleFunc("/update_account/", handlers.UpdateAccountDetailsByID(db))

	// transacation api routes
	mux.HandleFunc("/transaction", handlers.CreateTransactionHandler(db))
	return mux
}
