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
	Now string `json:"LocalTime, omitempty"`
}

type TimeZone1 struct {
	Now string `json:"time1, omitempty"`
}

type TimeZone2 struct {
	Now string `json:"time2, omitempty"`
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
	test := r.URL.Query()

	if test != nil {
		strDict := test
		for index, element := range strDict {
			fmt.Println("Index :", index, " Element :", element)
		}

		if test["tz1"] != nil {
			loc, err := time2.LoadLocation(test["tz1"][0])
			if err != nil {
				log.Println(err)
				output := "Not a valid timezone"
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(output)
			} else {
				time1 := TimeZone1{
					Now: time2.Now().In(loc).String(),
				}
				log.Println("Url Param 'tz1' is: " + string(test["tz1"][0]))
				log.Println("ZONE: ", loc, "Time: ", time1)
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(time1)
			}
		} else {
			output := "Error"
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(output)
		}

		if test["tz2"] != nil {
			loc1, err := time2.LoadLocation(test["tz2"][0])
			if err != nil {
				log.Println(err)
				output := "Not a valid timezone"
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(output)
			} else {
				time2 := TimeZone2{
					Now: time2.Now().In(loc1).String(),
				}
				log.Println("Url Param 'tz2' is: " + string(test["tz2"][0]))
				log.Println("ZONE: ", loc1, "Time: ", time2)
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(time2)
			}
		} else {
			output := "No TimeZone2 parameter passed in the Query"
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(output)
		}

	} else {
		output := "Unknown Error"
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(output)
	}

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	test := "\n System Functioning Normally"
	check := Health{
		health: test,
	}
	w.WriteHeader(http.StatusOK)
	log.Println(check)
	fmt.Fprint(w, check)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	test := "\n All Systems Ready"
	check := readiness{
		ready: test,
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
