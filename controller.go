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

	if err := json.NewEncoder(w).Encode(sendOfferRequest(r.URL.Query().Get("date"), r.URL.Query().Get("from"), r.URL.Query().Get("to")).FlightProduct); err != nil {
		log.Printf(err.Error())
		panic(err)
	}
}

func getOfferDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var url InputUrl
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(sendOfferDetailFromLink(url.URL, url.ReqType)); err != nil {
		log.Printf(err.Error())
		panic(err)
	}
}

func getAllStations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(sendStationsRequest().StationCities); err != nil {
		log.Printf(err.Error())
		panic(err)
	}
}
