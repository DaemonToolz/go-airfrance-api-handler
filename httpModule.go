package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func sendOfferRequest() string {
	return ""
	//offerHeaders := make(map[string]string)
	//offerHeaders["AFKL-TRAVEL-Host"] = "AF"
	//fmt.Println(string(sendRequest("GET", "https://api.airfranceklm.com/opendata/offers/v1/available-offers?departureDate=2022-06-30T00:00:00Z&d='1'&displayPriceContent='1'&displayPriceBalance=true", nil, offerHeaders)))
	//return string(sendRequest("GET", "https://api.airfranceklm.com/opendata/offers/v1/available-offers?departureDate=2022-06-30T00:00:00Z&d='1'&displayPriceContent='1'&displayPriceBalance=true", nil, offerHeaders))
}

func sendFlightRequest(id string) Flight {
	data := Flight{}
	json.Unmarshal(sendRequest("GET", fmt.Sprintf("https://api.airfranceklm.com/opendata/flightstatus/%s", id), nil, nil), &data)
	return data
}

func sendRequest(method string, uri string, inBody interface{}, headers map[string]string) []byte {
	client := &http.Client{}

	var req *http.Request
	if inBody != nil {
		bodyByte, _ := json.Marshal(inBody)
		req, _ = http.NewRequest(method, uri, bytes.NewBuffer(bodyByte))
	} else {
		req, _ = http.NewRequest(method, uri, nil)
	}
	req.Header.Add("Accept", "application/hal+json")
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

	return body
}
