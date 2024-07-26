package main

import (
	"fmt"
	"net/http"
)

func main() {
	//define routes
	http.HandleFunc("/greet", greet)

	//starting server
	http.ListenAndServe("localhost:8000", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!!")
}
