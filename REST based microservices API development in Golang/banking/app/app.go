package app

import (
	"banking/domain"
	"banking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	mux := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	//define routes
	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
