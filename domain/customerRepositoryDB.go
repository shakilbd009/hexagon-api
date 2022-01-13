package domain

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shakilbd009/hexagon-api/errs"
	"github.com/shakilbd009/hexagon-api/logger"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) ByID(id string) (*Customer, *errs.AppError) {
	query := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	//row := d.client.QueryRow(query, id)
	var c Customer
	err := d.client.Get(&c, query, id)
	//err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError(fmt.Sprintf("customer not found by id %s", id))
		}
		logger.Error("error while scanning customer by id " + err.Error())
		return nil, errs.NewUnexpectedError(unexpectedErr)
	}
	return &c, nil
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllQuery := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllQuery)
	} else {
		findAllQuery := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllQuery, status)
	}

	if err != nil {
		logger.Error("error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}

	// if err := sqlx.StructScan(rows, &customers); err != nil {
	// 	logger.Error("Error while scanning customers " + err.Error())
	// 	return nil, errs.NewUnexpectedError("unexpected database error")
	// }
	return customers, nil
}

func NewCustomerRepositoryDB(client *sqlx.DB) CustomerRepositoryDB {

	return CustomerRepositoryDB{client: client}
}
