package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	Next     string              `json:"next"`
	Previous string              `json:"previous"`
	Results  []map[string]string `json:"results"`
}

func Populate_config(url string, myconf *Config) error {
	client := &http.Client{}
	//var results []map[string]string
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request?")
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("issue performing the request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("issue getting body from resp")
	}
	errorz := json.Unmarshal(body, &myconf)
	if errorz != nil {
		return fmt.Errorf("issue decod ing into myconf")
	}
	for i, _ := range myconf.Results {
		fmt.Println(myconf.Results[i]["name"])
	}
	//fmt.Println(myconf.Next, myconf.Previous)
	//fmt.Printf("%T, %T", myconf.Next, myconf.Previous)
	return nil
}
