package dto

import "Harikrishnan-Ashok/ManagementMadeEazy/server/internal/models"

type CreateAccountDTO struct {
	AccType    models.AccountType `json:"acc_type"`
	AccName    string             `json:"acc_name"`
	AccContact string             `json:"acc_contact"`
	AccAddress string             `json:"acc_address"`
}

type SingleAccountDetails struct {
	AccID        string `json:"acc_id"`
	AccName      string `json:"acc_name"`
	AccContact   string `json:"acc_contact"`
	AccAddress   string `json:"acc_address"`
	AccArchived  bool   `json:"acc_archived"`
	AccCreatedAt string `json:"acc_created_at"`
}

type SuccessResponse struct {
	Result  string `json:"result"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type UpdateAccountDetails struct {
	AccName     *string `json:"acc_name"`
	AccContact  *string `json:"acc_contact"`
	AccAddress  *string `json:"acc_address"`
	AccArchived *bool   `json:"acc_archived"`
}
