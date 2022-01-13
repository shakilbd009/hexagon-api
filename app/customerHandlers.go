package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shakilbd009/hexagon-api/service"
)

const (
	content_type     = "Content-Type"
	application_xml  = "application/xml"
	application_json = "application/json"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomer(queryString)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		fmt.Println(err)
	} else {
		writeResponse(w, http.StatusOK, customers)
	}

}

func (ch CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer, err := ch.service.GetCustomer(vars["customer_id"])
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}

}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add(content_type, application_json)
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
