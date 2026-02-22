// Package repos
package repos

import (
	"database/sql"
	"log"
	"time"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/models"
)

func InsertIntoAccounts(db *sql.DB, data *models.Account) (sql.Result, error) {
	query := `INSERT INTO ACCOUNTS(acc_id,acc_type,acc_name,acc_address,acc_archived,acc_created_at) VALUES (?,?,?,?,?,?)`
	res, err := db.Exec(query, data.AccID, data.AccType, data.AccName, data.AccAddress, false, time.Now())
	if err != nil {
		return res, err
	}
	log.Println("data inserted successfully")
	return res, nil
}
