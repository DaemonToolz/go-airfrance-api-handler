package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getAllFlights(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pageNumber, _ := strconv.Atoi(vars.Get("page"))
	if err := json.NewEncoder(w).Encode(sendAllFlightsRequest(vars.Get("start"), vars.Get("end"), pageNumber)); err != nil {
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

func getAllOffers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(sendOfferRequest()); err != nil {
		log.Printf(err.Error())
		panic(err)
	}
}
