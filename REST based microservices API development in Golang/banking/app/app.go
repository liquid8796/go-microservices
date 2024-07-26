package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()

	//define routes
	http.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
