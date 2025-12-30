package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 200*time.Second)
	cfg := &Config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
