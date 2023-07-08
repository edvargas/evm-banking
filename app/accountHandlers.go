package app

import (
	"banking/dto"
	"banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (a AccountHandler) NewAccount(writer http.ResponseWriter, request *http.Request) {
	var newAccount dto.NewAccountRequest
	customerId := mux.Vars(request)["customer_id"]

	err := json.NewDecoder(request.Body).Decode(&newAccount)
	if err != nil {
		writeResponse(writer, http.StatusBadRequest, err.Error())
	} else {
		newAccount.CustomerId = customerId
		account, appError := a.service.NewAccount(newAccount)
		if appError != nil {
			writeResponse(writer, appError.Code, appError.Message)
		} else {
			writeResponse(writer, http.StatusCreated, account)
		}
	}
}
func (a AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	var newTransaction dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&newTransaction)
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	accountId := vars["account_id"]

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		newTransaction.CustomerId = customerId
		newTransaction.AccountId = accountId
		transaction, appError := a.service.MakeTransaction(newTransaction)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, transaction)
		}
	}
}
