package domain

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shakilbd009/hexagon-api/errs"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) ByID(id string) (*Customer, *errs.AppError) {
	query := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row := d.client.QueryRow(query, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError(fmt.Sprintf("customer not found by id %s", id))
		}
		log.Println("error while scanning customer by id " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &c, nil
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {

	findAllQuery := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	rows, err := d.client.Query(findAllQuery)
	if err != nil {
		log.Println("error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("error while querying customer table " + err.Error())
			return nil, errs.NewUnexpectedError(err.Error())
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	user := os.Getenv("dbUser")
	pass := os.Getenv("dbPass")
	host := os.Getenv("dbHost")
	port := os.Getenv("dbPort")
	client, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/banking", user, pass, host, port))
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client: client}
}
