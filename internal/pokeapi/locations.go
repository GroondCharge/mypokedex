package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(areaName string) (LocationArea, error) {
	url := baseURL + "location-area/" + areaName

	// Check cache first
	if data, ok := c.cache.Get(url); ok {
		fmt.Println("Found in cache!")
		locationArea := LocationArea{}
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}
	fmt.Println("Not found in Cache - fetching data")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}
	response, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return LocationArea{}, fmt.Errorf("location area '%s' not found", areaName)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationArea{}, err
	}

	// Add to cache
	c.cache.Add(url, data)

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}

func (c *Client) ListLocations(pageURL *string) (ResponseLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check cache first
	if data, ok := c.cache.Get(url); ok {
		fmt.Println("Cache hit!")
		locationsResponse := ResponseLocations{}
		err := json.Unmarshal(data, &locationsResponse)
		if err != nil {
			return ResponseLocations{}, err
		}
		return locationsResponse, nil
	}
	fmt.Println("Cache miss - fetching from API...")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocations{}, err
	}
	response, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseLocations{}, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return ResponseLocations{}, err
	}

	// Add to cache
	c.cache.Add(url, data)

	locationsResponse := ResponseLocations{}
	err = json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return ResponseLocations{}, err
	}
	return locationsResponse, nil
}
