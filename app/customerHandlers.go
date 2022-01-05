package app

import (
	"encoding/json"
	"encoding/xml"
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

	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		fmt.Println(err)
	}

	if r.Header.Get(content_type) == application_xml {
		w.Header().Add(content_type, application_xml)
		if err := xml.NewEncoder(w).Encode(customers); err != nil {
			fmt.Println(err)
		}
	}
	if r.Header.Get(content_type) == application_json {
		w.Header().Add(content_type, application_json)
		if err := json.NewEncoder(w).Encode(customers); err != nil {
			fmt.Println(err)
		}
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
