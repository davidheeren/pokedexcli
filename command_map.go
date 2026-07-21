package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type mapQuery struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []mapLocation `json:"results"`
}

type mapLocation struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

var mapNextUrl string = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
var mapPrevUrl string = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"

func commandMap() error {
	err := runMap(mapNextUrl)
	return err
}

func commandMapB() error {
	err := runMap(mapPrevUrl)
	return err
}

func runMap(apiUrl string) error {
	req, err := http.NewRequest("GET", apiUrl, nil) 
	client := &http.Client {
		Timeout: time.Millisecond * 5000,
	}

	// res, err := http.Get(apiUrl)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var query mapQuery
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&query)
	if err != nil {
		return err
	}

	offsetUrl, err := url.Parse(apiUrl)
	if err != nil {
		return err
	}

	offsetStr := offsetUrl.Query().Get("offset")
	if offsetStr == "" {
		return errors.New("url has no query prameter 'offset'")
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return err
	}

	fmt.Printf("----PAGE %d----\n", offset/20+1)

	for _, location := range query.Results {
		fmt.Println(location.Name)
	}

	if query.Next != "" {
		mapNextUrl = query.Next
	}
	if query.Previous != "" {
		mapPrevUrl = query.Previous
	}
	return nil
}
