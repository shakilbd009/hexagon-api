package domain

import "github.com/shakilbd009/hexagon-api/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ByID(id string) (*Customer, *errs.AppError)
}
