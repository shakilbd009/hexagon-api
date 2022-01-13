package service

import (
	"github.com/shakilbd009/hexagon-api/domain"
	"github.com/shakilbd009/hexagon-api/dto"
	"github.com/shakilbd009/hexagon-api/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerSevice struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerSevice) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {

	c, err := d.repo.ByID(id)
	if err != nil {
		return nil, err
	}
	return c.ToDTO(), nil
}

func (d DefaultCustomerSevice) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	customers, err := d.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	responses := make([]dto.CustomerResponse, len(customers), len(customers))
	for i, c := range customers {
		//responses = append(responses, *c.ToDTO())
		responses[i] = *c.ToDTO()
	}
	return responses, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerSevice {
	return DefaultCustomerSevice{repo: repository}
}
