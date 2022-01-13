package domain

import "github.com/shakilbd009/hexagon-api/errs"

type Transaction struct {
	TransactionID   string
	AccountID       string
	TransactionDate string
	TransactionType string
	Amount          float64
}

type TransactionRepository interface {
	NewTransaction(Transaction) (*Transaction, *errs.AppError)
	GetTransaction(string) (*Transaction, *errs.AppError)
}
