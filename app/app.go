package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shakilbd009/hexagon-api/domain"
	"github.com/shakilbd009/hexagon-api/service"
)

func Start() {
	router := mux.NewRouter()
	//wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
