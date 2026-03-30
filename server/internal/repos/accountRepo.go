// Package repos
package repos

import (
	"database/sql"
	"log"
	"time"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/models"
)

// InsertIntoAccounts to insert 1 Account with its details into the db
func InsertIntoAccounts(db *sql.DB, data *models.Account) (sql.Result, error) {
	query := `INSERT INTO ACCOUNTS(acc_id,acc_type,acc_name,acc_contact,acc_address,acc_archived,acc_created_at) VALUES (?,?,?,?,?,?,?)`
	res, err := db.Exec(query, data.AccID, data.AccType, data.AccName, data.AccContact, data.AccAddress, false, time.Now())
	if err != nil {
		return res, err
	}
	log.Println("data inserted successfully")
	return res, nil
}

// GetAccountDetailsByID to read a single Account details sorted by ID
func GetAccountDetailsByID(db *sql.DB, id string) (*models.Account, error) {
	qry := `SELECT acc_id,acc_type,acc_name,acc_contact,acc_address,acc_archived,acc_created_at FROM ACCOUNTS WHERE acc_id=?`
	res := db.QueryRow(qry, id)
	acc := &models.Account{}
	err := res.Scan(
		&acc.AccID,
		&acc.AccType,
		&acc.AccName,
		&acc.AccContact,
		&acc.AccAddress,
		&acc.AccArchived,
		&acc.AccCreatedAt,
	)
	if err != nil {
		log.Println("error while trying to retrive the data ")
		return acc, err
	}
	log.Println("retrived data")
	return acc, nil
}
