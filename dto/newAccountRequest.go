package dto

import (
	"strings"

	"github.com/shakilbd009/hexagon-api/errs"
)

type NewAccountRequest struct {
	CustomerID  string  `json:"customer_id,omitempty"`
	AccountType string  `json:"account_type,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("to open a new account you need to deposit atleast 5000.00")
	}
	if strings.ToLower(r.AccountType) != "savings" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("account type should be checking or savings")
	}
	return nil
}
