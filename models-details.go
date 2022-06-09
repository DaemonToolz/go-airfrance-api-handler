package main

type InputUrl struct {
	URL     string `json:"url"`
	ReqType string `json:"type"`
}

type TaxBreakdownDetailResult struct {
	Price TaxPriceDetail `json:"price"`
}

type TaxPriceDetail struct {
	TotalPrice                                   float32                    `json:"totalPrice"`
	TotalFare                                    float32                    `json:"totalFare"`
	TotalTaxes                                   float32                    `json:"totalTaxes"`
	TotalProducts                                int                        `json:"totalProducts"`
	TotalTaxesWithoutAirTransportationSurcharges float32                    `json:"totalTaxesWithoutAirTransportationSurcharges"`
	TotalSurcharges                              []TaxPriceDetailExtract    `json:"totalSurcharges"`
	TotalAirTransportationSurcharges             []TaxPriceDetailExtract    `json:"totalAirTransportationSurcharges"`
	PricePerPassengerTypes                       []TaxPricePerPassengerType `json:"pricePerPassengerTypes"`
	Currency                                     string                     `json:"currency"`
}

type TaxPricePerPassengerType struct {
	Products                    int                     `json:"products"`
	PassengerType               string                  `json:"passengerType"`
	Fare                        float32                 `json:"fare"`
	FareDetails                 []TaxPriceDetailExtract `json:"fareDetails"`
	Taxes                       []TaxPriceDetailExtract `json:"taxes"`
	Surcharges                  []TaxPriceDetailExtract `json:"surcharges"`
	AirTransportationSurcharges TaxPriceDetailExtract   `json:"airTransportationSurcharges"`
}

type TaxPriceDetailExtract struct {
	Code   string  `json:"code"`
	Amount float32 `json:"amount"`
	Name   string  `json:"name"`
}

type PriceDetailResult struct {
	PriceDetails PriceDetail `json:"priceDetails"`
}

type PriceDetail struct {
	Totals          []PriceTotal       `json:"totals"`
	PriceCategories []PricePerCategory `json:"priceCategories"`
}

type PriceTotal struct {
	Currency           string                    `json:"currency"`
	Amount             float32                   `json:"amount"`
	PricePerPassengers []PricePerPassengerDetail `json:"pricePerPassengers"`
}

type PricePerPassengerDetail struct {
	PassengerID int     `json:"passengerId"`
	Amount      float32 `json:"amount"`
}

type PricePerCategory struct {
	Code            string           `json:"code"`
	Label           string           `json:"label"`
	Totals          []PriceTotal     `json:"totals"`
	PriceComponents []PriceComponent `json:"priceComponents"`
}

type PriceComponent struct {
	Code                    string                    `json:"code"`
	Label                   string                    `json:"label"`
	Currency                string                    `json:"currency"`
	Amount                  float32                   `json:"amount"`
	PricePerPassengers      []PricePerPassengerDetail `json:"pricePerPassengers"`
	PriceComponentBreakdown []PriceComponentBreakdown `json:"priceComponentBreakdown,omitempty"`
}

type PriceComponentBreakdown struct {
	Code               string                    `json:"code"`
	Nature             string                    `json:"nature"`
	Label              string                    `json:"label"`
	Amount             float32                   `json:"amount"`
	PricePerPassengers []PricePerPassengerDetail `json:"pricePerPassengers"`
}

type TicketConditionsResponse struct {
	Passengers          []PassengerTicketCondition  `json:"passengers"`
	SameConditions      bool                        `json:"sameConditions"`
	Connections         []ConnectionTicketCondition `json:"connections"`
	BookingDateInterval BookingIntervalCondition    `json:"bookingDateInterval"`
	FlyingBlue          FlyingBlueTicketCondition   `json:"flyingBlue"`
	Disclaimer          TicketConditionDisclaimer   `json:"disclaimer"`
}

type TicketConditionDisclaimer struct {
	HandBaggageText string `json:"handBaggageText"`
	FlexibilityText string `json:"flexibilityText"`
}

type BookingIntervalCondition struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type TicketCommercialText struct {
	Relevance int    `json:"relevance"`
	Text      string `json:"text"`
}

type FlyingBlueTicketCondition struct {
	Earned                   int                      `json:"earned"`
	DetailsText              string                   `json:"detailsText"`
	QualifyingPoints         int                      `json:"qualifyingPoints"`
	UltimateQualifyingPoints int                      `json:"ultimateQualifyingPoints"`
	DetailsTextQualPoints    string                   `json:"detailsTextQualPoints"`
	CommercialText           PassengerTicketCondition `json:"commercialText"`
}

