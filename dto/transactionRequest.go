package dto

import "github.com/shakilbd009/hexagon-api/errs"

const (
	withdrawl = "withdrawl"
	deposit   = "deposit"
)

type NewTransactionRequest struct {
	CustomerID      string  `json:"-"`
	TransactionType string  `json:"transaction_type,omitempty"`
	TransactionDate string  `json:"transaction_date,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
	AccountID       string  `json:"account_id,omitempty"`
}

func (t NewTransactionRequest) Validate(fund float64) *errs.AppError {
	if t.TransactionType != withdrawl || t.TransactionType != deposit {
		return errs.NewValidationError("transaction type can only be deposit or withdrawl")
	}
	if t.TransactionType == withdrawl {
		if t.Amount < 0 {
			return errs.NewValidationError("amount cannot be less than zero")
		}
		if t.Amount > fund {
			return errs.NotEnoughFundError("you dont have enough fund for this withdrawl")
		}
	}
	return nil
}
