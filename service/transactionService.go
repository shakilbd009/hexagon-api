package service

import (
	"github.com/shakilbd009/hexagon-api/domain"
	"github.com/shakilbd009/hexagon-api/dto"
	"github.com/shakilbd009/hexagon-api/errs"
)

type TransactionService interface {
	NewTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultTransactionService struct {
	repo        domain.TransactionRepository
	accountRepo domain.AccountRepository
}

func NewDefaultTransactionService(tRepo domain.TransactionRepository, aRepo domain.AccountRepository) *DefaultTransactionService {
	return &DefaultTransactionService{repo: tRepo, accountRepo: aRepo}
}

func (d DefaultTransactionService) NewTransaction(t dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {

	account, err := d.accountRepo.Get(t.AccountID)
	if err != nil {
		return nil, err
	}
	if err := t.Validate(account.Amount); err != nil {
		return nil, err
	}
	newAmount := account.Amount - t.Amount
	if err := d.accountRepo.Update(account.AccountID, newAmount); err != nil {
		return nil, err
	}
	transaction := domain.Transaction{
		AccountID:       account.AccountID,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
	}
	newTransaction, err := d.repo.NewTransaction(transaction)
	if err != nil {
		return nil, err
	}

	return &dto.NewTransactionResponse{
		TransactionID: newTransaction.TransactionID,
		Amount:        newTransaction.Amount,
	}, nil
}
