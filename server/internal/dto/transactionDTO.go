package dto

import (
	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/models"

	"github.com/google/uuid"
)

type CreateTransactionDTO struct {
	CreditAccID  uuid.UUID              `json:"credit_acc_id"`
	DebitAccID   uuid.UUID              `json:"debit_acc_id"`
	TransAmount  int64                  `json:"trans_amount"`
	TransRemarks string                 `json:"trans_remarks"`
	TransType    models.TransactionType `json:"trans_type"`
}

type TransactionFiltersDto struct {
	Page      *int                    `json:"page"`
	StartDate *string                 `json:"start_date"`
	EndDate   *string                 `json:"end_date"`
	TransType *models.TransactionType `json:"trans_type"`
}
