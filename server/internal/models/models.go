// Package models contains the structs that will hold the various data from the db after fetching
package models

import "time"

type (
	AccountType     string
	TransactionType string
)

const (
	BusinessAcc AccountType = "BUSINESS"
	StaffAcc    AccountType = "STAFF"
	ConsumerAcc AccountType = "CONSUMER"
)

const (
	Sale     TransactionType = "SALE"
	Purchase TransactionType = "PURCHASE"
	Salary   TransactionType = "SALARY"
	Advance  TransactionType = "ADVANCE"
	Reversal TransactionType = "REVERSAL"
)

type Account struct {
	AccID        string
	AccType      AccountType
	AccName      string
	AccContact   string
	AccAddress   string
	AccArchived  bool
	AccCreatedAt time.Time
}

type BusinessAccount struct {
	Account
	GstIN string
}

type StaffAccount struct {
	Account
	Salary int64
}

type TransactionItem struct {
	TransID            string
	TransCreditedAccID string
	TransDebitedAccID  string
	TransAmount        int64
	TransType          TransactionType
	TransRemarks       string
	TransCreatedAt     time.Time
	TransReversalOf    *string
}
