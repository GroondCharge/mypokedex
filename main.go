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
	"mapb": {
		name:        "mapb",
		description: "Provides a list of location-area points on previous page",
		callback:    commandMapb,
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
			break
		}
		keyword := strings.Fields(strings.ToLower(scanner.Text()))[0]
		command, ok := commands[keyword]
		if !ok {
			fmt.Printf("command was not found")
			continue
		}
		err := command.callback(config_pointer)
		if err != nil {
			fmt.Println(err)
		}

		//continue
		//keyword := strings.Fields(strings.ToLower(scanner.Text()))[0]

	}
}
