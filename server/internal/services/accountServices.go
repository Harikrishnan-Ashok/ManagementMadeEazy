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

func GetAllAccountService(db *sql.DB, archived bool) ([]*dto.SingleAccountDetails, error) {
	accounts, err := repos.GetAllAccounts(db, archived)
	if err != nil {
		return nil, err
	}

	result := []*dto.SingleAccountDetails{}

	for _, acc := range accounts {
		result = append(result, &dto.SingleAccountDetails{
			AccID:      acc.AccID,
			AccName:    acc.AccName,
			AccContact: acc.AccContact,
			AccAddress: acc.AccAddress,
		})
	}

	return result, nil
}

func GetSingleAccountByID(db *sql.DB, id uuid.UUID) (*dto.SingleAccountDetails, error) {
	result, err := repos.GetAccountDetailsByID(db, id)
	if err != nil {
		return nil, err
	}
	account := &dto.SingleAccountDetails{
		AccID:        result.AccID,
		AccName:      result.AccName,
		AccContact:   result.AccContact,
		AccAddress:   result.AccAddress,
		AccArchived:  result.AccArchived,
		AccCreatedAt: result.AccCreatedAt,
	}

	return account, nil
}

func UpdateAccountDetailsByID(
	db *sql.DB,
	id uuid.UUID,
	body *dto.UpdateAccountDetails,
) error {
	updates := []string{}
	args := []any{}

	if body.AccName != nil {
		updates = append(updates, "acc_name=?")
		args = append(args, *body.AccName)
	}

	if body.AccContact != nil {
		updates = append(updates, "acc_contact=?")
		args = append(args, *body.AccContact)
	}

	if body.AccAddress != nil {
		updates = append(updates, "acc_address=?")
		args = append(args, *body.AccAddress)
	}

	if body.AccArchived != nil {
		updates = append(updates, "acc_archived=?")
		args = append(args, *body.AccArchived)
	}

	if len(updates) == 0 {
		return nil
	}

	return repos.UpdateAccountDetails(db, id, updates, args)
}
