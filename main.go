package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/pokeapi"
	"strings"
)

var commands = map[string]cliCommand{
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
}

func main() {

	config := pokeapi.Config{}
	config_pointer := &config
	for {
		fmt.Print("Pokedex >")
		reader := bufio.NewReader(os.Stdin)
		scanner := bufio.NewScanner(reader)
		checker := scanner.Scan()
		if !checker {
			fmt.Printf("No more commands were given")
		}
		keyword := strings.Fields(strings.ToLower(scanner.Text()))[0]
		command, ok := commands[keyword]
		if !ok {
			fmt.Printf("command was not found")
			continue
		}
		command.callback(config_pointer)

	}
}
