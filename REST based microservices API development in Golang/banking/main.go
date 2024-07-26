package main

import "net/http"

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {

	})
}
