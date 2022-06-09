package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

// curl -X GET "network-and-schedule/typical-flight-schedule/iata-seasons" -H "Api-Key:****"
const (
	NetworkSchedule = "network-and-schedule"
)

func sendAllFlightsRequest(begin string, end string, pageNumber int) FlightsData {
	data := FlightsData{}
	json.Unmarshal(sendRequest("GET", fmt.Sprintf("https://api.airfranceklm.com/opendata/flightstatus?serviceType=J&startRange=%s&endRange=%s&pageSize=12&origin=CDG&departureCity=&pageNumber=%d", begin, end, pageNumber), nil, nil), &data)
	return data
}

func sendOfferDetailRequest(plannedDeparture string, from string, to string) Itineraries {
	result := Itineraries{}

	offerHeaders := make(map[string]string)
	offerHeaders["AFKL-TRAVEL-Host"] = "AF"
	offerHeaders["AFKL-Travel-Country"] = "FR"
	offerHeaders["Content-Type"] = "application/json"
	offerHeaders["Accept"] = "application/hal+json"

	aop := AvailableOfferParam{}
	aop.CommercialCabin = []string{"ALL"}
	aop.PassengerCount = PassengerCount{
		Adult: 1,
	}
	aop.RequestedConnections = []RequestedConnectionParam{{
		DepartureDate: plannedDeparture, // "2022-07-15",
		Origin:        RequestedParam{Airport: LocationParam{Code: from}},
		Destination:   RequestedParam{Airport: LocationParam{Code: to}},
	}}

	json.Unmarshal(sendRequest("POST", "https://api.airfranceklm.com/opendata/offers/v1/available-offers", aop, offerHeaders), &result)

	return result
}

func sendOfferRequest(plannedDeparture string, from string, to string) Itineraries {
	result := Itineraries{}

	offerHeaders := make(map[string]string)
	offerHeaders["AFKL-TRAVEL-Host"] = "AF"
	offerHeaders["AFKL-Travel-Country"] = "FR"
	offerHeaders["Content-Type"] = "application/json"
	offerHeaders["Accept"] = "application/hal+json"

	aop := AvailableOfferParam{}
	aop.CommercialCabin = []string{"ALL"}
	aop.PassengerCount = PassengerCount{
		Adult: 1,
	}
	aop.RequestedConnections = []RequestedConnectionParam{{
		DepartureDate: plannedDeparture, // "2022-07-15",
		Origin:        RequestedParam{Airport: LocationParam{Code: from}},
		Destination:   RequestedParam{Airport: LocationParam{Code: to}},
	}}

	json.Unmarshal(sendRequest("POST", "https://api.airfranceklm.com/opendata/offers/v1/available-offers", aop, offerHeaders), &result)

	return result
}

func sendFlightRequest(id string) Flight {
	data := Flight{}
	headers := make(map[string]string)
	headers["Accept"] = "application/hal+json"
	json.Unmarshal(sendRequest("GET", fmt.Sprintf("https://api.airfranceklm.com/opendata/flightstatus/%s", id), nil, headers), &data)
	return data
}

func sendStationsRequest() StationCitiesResponse {
	data := StationCitiesResponse{}
	json.Unmarshal(sendRequest("GET", "https://api.airfranceklm.com/opendata/network-and-schedule/typical-flight-schedule/stations", nil, nil), &data)
	return data
}

func sendOfferDetailFromLink(input string, dataType string) interface{} {

	offerHeaders := make(map[string]string)
	offerHeaders["AFKL-TRAVEL-Host"] = "AF"
	offerHeaders["AFKL-TRAVEL-Country"] = "FR"
	offerHeaders["Content-Type"] = "application/json"
	offerHeaders["Accept"] = "application/hal+json"

	var data = reflect.New(RequestRedirect[dataType]).Interface()
	json.Unmarshal(sendRequest("GET", input, nil, offerHeaders), &data)
	return data

}

func sendRequest(method string, uri string, inBody interface{}, headers map[string]string) []byte {
	client := &http.Client{}

	var req *http.Request
	if inBody != nil {
		bodyByte, _ := json.Marshal(inBody)
		fmt.Println(string(bodyByte))
		req, _ = http.NewRequest(method, uri, bytes.NewBuffer(bodyByte))
	} else {
		req, _ = http.NewRequest(method, uri, nil)
	}
	req.Header.Add("Api-Key", appConfig.Key)
	req.Header.Add("Accept-Language", "en-GB")

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Status)
	fmt.Println(res.StatusCode)

	return body
}
