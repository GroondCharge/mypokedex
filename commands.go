package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"pokedex/internal/pokeapi"
)

func commandExit(config *Config, args ...string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
func commandHelp(config *Config, args ...string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}
func commandMap(config *Config, args ...string) error {
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
func commandMapb(config *Config, args ...string) error {
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

func commandExplore(config *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("please provide a location area name\n")
	}
	areaName := args[0]
	fmt.Printf("Exploring %s...\n", areaName)

	locationArea, err := config.pokeapiClient.GetLocationArea(areaName)
	if err != nil {
		return err
	}
	fmt.Print("Found Pokemon:\n")
	for _, eachEncounter := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", eachEncounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(config *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("please provide a name of the pokemonn you're trying to catch\n")
	}
	if len(config.Invetory) == 0 {
		config.Invetory = map[string]pokeapi.PokemonStruct{}
	}
	pokemonName := args[0]
	pokemonObject, err := config.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	if _, ok := config.Invetory[pokemonName]; ok {
		fmt.Printf("Pokemon %s already in inventory, continuing\n", pokemonName)
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	randomInt := rand.Intn(100)
	var chancerOfCatching float32
	chancerOfCatching = (float32(randomInt)) / (float32(pokemonObject.BaseExperience) / 100)
	if chancerOfCatching > 50.0 {
		fmt.Printf("Caught %s, adding to pokedex\n", pokemonName)
		config.Invetory[pokemonName] = pokemonObject
	} else {
		fmt.Printf("Unlucky, did not catch %s, better luck next time\n", pokemonName)
	}
	return nil
}

func commandInspect(config *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("provide a name of the pokemon")
	}
	if pokemon, ok := config.Invetory[args[0]]; ok {
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, eachType := range pokemon.Types {
			fmt.Printf("- %s\n", eachType.Type.Name)
		}
	} else {
		return fmt.Errorf("You haven't caught this Pokemon yet")
	}
	return nil
}
