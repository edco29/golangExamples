package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	SONAR_URL    = "https://sonarcloud.io/api/projects/search"
	TOKEN        = ""
	ORGANIZATION = ""
)

var bearer = "Bearer " + TOKEN

type Pagging struct {
	PageIndex int64 `json:"pageIndex"`
	PageSize  int64 `json:"pageSize"`
	Total     int64 `json:"total"`
}

type Component struct {
	Organization string `json:"organization"`
	Key          string `json:"key"`
	Name         string `json:"name"`
	Qualifier    string `json:"qualifier"`
	Visibility   string `json:"visibility"`
}

type Components struct {
	Paging     Pagging
	Components []Component
}

func main() {

	// Create a new request using http
	req, err := http.NewRequest("GET", SONAR_URL, http.NoBody)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// pass query params
	q := req.URL.Query()
	q.Add("organization", ORGANIZATION)
	req.URL.RawQuery = q.Encode()

	// Send req using http Client
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	var c1 Components
	json.Unmarshal(body, &c1)
	fmt.Print(c1.Components)
}
