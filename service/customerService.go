package service

import (
	"github.com/shakilbd009/hexagon-api/domain"
	"github.com/shakilbd009/hexagon-api/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerSevice struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerSevice) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return d.repo.ByID(id)
}

func (d DefaultCustomerSevice) GetAllCustomer() ([]domain.Customer, *errs.AppError) {
	return d.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerSevice {
	return DefaultCustomerSevice{repo: repository}
}
