package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	time2 "time"

	"github.com/gorilla/mux"
)

type Time struct {
	Now string `json:"time, omitempty"`
}

type TimeZone struct {
	Now string `json:"time, omitempty"`
}

type Health struct {
	health string `json:"Health, omitempty"`
}

type readiness struct {
	ready string `json:"readiness, omitempty"`
}

func GetTime(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	time := Time{
		Now: time2.Now().String(),
	}

	fmt.Println("Current Time: ", time)
	log.Println("Current Time: ", time)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(time)

}

func GetTimeZone(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["tz1"]
	test := r.URL.Query()

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'tz1' is missing")
	} else {
		log.Println(keys[0])
	}

	strDict := test
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}

	tz1 := keys[0]

	log.Println("Url Param 'tz1' is: " + string(tz1))

	loc, _ := time2.LoadLocation(tz1)

	time := TimeZone{
		Now: time2.Now().In(loc).String(),
	}

	log.Println("ZONE: ", loc, "Time: ", time)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(time)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	check := Health{
		health: "Sytem functioning normally",
	}
	w.WriteHeader(http.StatusOK)
	log.Println(check)
	fmt.Fprint(w, check)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	check := readiness{
		ready: "Sytem Ready",
	}
	w.WriteHeader(http.StatusOK)
	log.Println(check)
	fmt.Fprint(w, check)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetTime).Methods("GET")
	router.HandleFunc("/time", GetTimeZone).Methods("GET")
	router.HandleFunc("/health", healthHandler).Methods("GET")
	router.HandleFunc("/readiness", readinessHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8086", router))
}
