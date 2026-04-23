package repos

import (
	"database/sql"
	"log"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/dto"
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

func GetTransactionRows(db *sql.DB, filters dto.TransactionFiltersDto) ([]models.TransactionItem, error) {
	query := `SELECT 
		trans_id,
		credit_acc_id,
		debit_acc_id,
		trans_amount,
		trans_type,
		trans_remarks,
		trans_created_at,
		trans_reversal_of_trans_id
	FROM transactions
	WHERE 1=1`

	args := []any{}

	// Filters
	if filters.StartDate != nil {
		query += " AND trans_created_at >= ?"
		args = append(args, *filters.StartDate)
	}

	if filters.EndDate != nil {
		query += " AND trans_created_at <= ?"
		args = append(args, *filters.EndDate)
	}

	if filters.TransType != nil {
		query += " AND trans_type = ?"
		args = append(args, *filters.TransType)
	}

	// Pagination
	limit := 10
	offset := 0

	if filters.Page != nil {
		offset = (*filters.Page) * limit
	}

	query += " ORDER BY trans_created_at DESC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	// Execute
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []models.TransactionItem{}

	for rows.Next() {
		var t models.TransactionItem

		err := rows.Scan(
			&t.TransID,
			&t.TransCreditedAccID,
			&t.TransDebitedAccID,
			&t.TransAmount,
			&t.TransType,
			&t.TransRemarks,
			&t.TransCreatedAt,
			&t.TransReversalOf,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
