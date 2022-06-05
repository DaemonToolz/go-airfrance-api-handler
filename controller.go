package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getAllFlights(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(sendAllFlightsRequest()); err != nil {
		log.Printf(err.Error())
		panic(err)
	}
}

func getFlight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(sendFlightRequest(vars["flight_id"])); err != nil {
		log.Printf(err.Error())
		panic(err)
	}
}
