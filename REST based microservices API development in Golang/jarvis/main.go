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

	if len(multiTz) == 1 {
		if tz == "" {
			current_time = time.Now().UTC().String()
			res["current_time"] = current_time
		} else {
			loc, err := time.LoadLocation(tz)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "invalid timezone")
				return
			}

			current_time = time.Now().In(loc).String()
			res["current_time"] = current_time
		}
	} else if len(multiTz) > 1 {
		for _, _tz := range multiTz {
			loc, err := time.LoadLocation(_tz)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "invalid timezone %s", _tz)
				return
			}

			res[_tz] = time.Now().In(loc).String()
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
