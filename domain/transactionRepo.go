package domain

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/shakilbd009/hexagon-api/errs"
	"github.com/shakilbd009/hexagon-api/logger"
)

type TransactionRepositoryDB struct {
	client *sqlx.DB
}

func (d TransactionRepositoryDB) NewTransaction(t Transaction) (*Transaction, *errs.AppError) {

	query := "insert into transactions ( account_id, amount, transaction_type) values ( ?, ?, ?)"
	result, err := d.client.Exec(query, t.AccountID, t.Amount, t.TransactionType)
	if err != nil {
		logger.Error("error while updating transaction: " + err.Error())
		return nil, errs.NewUnexpectedError(unexpectedErr)
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while getting last insertId: " + err.Error())
		return nil, errs.NewUnexpectedError(unexpectedErr)
	}
	t.TransactionID = strconv.FormatInt(id, 10)
	return &t, nil
}

func (d TransactionRepositoryDB) GetTransaction(id string) (*Transaction, *errs.AppError) {
	var t Transaction
	query := "select account_id,customer_id, opening_date, account_type, amount, status from accounts where account_id = ?"
	err := d.client.Get(&t, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError(fmt.Sprintf("account not found by id %s", id))
		}
		logger.Error("error while scanning customer by id " + err.Error())
		return nil, errs.NewUnexpectedError(unexpectedErr)
	}
	return &t, nil
}

func NewTransactionRepositoryDB(client *sqlx.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{client: client}
}
