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

func sendAllFlightsRequest() FlightsData {
	data := FlightsData{}
	json.Unmarshal(sendRequest("GET", "https://api.airfranceklm.com/opendata/flightstatus?serviceType=J", nil), &data)
	return data
}

func sendFlightRequest(id string) Flight {
	data := Flight{}
	json.Unmarshal(sendRequest("GET", fmt.Sprintf("https://api.airfranceklm.com/opendata/flightstatus/%s", id), nil), &data)

	return data
}

func sendRequest(method string, uri string, inBody interface{}) []byte {
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