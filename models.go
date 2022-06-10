package main

import (
	"time"
)

// Get all flights
type FlightsData struct {
	Flights  []Flight `json:"operationalFlights"`
	PageData PageData `json:"page"`
}

type PageData struct {
	Size        int `json:"pageSize"`
	Number      int `json:"pageNumber"`
	FullCount   int `json:"fullCount"`
	ItemPerPage int `json:"pageCount"`
	TotalPages  int `json:"totalPages"`
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

// Find fare
//   "required" : [ "bookingFlow", "commercialCabins", "customer", "focusedConnection", "passengers", "reservationId" ],
type AvailableOfferParam struct {
	//BookignFlow          string                     `json:"bookingFlow"` // [ "REWARD", "CORPORATE", "LEISURE", "STAFF" ]
	//Passengers           []PassengerParam           `json:"passengers"`
	PassengerCount  PassengerCount `json:"passengerCount"`
	CommercialCabin []string       `json:"commercialCabins"` // [ "ECONOMY", "PREMIUM", "BUSINESS", "FIRST", "ALL" ]
	//FocusedConnection    int                        `json:"focusedConnection"`
	RequestedConnections []RequestedConnectionParam `json:"requestedConnections"`
}

const ()

type PassengerCount struct {
	Adult    int `json:"ADT"`
	Children int `json:"CHD"`
	Infants  int `json:"INF"`
}

type PassengerParam struct {
	Id           int    `json:"id"`
	Type         string `json:"type"`
	PNRId        int    `json:"pnrId"`
	BirthDate    string `json:"birthDate"`
	MinAge       int    `json:"minAge"`
	MaxAge       int    `json:"maxAge"`
	TicketNumber string `json:"ticketNumber"`
}

type RequestedConnectionParam struct {
	DepartureDate string         `json:"departureDate"`
	Origin        RequestedParam `json:"origin"`
	Destination   RequestedParam `json:"destination"`
}

type RequestedParam struct {
	Airport LocationParam `json:"airport"`
}

type LocationParam struct {
	//Type string `json:"type"`
	Code string `json:"code"`
}

// Available offer response
type Itineraries struct {
	FlightProduct  []Itinerary `json:"itineraries"`
	TotalPriceText string      `json:"totalPriceText"`
}

type Itinerary struct {
	FlightProducts []FlightProduct    `json:"flightProducts"`
	Connections    []ConnectionDetail `json:"connections"`
}

type FlightProduct struct {
	Passengers    []PassengerParam `json:"passengers"`
	PriceResponse PriceResponse    `json:"price"`
	Connection    []Connection     `json:"connections"`
	Links         Links            `json:"_links"`
}

type PriceResponse struct {
	DisplayPrice      float32             `json:"displayPrice"`
	TotalPrice        float32             `json:"totalPrice"`
	PricePerPassenger []PricePerPassenger `json:"pricePerPassengerTypes"`
	Flexibility       bool                `json:"flexibilityWaiver"`
	Currency          string              `json:"currency"`
	DisplayType       string              `json:"displayType"`
}

type PricePerPassenger struct {
	PassengerType string           `json:"passengerType"`
	Fare          float32          `json:"fare"`
	Taxes         float32          `json:"taxes"`
	Products      int              `json:"products"`
	PrimaryPax    bool             `json:"primaryPax"`
	Surcharge     []SurchargePrice `json:"surcharges"`
}

type SurchargePrice struct {
	Amount float32 `json:"amount"`
	Code   string  `json:"code"`
}

type Connection struct {
	AvailableSeats       int           `json:"numberOfSeatsAvailable"`
	FareBasis            FareBasis     `json:"fareBasis"`
	Segments             []Segment     `json:"segments"`
	FareFamily           FareFamily    `json:"fareFamily"`
	CommercialCabin      string        `json:"commercialCabin"`
	CommercialCabinLabel string        `json:"commercialCabinLabel"`
	PriceResponse        PriceResponse `json:"price"`
	Links                Links         `json:"_links"`
}

type FareBasis struct {
	Code string `json:"code"`
}

type Links struct {
	FlightDetails    Link `json:"flightDetails"`
	TicketConditions Link `json:"ticketConditions"`
	ShoppingCart     Link `json:"shoppingCart"`
	TaxBreakdown     Link `json:"taxBreakdown"`
	PriceDetails     Link `json:"priceDetails"`
	UpsellOffers     Link `json:"upsellOffers"`
	RelatedProducts  Link `json:"relatedProducts"`
	Information      Link `json:"information"`
}

type Link struct {
	URL         string `json:"href"`
	Templated   bool   `json:"templated"`
	UseRootPath bool   `json:"useRootPath"`
}

type FareFamily struct {
	Code      string `json:"code"`
	Hierarchy int    `json:"hierarchy"`
}

type SellingClass struct {
	Code string `json:"code"`
}

type Cabin struct {
	Class string `json:"class"`
}

type Segment struct {
	Cabin        Cabin        `json:"cabin"`
	FareBasis    FareBasis    `json:"fareBasis"`
	SellingClass SellingClass `json:"sellingClass"`
}

type ConnectionDetail struct {
	Duration      int             `json:"duration"`
	SegmentDetail []SegmentDetail `json:"segments"`
}

type SegmentDetail struct {
	ArrivalDateTime   string          `json:"arrivalDateTime"`
	DepartureDateTime string          `json:"departureDateTime"`
	Destination       Airport         `json:"destination"`
	Origin            Airport         `json:"origin"`
	HighestPriority   bool            `json:"highestPriority"`
	FlightDuration    int             `json:"flightDuration"`
	DateVariation     int             `json:"dateVariation"`
	MarketingFlight   MarketingFlight `json:"marketingFlight"`
}

type MarketingFlight struct {
	Number  string  `json:"number"`
	Carrier Carrier `json:"carrier"`
}

type OperatingFlight struct {
	Number  string  `json:"number"`
	Carrier Carrier `json:"carrier"`
}

type EquipmentType struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	ACVCode string `json:"acvCode"`
	Links   Links  `json:"_links"`
}

type Carrier struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type StationCitiesResponse struct {
	StationCities []StationCity `json:"stationCities"`
}

type StationCity struct {
	IataStationCode string `json:"iataStationCode"`
	Name            string `json:"name"`
	IataCityCode    string `json:"iataCityCode"`
	CityName        string `json:"cityName"`
	Iso2CountryCode string `json:"iso2CountryCode"`
	CountryName     string `json:"countryName"`
	SubRegionCode   string `json:"subRegionCode"`
	RegionCode      string `json:"regionCode"`
	SubRegionName   string `json:"subRegionName"`
	RegionName      string `json:"regionName"`
}
