package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type Config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	Invetory      map[string]pokeapi.PokemonStruct
}

func startRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex>")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg, words[1:]...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("Unknown command, here are some of possible uses:\n")
			commandHelp(cfg)
			continue
		}
	}
}
func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Prints list of commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Provides a list of location-area points",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Provides a list of location-area points on previous page",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists Pokemon found in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Used for catching pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Used for inspecting caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Used for listing contents of your Pokedex",
			callback:    commandPokedex,
		},
	}
}
