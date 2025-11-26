package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(config *Config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
func commandHelp(config *Config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}
func commandMap(config *Config) error {
	locationsResponse, err := config.pokeapiClient.ListLocations(config.Next)
	if err != nil {
		return err
	}
	config.Next = locationsResponse.Next
	config.Previous = locationsResponse.Previous
	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}
func commandMapb(config *Config) error {
	if config.Previous == nil {
		return errors.New("you're on the first page")
	}
	locationsResponse, err := config.pokeapiClient.ListLocations(config.Previous)
	if err != nil {
		return err
	}
	config.Next = locationsResponse.Next
	config.Previous = locationsResponse.Previous
	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}
