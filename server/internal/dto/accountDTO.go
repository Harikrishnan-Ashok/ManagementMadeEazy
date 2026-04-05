package dto

import "Harikrishnan-Ashok/ManagementMadeEazy/server/internal/models"

type CreateAccountDTO struct {
	AccType    models.AccountType `json:"acc_type"`
	AccName    string             `json:"acc_name"`
	AccContact string             `json:"acc_contact"`
	AccAddress string             `json:"acc_address"`
}

type SuccessResponse struct {
	Result  string `json:"result"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
