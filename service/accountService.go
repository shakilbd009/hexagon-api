package service

import (
	"time"

	"github.com/shakilbd009/hexagon-api/domain"
	"github.com/shakilbd009/hexagon-api/dto"
	"github.com/shakilbd009/hexagon-api/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	GetAccount(string) (*dto.GetAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) GetAccount(id string) (*dto.GetAccountResponse, *errs.AppError) {
	account, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return &dto.GetAccountResponse{
		AccountID:   account.AccountID,
		CustomerID:  account.CustomerID,
		OpeningDate: account.OpeningDate,
		AccountType: account.AccountType,
		Amount:      account.Amount,
		Status:      account.Status,
	}, nil

}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	d := domain.Account{
		AccountID:   "",
		CustomerID:  req.CustomerID,
		OpeningDate: time.Now().Format("2006-01-02, 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	account, err := s.repo.Save(d)
	if err != nil {
		return nil, err
	}
	return account.ToNewAccountResponseDTO(), nil
}

func NewDefaultAccountService(repo domain.AccountRepository) *DefaultAccountService {
	return &DefaultAccountService{repo: repo}
}
