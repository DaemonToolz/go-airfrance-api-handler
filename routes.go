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
		"/flights",
		getAllFlights,
	},

	Route{
		"Get Flight",
		"GET",
		"/flights/{flight_id}",
		getFlight,
	},
}

// 2h68dyzmx4tbxx4asxua599b Key
// N1kI9WL6cM Secret
