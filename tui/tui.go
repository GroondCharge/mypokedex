package tui

import (
	"math/rand"
	"pokedex/internal/pokeapi"

	tea "github.com/charmbracelet/bubbletea"
)

// viewState represents which screen is currently displayed
type viewState int

const (
	locationListView  viewState = iota // main: browsing location areas
	areaDetailView                      // viewing Pokemon encounters in a location
	pokemonDetailView                   // inspecting a caught Pokemon's stats
	pokedexListView                     // browsing your caught Pokemon collection
)

// Config mirrors the main package Config for TUI use
type Config struct {
	PokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	Inventory     map[string]pokeapi.PokemonStruct
}

// Model is the main Bubble Tea model
type Model struct {
	// Core
	cfg     *Config
	view    viewState
	loading bool
	err     error
	message string // status messages (catch result, etc.)

	// Location List View state
	locations []pokeapi.ResponseLocations
	locItems  []string // location names for display
	locCursor int

	// Area Detail View state
	currentAreaName string
	currentArea     *pokeapi.LocationArea
	areaCursor      int

	// Pokemon Detail View state
	currentPokemon *pokeapi.PokemonStruct

	// Pokedex View state
	pokedexCursor int
	pokedexNames  []string // cached list of caught Pokemon names

	// Dimensions
	width  int
	height int
}

// New creates a new TUI model with the given config
func New(cfg *Config) Model {
	if cfg.Inventory == nil {
		cfg.Inventory = make(map[string]pokeapi.PokemonStruct)
	}
	return Model{
		cfg:     cfg,
		view:    locationListView,
		loading: true,
	}
}

// Init implements tea.Model - fetches initial locations
func (m Model) Init() tea.Cmd {
	return m.fetchLocations(nil)
}

// Update implements tea.Model - handles messages and key events
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		// Clear any status message on keypress
		m.message = ""
		return m.handleKeyPress(msg)

	case locationsMsg:
		return m.handleLocationsMsg(msg)

	case areaMsg:
		return m.handleAreaMsg(msg)

	case pokemonMsg:
		return m.handlePokemonMsg(msg)

	case catchResultMsg:
		return m.handleCatchResult(msg)

	case errMsg:
		m.loading = false
		m.err = msg.err
		return m, nil
	}

	return m, nil
}

// View implements tea.Model - renders the current screen
func (m Model) View() string {
	if m.loading {
		return m.renderLoading()
	}
	if m.err != nil {
		return m.renderError()
	}

	switch m.view {
	case locationListView:
		return m.renderLocationList()
	case areaDetailView:
		return m.renderAreaDetail()
	case pokemonDetailView:
		return m.renderPokemonDetail()
	case pokedexListView:
		return m.renderPokedexList()
	}

	return "Unknown view"
}

// handleKeyPress routes key events based on current view
func (m Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Global keys
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "?":
		// TODO: toggle help overlay
		return m, nil
	}

	// View-specific keys
	switch m.view {
	case locationListView:
		return m.handleLocationListKeys(msg)
	case areaDetailView:
		return m.handleAreaDetailKeys(msg)
	case pokemonDetailView:
		return m.handlePokemonDetailKeys(msg)
	case pokedexListView:
		return m.handlePokedexListKeys(msg)
	}

	return m, nil
}

// --- Location List View ---

func (m Model) handleLocationListKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "j", "down":
		if m.locCursor < len(m.locItems)-1 {
			m.locCursor++
		}
	case "k", "up":
		if m.locCursor > 0 {
			m.locCursor--
		}
	case "n":
		if m.cfg.Next != nil {
			m.loading = true
			return m, m.fetchLocations(m.cfg.Next)
		}
	case "N":
		if m.cfg.Previous != nil {
			m.loading = true
			return m, m.fetchLocations(m.cfg.Previous)
		}
	case "enter":
		if len(m.locItems) > 0 {
			m.loading = true
			m.currentAreaName = m.locItems[m.locCursor]
			return m, m.fetchArea(m.currentAreaName)
		}
	case "p":
		m.view = pokedexListView
		m.pokedexCursor = 0
		m.updatePokedexNames()
	}
	return m, nil
}

// --- Area Detail View ---

func (m Model) handleAreaDetailKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "j", "down":
		if m.currentArea != nil && m.areaCursor < len(m.currentArea.PokemonEncounters)-1 {
			m.areaCursor++
		}
	case "k", "up":
		if m.areaCursor > 0 {
			m.areaCursor--
		}
	case "c":
		// Catch the selected Pokemon
		if m.currentArea != nil && len(m.currentArea.PokemonEncounters) > 0 {
			pokemonName := m.currentArea.PokemonEncounters[m.areaCursor].Pokemon.Name
			m.loading = true
			return m, m.catchPokemon(pokemonName)
		}
	case "enter":
		// View Pokemon details if caught
		if m.currentArea != nil && len(m.currentArea.PokemonEncounters) > 0 {
			pokemonName := m.currentArea.PokemonEncounters[m.areaCursor].Pokemon.Name
			if pokemon, ok := m.cfg.Inventory[pokemonName]; ok {
				m.currentPokemon = &pokemon
				m.view = pokemonDetailView
			}
		}
	case "esc", "backspace":
		m.view = locationListView
		m.areaCursor = 0
	case "p":
		m.view = pokedexListView
		m.pokedexCursor = 0
		m.updatePokedexNames()
	}
	return m, nil
}

