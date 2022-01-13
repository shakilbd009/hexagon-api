package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shakilbd009/hexagon-api/dto"
	"github.com/shakilbd009/hexagon-api/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah AccountHandler) newAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var request dto.NewAccountRequest
	request.CustomerID = vars["customer_id"]
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	}
	resp, appErr := ah.service.NewAccount(request)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
	}
	writeResponse(w, http.StatusOK, resp)
}

func (ah AccountHandler) getAccount(w http.ResponseWriter, r *http.Request) {
	if id, ok := mux.Vars(r)["account_id"]; ok {
		var request dto.GetAccountRequest
		request.AccountID = id
		resp, err := ah.service.GetAccount(id)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
			return
		}
		writeResponse(w, http.StatusOK, resp)
	}
}
