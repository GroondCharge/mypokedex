package main

import (
	"fmt"
	"os"
	"pokedex/pokeapi"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error
}

func commandExit(config *pokeapi.Config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
func commandHelp(config *pokeapi.Config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}
func commandMap(config *pokeapi.Config) error {
	if config.Next == "" {
		pokeapi.Populate_config("https://pokeapi.co/api/v2/location-area/", config)
	} else {
		pokeapi.Populate_config(config.Next, config)
	}
	//fmt.Print(config.Next)
	//fmt.Print(config.Previous)
	return nil
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
