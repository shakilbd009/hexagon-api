package domain

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/shakilbd009/hexagon-api/errs"
	"github.com/shakilbd009/hexagon-api/logger"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	insertQuery := "insert into accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := d.client.Exec(insertQuery, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError(unexpectedErr)
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError(unexpectedErr)
	}
	a.AccountID = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositoryDB) Update(account_id string, amount float64) *errs.AppError {
	updateQuery := "update accounts set amount = ? where account_id = ?"
	_, err := d.client.Exec(updateQuery, amount, account_id)
	if err != nil {
		logger.Error("error while updating account: " + err.Error())
		return errs.NewUnexpectedError(unexpectedErr)
	}
	return nil
}

func (d AccountRepositoryDB) Get(id string) (*Account, *errs.AppError) {
	var t Account
	println(id)
	query := "select account_id,customer_id, opening_date, account_type, amount, status from accounts where account_id = ?"
	err := d.client.Get(&t, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error(err.Error())
			return nil, errs.NewNotFoundError(fmt.Sprintf("account not found by id %s", id))
		}
		logger.Error("error while getting account by id " + err.Error())
		return nil, errs.NewUnexpectedError(unexpectedErr)
	}

	return &t, nil
}

func NewAccountRepositoryDB(client *sqlx.DB) *AccountRepositoryDB {
	return &AccountRepositoryDB{client: client}
}
