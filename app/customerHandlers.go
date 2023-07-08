package app

import (
	"banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, request *http.Request) {
	status := request.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(status)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandler) getCustomer(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(writer, err.Code, err.AsMessage())
	} else {
		writeResponse(writer, http.StatusOK, customer)
	}
}

func writeResponse(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		panic(err)
	}
}
