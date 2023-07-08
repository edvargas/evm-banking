package dto

import (
	"banking/errs"
	"strings"
)

type NewTransactionRequest struct {
	CustomerId      string  `json:"customer_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (receiver NewTransactionRequest) Validate() *errs.AppError {
	if strings.ToLower(receiver.TransactionType) != "withdraw" && strings.ToLower(receiver.TransactionType) != "deposit" {
		return errs.NewValidationError("Transaction type should be withdraw or deposit")
	}

	return nil
}

func (receiver NewTransactionRequest) IsWithdraw() bool {
	return strings.ToLower(receiver.TransactionType) == "withdraw"
}
