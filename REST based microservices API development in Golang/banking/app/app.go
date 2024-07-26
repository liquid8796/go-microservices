package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	mux := mux.NewRouter()

	//define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)
	mux.HandleFunc("/customers/{customer_id}", getCustomer)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	mux.Vars(r)
}
