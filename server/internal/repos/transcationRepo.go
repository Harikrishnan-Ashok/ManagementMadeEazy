package repos

import (
	"database/sql"
	"log"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/models"
)

func CreateTransactionRow(db *sql.DB, data *models.TransactionItem) (sql.Result, error) {
	qry := `INSERT INTO TRANSACTIONS(
    trans_id,
    credit_acc_id,
    debit_acc_id,
    trans_amount,
    trans_type,
    trans_remarks,
    trans_reversal_of_trans_id,
    trans_created_at
) VALUES(?,?,?,?,?,?,?,?)`

	res, err := db.Exec(
		qry,
		data.TransID,
		data.TransCreditedAccID,
		data.TransDebitedAccID,
		data.TransAmount,
		data.TransType,
		data.TransRemarks,
		data.TransReversalOf,
		data.TransCreatedAt,
	)
	if err != nil {
		log.Println(err)
		return res, err
	}
	return res, nil
}
