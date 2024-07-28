package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/api/time", getTime).Methods(http.MethodGet)

}

func getTime(w http.ResponseWriter, r *http.Request) {
	tz := r.URL.Query().Get("tz")
	res := map[string]string{}
	var current_time string

	if tz == "" {
		current_time = time.UTC.String()
		res["current_time"] = current_time
	} else {
		loc, _ := time.LoadLocation(tz)
		current_time = time.Now().In(loc).String()
	}

	w.Header().Add("Content-Type", "application/xml")
	json.NewEncoder(w).Encode(current_time)
}
