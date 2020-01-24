package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func makeRequest(url string) *http.Response {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return response
}

func closeResponse(response *http.Response) {
	if err := response.Body.Close(); err != nil {
		panic(err)
	}
}

func getBody(response *http.Response) []byte {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func getJSON(body []byte) *Response {
	data := Response{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	return &data
}

func getUrl(url string) *Response {
	response := makeRequest(url)
	defer closeResponse(response)
	body := getBody(response)
	data := getJSON(body)
	return data
}
