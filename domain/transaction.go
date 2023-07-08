package domain

import (
	"banking/dto"
	"strings"
)

type Transaction struct {
	TransactionId   string
	AccountId       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

func (t Transaction) ToNewTransactionResponseDto() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}

func (t Transaction) IsWithdraw() bool {
	return strings.ToLower(t.TransactionType) == "withdraw"
}
