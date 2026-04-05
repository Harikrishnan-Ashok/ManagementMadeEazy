package main

import (
	"fmt"
	"log"
	"net/http"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/db"
	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/router"
)

func main() {
	fmt.Println("Trying to connect with DB")

	conn, err := db.OpenDB("./mmeMain.db")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Connected to mmeMain successfully")

	r := router.SetupRouter(conn)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Server running at http://localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
