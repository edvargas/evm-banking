package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"strings"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	GetAccount(accountId string) (*dto.AccountResponse, *errs.AppError)
	MakeTransaction(request dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (d DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	account := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: strings.ToLower(req.AccountType),
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, err := d.repo.Save(account)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()

	return &response, nil
}

func (d DefaultAccountService) GetAccount(accountId string) (*dto.AccountResponse, *errs.AppError) {
	account, err := d.repo.ById(accountId)

	if err != nil {
		return nil, err
	}

	response := account.ToAccountResponseDto()

	return &response, nil
}

func (d DefaultAccountService) MakeTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	if req.IsWithdraw() {
		account, appError := d.repo.ById(req.AccountId)
		if appError != nil {
			return nil, appError
		}

		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}

	transaction := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	newTransaction, appError := d.repo.SaveTransaction(transaction)
	if appError != nil {
		return nil, appError
	}

	response := newTransaction.ToNewTransactionResponseDto()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
