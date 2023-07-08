package domain

import (
	"banking/dto"
	"banking/errs"
)

type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

func (a Account) ToAccountResponseDto() dto.AccountResponse {
	return dto.AccountResponse{
		AccountId:   a.AccountId,
		CustomerId:  a.CustomerId,
		AccountType: a.AccountType,
		Amount:      a.Amount,
		OpeningDate: a.OpeningDate,
		Status:      a.Status,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	ById(string) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}
