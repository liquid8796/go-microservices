package main

import (
	"net/http"
)

func main() {
	//define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	//starting server
	http.ListenAndServe("localhost:8000", nil)
}