// --- Pokemon Detail View ---

func (m Model) handlePokemonDetailKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc", "backspace":
		// Go back to previous view
		if m.currentArea != nil {
			m.view = areaDetailView
		} else {
			m.view = pokedexListView
		}
		m.currentPokemon = nil
	}
	return m, nil
}

// --- Pokedex List View ---

func (m Model) handlePokedexListKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "j", "down":
		if m.pokedexCursor < len(m.pokedexNames)-1 {
			m.pokedexCursor++
		}
	case "k", "up":
		if m.pokedexCursor > 0 {
			m.pokedexCursor--
		}
	case "enter":
		if len(m.pokedexNames) > 0 {
			name := m.pokedexNames[m.pokedexCursor]
			if pokemon, ok := m.cfg.Inventory[name]; ok {
				m.currentPokemon = &pokemon
				m.currentArea = nil // mark that we came from pokedex
				m.view = pokemonDetailView
			}
		}
	case "esc", "backspace", "p":
		m.view = locationListView
	}
	return m, nil
}

func (m *Model) updatePokedexNames() {
	m.pokedexNames = make([]string, 0, len(m.cfg.Inventory))
	for name := range m.cfg.Inventory {
		m.pokedexNames = append(m.pokedexNames, name)
	}
}

// --- Messages ---

type locationsMsg struct {
	response pokeapi.ResponseLocations
}

type areaMsg struct {
	area pokeapi.LocationArea
}

type pokemonMsg struct {
	pokemon pokeapi.PokemonStruct
}

type catchResultMsg struct {
	pokemon pokeapi.PokemonStruct
	caught  bool
}

type errMsg struct {
	err error
}

// --- Commands ---

func (m Model) fetchLocations(url *string) tea.Cmd {
	return func() tea.Msg {
		resp, err := m.cfg.PokeapiClient.ListLocations(url)
		if err != nil {
			return errMsg{err: err}
		}
		return locationsMsg{response: resp}
	}
}

func (m Model) fetchArea(name string) tea.Cmd {
	return func() tea.Msg {
		area, err := m.cfg.PokeapiClient.GetLocationArea(name)
		if err != nil {
			return errMsg{err: err}
		}
		return areaMsg{area: area}
	}
}

func (m Model) catchPokemon(name string) tea.Cmd {
	return func() tea.Msg {
		pokemon, err := m.cfg.PokeapiClient.GetPokemon(name)
		if err != nil {
			return errMsg{err: err}
		}

		// Catch logic (same as original)
		randomInt := rand.Intn(100)
		chanceOfCatching := float32(randomInt) / (float32(pokemon.BaseExperience) / 100)
		caught := chanceOfCatching > 50.0

		return catchResultMsg{pokemon: pokemon, caught: caught}
	}
}

// --- Message Handlers ---

func (m Model) handleLocationsMsg(msg locationsMsg) (tea.Model, tea.Cmd) {
	m.loading = false
	m.cfg.Next = msg.response.Next
	m.cfg.Previous = msg.response.Previous

	m.locItems = make([]string, len(msg.response.Results))
	for i, loc := range msg.response.Results {
		m.locItems[i] = loc.Name
	}
	m.locCursor = 0

	return m, nil
}

func (m Model) handleAreaMsg(msg areaMsg) (tea.Model, tea.Cmd) {
	m.loading = false
	m.currentArea = &msg.area
	m.areaCursor = 0
	m.view = areaDetailView
	return m, nil
}

func (m Model) handlePokemonMsg(msg pokemonMsg) (tea.Model, tea.Cmd) {
	m.loading = false
	m.currentPokemon = &msg.pokemon
	m.view = pokemonDetailView
	return m, nil
}

func (m Model) handleCatchResult(msg catchResultMsg) (tea.Model, tea.Cmd) {
	m.loading = false

	if _, ok := m.cfg.Inventory[msg.pokemon.Name]; ok {
		m.message = msg.pokemon.Name + " is already in your Pokedex!"
		return m, nil
	}

	if msg.caught {
		m.cfg.Inventory[msg.pokemon.Name] = msg.pokemon
		m.message = "Caught " + msg.pokemon.Name + "!"
	} else {
		m.message = msg.pokemon.Name + " escaped!"
	}

	return m, nil
}
