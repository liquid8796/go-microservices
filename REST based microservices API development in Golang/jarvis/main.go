package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/api/time", func(w http.ResponseWriter, r *http.Request) {})
}
