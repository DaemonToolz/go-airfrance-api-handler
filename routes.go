package main

import (
	"net/http"
)

// Route Chemin Web
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes Ensemble de chemins HTTP
type Routes []Route

// A mettre dans un JSON (et charger via Swagger ?)
var routes = Routes{
	Route{
		"Get Flights",
		"GET",
		"/api/air/v1/flights",
		getAllFlights,
	},

	Route{
		"Get Flight",
		"GET",
		"/api/air/v1/flights/{flight_id}",
		getFlight,
	},

	Route{
		"Get Offers",
		"GET",
		"/api/air/v1/offers",
		getAllOffers,
	},
}
