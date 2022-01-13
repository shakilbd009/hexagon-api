package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shakilbd009/hexagon-api/dto"
	"github.com/shakilbd009/hexagon-api/errs"
	"github.com/shakilbd009/hexagon-api/service"
)

type transactionHandler struct {
	service service.TransactionService
}

func (t transactionHandler) newTransaction(w http.ResponseWriter, r *http.Request) {
	var nt dto.NewTransactionRequest
	vars := mux.Vars(r)
	account_id, ok := vars["account_id"]
	if !ok {
		appErr := errs.NewUnprocessableEntityError("account_id is missing in qparameter")
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}
	customer_id, ok := vars["customer_id"]
	if !ok {
		appErr := errs.NewUnprocessableEntityError("customer_id is missing in qparameter")
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&nt); err != nil {
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	nt.AccountID = account_id
	nt.CustomerID = customer_id
	resp, appErr := t.service.NewTransaction(nt)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, &resp)
}
