package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
