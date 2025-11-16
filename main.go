package main

import (
	"bufio"
	"fmt"
	"os"
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
}

func main() {

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
		command.callback()

	}
}
