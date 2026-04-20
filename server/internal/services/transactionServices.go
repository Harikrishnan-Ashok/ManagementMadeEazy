package services

import (
	"database/sql"
	"time"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/dto"
	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/models"
	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/repos"

	"github.com/google/uuid"
)

func CreateTransactionService(db *sql.DB, req dto.CreateTransactionDTO) error {
	transaction := models.TransactionItem{
		TransID:            uuid.New().String(),
		TransCreditedAccID: req.CreditAccID.String(),
		TransDebitedAccID:  req.DebitAccID.String(),
		TransAmount:        req.TransAmount,
		TransType:          req.TransType,
		TransRemarks:       req.TransRemarks,
		TransCreatedAt:     time.Now().Format(time.RFC3339Nano),
		TransReversalOf:    nil,
	}

	// call repo fun
	_, err := repos.CreateTransactionRow(db, &transaction)
	if err != nil {
		return err
	}
	return nil
}
