// Package db
package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

func OpenDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	// Set foreign keys via Exec, not connection string
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	// Verify it actually took
	var fk int
	db.QueryRow("PRAGMA foreign_keys").Scan(&fk)
	if fk != 1 {
		log.Println("foreign key pragma failed")
		return nil, fmt.Errorf("foreign keys not enabled")
	} else {
		fmt.Println("success")
	}

	return db, nil
}
