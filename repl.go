package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
type config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

func testing_api() {
	myconf := config{}
	client := &http.Client{}
	//req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location-area", nil)
	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location-area?offset=20&limit=20", nil)
	if err != nil {
		fmt.Print("Error creating request?")
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("issue performing the request")
	}
	fmt.Println()
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("issue getting body from resp")
	}
	errorz := json.Unmarshal(body, &myconf)
	if errorz != nil {
		fmt.Print("issue decod ing into myconf")
	}
	fmt.Println(myconf.Next, myconf.Previous)

}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
