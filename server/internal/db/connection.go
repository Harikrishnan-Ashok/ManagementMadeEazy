// Package db contains the connection files and the sql schemas
package db

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

func OpenDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return db, err
	}

	_, err = db.Exec("PRAGMA forein_keys = ON;")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("ENABLED forein_keys")
	}
	return db, nil
}
