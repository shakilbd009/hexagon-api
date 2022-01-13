package domain

import (
	"github.com/shakilbd009/hexagon-api/dto"
	"github.com/shakilbd009/hexagon-api/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}

func (c Customer) statusAsText() string {
	status := "active"
	if c.Status == "0" {
		status = "inactive"
	}
	return status
}

func (c Customer) ToDTO() *dto.CustomerResponse {

	return &dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText()}
}
