// Package repos
package repos

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/models"

	"github.com/google/uuid"
)

// InsertIntoAccounts to insert 1 Account with its details into the db
func InsertIntoAccounts(db *sql.DB, data *models.Account) (sql.Result, error) {
	query := `INSERT INTO ACCOUNTS(acc_id,acc_type,acc_name,acc_contact,acc_address,acc_archived,acc_created_at) VALUES (?,?,?,?,?,?,?)`
	res, err := db.Exec(query, data.AccID, data.AccType, data.AccName, data.AccContact, data.AccAddress, data.AccArchived, data.AccCreatedAt)
	if err != nil {
		log.Println(err)
		return res, err
	}
	return res, nil
}

// GetAccountDetailsByID to read a single Account details sorted by ID
func GetAccountDetailsByID(db *sql.DB, id uuid.UUID) (*models.Account, error) {
	qry := `SELECT acc_id,acc_type,acc_name,acc_contact,acc_address,acc_archived,acc_created_at FROM ACCOUNTS WHERE acc_id=?`
	res := db.QueryRowContext(context.Background(), qry, id)
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
	return acc, nil
}

func GetAllAccounts(db *sql.DB, archived bool) ([]*models.Account, error) {
	qry := `SELECT acc_id, acc_name, acc_contact, acc_address FROM accounts WHERE acc_archived=?`
	rows, err := db.QueryContext(context.Background(), qry, archived)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	accounts := []*models.Account{}
	for rows.Next() {
		var acc models.Account
		err := rows.Scan(
			&acc.AccID,
			&acc.AccName,
			&acc.AccContact,
			&acc.AccAddress,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &acc)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}

func UpdateAccountDetails(
	db *sql.DB,
	id uuid.UUID,
	updates []string,
	args []any,
) error {
	query := "UPDATE accounts SET " +
		strings.Join(updates, ", ") +
		" WHERE acc_id=?"

	args = append(args, id)

	_, err := db.Exec(query, args...)
	return err
}
