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

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", GetTime).Methods("GET")
	log.Fatal(http.ListenAndServe(":8086", router))
}