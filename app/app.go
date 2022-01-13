package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/shakilbd009/hexagon-api/domain"
	"github.com/shakilbd009/hexagon-api/logger"
	"github.com/shakilbd009/hexagon-api/service"
)

var (
	address string
	port    string
)

func sanityCheck() {
	address = os.Getenv("SERVER_ADDRESS")
	port = os.Getenv("SERVER_PORT")
	if address == "" || port == "" {
		logger.Fatal("environment variables not defined correctly")
	}
}

func Start() {
	sanityCheck()
	router := mux.NewRouter()
	client := getDbClient()
	//wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB(client))}
	ah := AccountHandler{
		service: service.NewDefaultAccountService(domain.NewAccountRepositoryDB(client))}
	th := transactionHandler{
		service: service.NewDefaultTransactionService(domain.NewTransactionRepositoryDB(client), domain.NewAccountRepositoryDB(client)),
	}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.newAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/account/{account_id:[0-9]+}", ah.getAccount).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}/transactions", th.newTransaction).Methods(http.MethodPost)

	logger.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router).Error())
}

func getDbClient() *sqlx.DB {
	user := os.Getenv("dbUser")
	pass := os.Getenv("dbPass")
	host := os.Getenv("dbHost")
	port := os.Getenv("dbPort")
	client, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/banking", user, pass, host, port))
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
