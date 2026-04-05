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
	return mux
}
