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
	
	fmt.Println("Current Time: ", time)
	log.Println("Current Time: ", time)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(time)

}

func GetTimeZone(w http.ResponseWriter, r *http.Request){
	
	keys, ok := r.URL.Query()["tz1"]
	// test := r.URL.Query()
	// log.Println(test["tz2"][0])
	
	if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'tz1' is missing")
	} else {
		log.Println(keys[0])
	}
	
	// if test["tz2"][0] != ""{
	// 	log.Println(test["tz2"][0])
	// } else {
	// 	log.Println("tz2 is missing")
	// }

	
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

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", GetTime).Methods("GET")
	router.HandleFunc("/time", GetTimeZone).Methods("GET")
	log.Fatal(http.ListenAndServe(":8086", router))
}