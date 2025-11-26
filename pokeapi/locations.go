package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (ResponseLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
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
	locationsResponse := ResponseLocations{}
	err = json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return ResponseLocations{}, err
	}
	return locationsResponse, nil
}
