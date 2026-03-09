package repos

import (
	"database/sql"
	"fmt"
	"time"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/models"

	"github.com/google/uuid"
)

func CreateTransactionRow(
	db *sql.DB,
	creditAccID string,
	debitAccID string,
	amount int64,
	transcationType models.TransactionType,
	remarks string,
	reversalID *string,
) (string, error) {
	fmt.Println("CAID:", creditAccID)
	fmt.Println("DAID:", debitAccID)
	qry := `INSERT INTO TRANSACTIONS(trans_id,credit_acc_id,debit_acc_id,trans_amount,trans_type,trans_remarks,trans_reversal_of_trans_id,trans_created_at) VALUES(?,?,?,?,?,?,?,?)`
	transID := uuid.NewString()
	_, err := db.Exec(qry, transID, creditAccID, debitAccID, amount, transcationType, remarks, reversalID, time.Now().UTC().Format(time.RFC1123))
	if err != nil {
		return "", err
	} else {
		return transID, nil
	}
}
