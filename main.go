package main

import (
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"pokedex/tui"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 60*time.Second)

	// Create TUI config
	cfg := &tui.Config{
		PokeapiClient: pokeClient,
		Inventory:     make(map[string]pokeapi.PokemonStruct),
	}

	// Start Bubble Tea TUI
	p := tea.NewProgram(tui.New(cfg), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running TUI: %v\n", err)
		os.Exit(1)
	}
}
