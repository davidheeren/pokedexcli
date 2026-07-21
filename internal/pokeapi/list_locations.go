package pokeapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) ListLocations(apiUrl string) (MapLocations, error) {
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return MapLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return MapLocations{}, err
	}
	defer res.Body.Close()

	var locations MapLocations
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return MapLocations{}, err
	}

	return locations, nil
}

func LocationsPageNumber(apiUrl string) (int, error) {
	offsetUrl, err := url.Parse(apiUrl)
	if err != nil {
		return 0, err
	}

	offsetStr := offsetUrl.Query().Get("offset")
	if offsetStr == "" {
		return 0, errors.New("url has no query prameter 'offset'")
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return 0, err
	}

	pageNum := offset / 20 + 1
	return pageNum, nil
}
