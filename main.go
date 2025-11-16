package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		fmt.Print("Pokedex >")
		reader := bufio.NewReader(os.Stdin)
		scanner := bufio.NewScanner(reader)
		checker := scanner.Scan()
		if !checker {
			fmt.Printf("No more commands were given")
		}
		fmt.Printf("Your command was: %s\n", strings.Fields(strings.ToLower(scanner.Text()))[0])

	}
}
