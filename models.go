package main

import "time"

//
type FlightsData struct {
	Flights []Flight `json:"operationalFlights"`
}

type Flight struct {
	Number        int         `json:"flightNumber"`
	ScheduledDate string      `json:"flightScheduleDate"`
	Id            string      `json:"id"`
	CompleteRoute string      `json:"route"`
	Airline       Airline     `json:"airline"`
	Status        string      `json:"flightStatusPublic"`
	Legs          []FlightLeg `json:"flightLegs"`
}

type FlightLeg struct {
	Status          string                   `json:"status"`
	PublicStatus    string                   `json:"legStatusPublic"`
	IATAServiceType string                   `json:"serviceType"`
	Restricted      bool                     `json:"restricted"`
	Completion      string                   `json:"completionPercentage"`
	Departure       FlightStationInformation `json:"departureInformation"`
	Arrival         FlightStationInformation `json:"arrivalInformation"`
	Aircraft        Aircraft                 `json:"aircraft"`
	Irregulary      Irregularities           `json:"irregulary"`
}

type FlightStationInformation struct {
	Airport Airport       `json:"airport"`
	Times   FlightETAData `json:"times"`
}

type FlightETAData struct {
	Scheduled        time.Time `json:"scheduled"`
	LatestPublished  time.Time `json:"latestPublished"`
	EstimatedPublic  time.Time `json:"estimatedPublic"`
	EstimatesTakeOff time.Time `json:"estimatedTakeOffTime"`
}

type Airport struct {
	Name        string      `json:"name"`
	Code        string      `json:"code"`
	City        City        `json:"city"`
	Coordinates Coordinates `json:"location"`
}

type City struct {
	Name    string  `json:"name"`
	Code    string  `json:"code"`
	Country Country `json:"country"`
}

type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Coordinates struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type Place struct {
	Gates        []string `json:"gateNumber"`
	Parking      string   `json:"parkingPosition"`
	PierCode     string   `json:"pierCode"`
	Terminal     string   `json:"terminal"`
	BoardingPier string   `json:"boardingPier"`
}

type Aircraft struct {
	Registration     string `json:"registration"`
	TypeCode         string `json:"typeCode"`
	TypeName         string `json:"typeName"`
	OwnerAirlineCode string `json:"ownerAirlineCode"`
	OwnerAirlineName string `json:"ownerAirlineName"`
}

type Airline struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Irregularities struct {
	Cancelled string `json:"cancelled"`
}

type StationData struct {
	Station []Station `json:"stationCities"`
}

type Station struct {
	IATAStationCode string `json:"iataStationCode"`
	Name            string `json:"name"`
	IATACityCode    string `json:"iataCityCode"`
	City            string `json:"cityName"`
	ISOCountry      string `json:"iso2CountryCode"`
	Country         string `json:"countryName"`
}

/*
 {
      "flightLegs": [
        {

          "irregularity": {
            "cancelled": "N"
          },
        }
      ],
    },
*/
