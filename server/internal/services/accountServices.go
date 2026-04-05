package services

import (
	"database/sql"
	"strings"
	"time"

	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/dto"
	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/models"
	"Harikrishnan-Ashok/ManagementMadeEazy/server/internal/repos"

	"github.com/google/uuid"
)

func CreateAccountService(db *sql.DB, requestData dto.CreateAccountDTO) error {
	acc := models.Account{
		AccID:        uuid.New().String(),
		AccType:      models.AccountType(strings.ToUpper(string(requestData.AccType))),
		AccName:      requestData.AccName,
		AccContact:   requestData.AccContact,
		AccAddress:   requestData.AccAddress,
		AccArchived:  false,
		AccCreatedAt: time.Now().Format(time.RFC3339Nano),
	}

	_, err := repos.InsertIntoAccounts(db, &acc)
	if err != nil {
		return err
	}
	return nil
}
