package domain

import (
	"github.com/shakilbd009/hexagon-api/dto"
	"github.com/shakilbd009/hexagon-api/errs"
)

type Account struct {
	AccountID   string  `db:"account_id,omitempty"`
	CustomerID  string  `db:"customer_id,omitempty"`
	OpeningDate string  `db:"opening_date,omitempty"`
	AccountType string  `db:"account_type,omitempty"`
	Amount      float64 `db:"amount,omitempty"`
	Status      string  `db:"status,omitempty"`
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	Get(id string) (*Account, *errs.AppError)
	Update(account_id string, amount float64) *errs.AppError
}

func (a Account) ToNewAccountResponseDTO() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{AccountID: a.AccountID}
}
