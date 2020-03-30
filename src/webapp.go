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

func GetTime(w http.ResponseWriter, r *http.Request) {
	
	time := Time{
		Now: time2.Now().String(),
	}

	fmt.Println("UserAgent %v", r.Header.Get("UserAgent"))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(time)

}

func GetTimeZone(w http.ResponseWriter, r *http.Request){
	
	keys, ok := r.URL.Query()["tz1"]
	
	if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }
	
	tz1 := keys[0]
	
	log.Println("Url Param 'key' is: " + string(tz1))

	loc, _ := time2.LoadLocation(tz1)

	time := TimeZone{
		Now: time2.Now().In(loc).String(),
	}

    fmt.Println("ZONE: ", loc, "Time: ", time)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(time)
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", GetTime).Methods("GET")
	router.HandleFunc("/time", GetTimeZone).Methods("GET")
	log.Fatal(http.ListenAndServe(":8086", router))
}