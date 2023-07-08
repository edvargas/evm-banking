package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new acount: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}

func (d AccountRepositoryDb) ById(accountId string) (*Account, *errs.AppError) {
	accountSelect := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where account_id = ?"
	var account Account
	err := d.client.Get(&account, accountSelect, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Account not found")
		}
		logger.Error("Error while scanning accounts " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &account, nil
}

// transactional method
func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while creating transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"

	result, _ := tx.Exec(sqlInsert, t.AccountId, t.Amount, t.TransactionId, t.TransactionDate)

	if t.IsWithdraw() {
		_, err = tx.Exec("UPDATE accounts SET amount = amount - ? where account_id = ?", t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec("UPDATE accounts SET amount = amount + ? where account_id = ?", t.Amount, t.AccountId)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Error while updating account")
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("Error while committing new transaction: " + err.Error())
		return nil, errs.NewValidationError("Error while committing transaction: " + err.Error())
	}

	account, appErr := d.ById(t.AccountId)
	if err != nil {
		return nil, appErr
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, errs.NewUnexpectedError("Error getting last insert id")
	}

	t.Amount = account.Amount
	t.TransactionId = strconv.FormatInt(id, 10)
	return &t, nil
}
