package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/api/time", getTime).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func getTime(w http.ResponseWriter, r *http.Request) {
	tz := r.URL.Query().Get("tz")
	multiTz := strings.Split(tz, ",")

	fmt.Printf("tz = %s, len = %d\n", tz, len(multiTz))

	res := map[string]string{}
	var current_time string

	if tz == "" {
		current_time = time.Now().UTC().String()

	} else {
		loc, err := time.LoadLocation(tz)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Location not found!")
			return
		}

		current_time = time.Now().In(loc).String()
	}

	res["current_time"] = current_time
	data, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	// w.Write(data)
	json.NewEncoder(w).Encode(string(data))
}
