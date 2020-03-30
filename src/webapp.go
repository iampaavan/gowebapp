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

func GetTime(w http.ResponseWriter, r *http.Request) {
	
	time := Time{
		Now: time2.Now().String(),
	}

	fmt.Println("UserAgent %v", r.Header.Get("UserAgent"))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(time)

}

func GetTimeZone(w http.ResponseWriter, r *http.Request){
	
	keys, ok := r.URL.Query()["key"]
	
	if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }
	
	key := keys[0]
	
	log.Println("Url Param 'key' is: " + string(key))

	loc, _ := time2.LoadLocation(key)
    now := time2.Now().In(loc)
    fmt.Println("ZONE : ", loc, " Time : ", now)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(now)
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", GetTime).Methods("GET")
	router.HandleFunc("/time", GetTimeZone).Methods("GET")
	log.Fatal(http.ListenAndServe(":8086", router))
}