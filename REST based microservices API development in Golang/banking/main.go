package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {
	//define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	//starting server
	http.ListenAndServe("localhost:8000", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Nam", City: "HCM", Zipcode: "700000"},
		{Name: "Trang", City: "HCM", Zipcode: "700900"},
	}

	w.Header().Add("Content-Type", "application/xml")

	// json.NewEncoder(w).Encode(customers)
	xml.NewEncoder(w).Encode(customers)
}
