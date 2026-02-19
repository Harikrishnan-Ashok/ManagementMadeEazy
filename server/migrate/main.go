package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	// Connect to the SQLite database
	db, err := sql.Open("sqlite", "./mmeMain.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = db.Exec("PRAGMA forein_keys = ON;")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("enforced Forien Keys")

	defer db.Close()
	fmt.Println("Connected to the SQLite database successfully.")

	schema, err := os.ReadFile("./internal/db/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Schema Applied")
}