type PassengerTicketCondition struct {
	ID               int    `json:"id"`
	Type             string `json:"type"`
	PrimaryPassenger bool   `json:"primaryPassenger"`
}

type FarePercentage struct {
	FarePercentage int `json:"farePercentage"`
}

type SkyPriorityCondition struct {
	CommercialText TicketCommercialText `json:"commercialText"`
	DetailsText    string               `json:"detailsText"`
	Allowed        bool                 `json:"allowed"`
}

type AdvancedPurchaseCondition struct {
	StayDuration StayDuration `json:"stayDuration"`
	DetailsText  string       `json:"detailsText"`
}

type StayDuration struct {
	Duration int    `json:"duration"`
	StayUnit string `json:"stayUnit"`
}

type CancelCondition struct {
	AllowedBeforeDeparture     bool                 `json:"allowedBeforeDeparture"`
	AllowedAfterDeparture      bool                 `json:"allowedAfterDeparture"`
	CommercialText             TicketCommercialText `json:"commercialText"`
	AfterDepartureDetailsText  string               `json:"afterDepartureDetailsText"`
	BeforeDepartureDetailsText string               `json:"beforeDepartureDetailsText"`
	ExtraCostsText             string               `json:"extraCostsText"`
}

type ChangeCondition struct {
	AllowedBeforeDeparture     bool                 `json:"allowedBeforeDeparture"`
	AllowedAfterDeparture      bool                 `json:"allowedAfterDeparture"`
	CommercialText             TicketCommercialText `json:"commercialText"`
	AfterDepartureDetailsText  string               `json:"afterDepartureDetailsText"`
	BeforeDepartureDetailsText string               `json:"beforeDepartureDetailsText"`
}

type NoShowCondition struct {
	CommercialText TicketCommercialText `json:"commercialText"`
	DetailsText    string               `json:"detailsText"`
	Allowed        bool                 `json:"allowed"`
}

type ConditionPerPassengerDetail struct {
	InfantDiscountCondition   FarePercentage            `json:"infantDiscountCondition,omitempty"`
	ChildrenDiscountCondition FarePercentage            `json:"childrenDiscountCondition,omitempty"`
	SkyPriorityCondition      SkyPriorityCondition      `json:"skyPriorityCondition,omitempty"`
	AdvancePurchaseCondition  AdvancedPurchaseCondition `json:"advancePurchaseCondition,omitempty"`
	CancelCondition           CancelCondition           `json:"cancelCondition,omitempty"`
	ChangeCondition           ChangeCondition           `json:"changeCondition,omitempty"`
	NoShowCondition           NoShowCondition           `json:"noShowCondition,omitempty"`
}

type FareFamilyCondition struct {
	Code                  string `json:"code"`
	Hierarchy             int    `json:"hierarchy"`
	Title                 string `json:"title"`
	CommercialDescription string `json:"commercialDescription"`
}

type BaggageAllowance struct {
	Quantity               int                  `json:"quantity"`
	Type                   string               `json:"type"`
	CommercialText         TicketCommercialText `json:"commercialText"`
	DetailsText            string               `json:"detailsText"`
	HandBaggageDetailsText string               `json:"handBaggageDetailsText"`
}

type FlyingBlueDetail struct {
	Earned                   int                  `json:"earned"`
	DetailsText              string               `json:"detailsText"`
	QualifyingPoints         int                  `json:"qualifyingPoints"`
	UltimateQualifyingPoints int                  `json:"ultimateQualifyingPoints"`
	DetailsTextQualPoints    string               `json:"detailsTextQualPoints"`
	CommercialText           TicketCommercialText `json:"commercialText"`
}

type ConditionPerPassenger struct {
	PassengerID      int                           `json:"passengerId"`
	Conditions       []ConditionPerPassengerDetail `json:"conditions"`
	FareFamily       FareFamilyCondition           `json:"fareFamily"`
	BaggageAllowance BaggageAllowance              `json:"baggageAllowance"`
	FlyingBlue       FlyingBlueDetail              `json:"flyingBlue"`
	TravelClassText  string                        `json:"travelClassText"`
	TripText         string                        `json:"tripText"`
}

type ConnectionTicketCondition struct {
	TravelDateIntervals    []BookingIntervalCondition    `json:"travelDateIntervals"`
	ConditionsPerPassenger []ConditionPerPassenger       `json:"conditionsPerPassenger"`
	Conditions             []ConditionPerPassengerDetail `json:"conditions"`
	FareFamily             FareFamilyCondition           `json:"fareFamily"`
	BaggageAllowance       BaggageAllowance              `json:"baggageAllowance"`
	FlyingBlue             FlyingBlueDetail              `json:"flyingBlue"`
	TravelClassText        string                        `json:"travelClassText"`
	TripText               string                        `json:"tripText"`
}
