package tui

import (
	"fmt"
	"strings"
)

func (m Model) renderLoading() string {
	return titleStyle.Render("Loading...")
}

func (m Model) renderError() string {
	return errorStyle.Render(fmt.Sprintf("Error: %v\n\nPress any key to continue", m.err))
}

func (m Model) renderLocationList() string {
	var b strings.Builder

	// Title
	b.WriteString(titleStyle.Render("Location Areas"))
	b.WriteString("\n")
	b.WriteString(subtitleStyle.Render("Select a location to explore"))
	b.WriteString("\n\n")

	// Location list
	for i, name := range m.locItems {
		cursor := "  "
		style := normalItemStyle
		if i == m.locCursor {
			cursor = "> "
			style = selectedStyle
		}
		b.WriteString(cursor + style.Render(name) + "\n")
	}

	// Status message
	if m.message != "" {
		b.WriteString(messageStyle.Render(m.message))
		b.WriteString("\n")
	}

	// Pagination info
	b.WriteString("\n")
	pageInfo := ""
	if m.cfg.Previous != nil {
		pageInfo += "[N] prev  "
	}
	if m.cfg.Next != nil {
		pageInfo += "[n] next"
	}
	if pageInfo != "" {
		b.WriteString(subtitleStyle.Render(pageInfo))
		b.WriteString("\n")
	}

	// Help
	b.WriteString(helpStyle.Render(locationListHelp))

	return b.String()
}

func (m Model) renderAreaDetail() string {
	var b strings.Builder

	// Title
	title := fmt.Sprintf("Exploring: %s", m.currentAreaName)
	b.WriteString(titleStyle.Render(title))
	b.WriteString("\n")
	b.WriteString(subtitleStyle.Render("Pokemon found in this area"))
	b.WriteString("\n\n")

	if m.currentArea == nil || len(m.currentArea.PokemonEncounters) == 0 {
		b.WriteString(subtitleStyle.Render("No Pokemon found here"))
		b.WriteString("\n")
	} else {
		for i, encounter := range m.currentArea.PokemonEncounters {
			name := encounter.Pokemon.Name
			cursor := "  "
			style := normalItemStyle
			if i == m.areaCursor {
				cursor = "> "
				style = selectedStyle
			}

			// Check if caught
			caught := ""
			if _, ok := m.cfg.Inventory[name]; ok {
				caught = caughtBadgeStyle.Render(" [caught]")
			}

			b.WriteString(cursor + style.Render(name) + caught + "\n")
		}
	}

	// Status message
	if m.message != "" {
		b.WriteString("\n")
		b.WriteString(messageStyle.Render(m.message))
	}

	// Help
	b.WriteString("\n")
	b.WriteString(helpStyle.Render(areaDetailHelp))

	return b.String()
}

func (m Model) renderPokemonDetail() string {
	if m.currentPokemon == nil {
		return "No Pokemon selected"
	}

	var b strings.Builder
	p := m.currentPokemon

	// Title
	b.WriteString(titleStyle.Render(strings.ToUpper(p.Name)))
	b.WriteString("\n\n")

	// Types
	b.WriteString("Types: ")
	for i, t := range p.Types {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(getTypeStyle(t.Type.Name).Render(t.Type.Name))
	}
	b.WriteString("\n\n")

	// Basic info
	b.WriteString(statLabelStyle.Render("Height:"))
	b.WriteString(statValueStyle.Render(fmt.Sprintf("%d", p.Height)))
	b.WriteString("\n")
	b.WriteString(statLabelStyle.Render("Weight:"))
	b.WriteString(statValueStyle.Render(fmt.Sprintf("%d", p.Weight)))
	b.WriteString("\n\n")

	// Stats
	b.WriteString(subtitleStyle.Render("Stats"))
	b.WriteString("\n")
	for _, stat := range p.Stats {
		label := statLabelStyle.Render(stat.Stat.Name + ":")
		value := statValueStyle.Render(fmt.Sprintf("%d", stat.BaseStat))
		bar := renderStatBar(stat.BaseStat)
		b.WriteString(fmt.Sprintf("%s %s %s\n", label, value, bar))
	}

	// Help
	b.WriteString("\n")
	b.WriteString(helpStyle.Render(pokemonDetailHelp))

	return boxStyle.Render(b.String())
}

func (m Model) renderPokedexList() string {
	var b strings.Builder

	// Title
	b.WriteString(titleStyle.Render("Your Pokedex"))
	b.WriteString("\n")
	caught := len(m.cfg.Inventory)
	b.WriteString(subtitleStyle.Render(fmt.Sprintf("%d Pokemon caught", caught)))
	b.WriteString("\n\n")

	if len(m.pokedexNames) == 0 {
		b.WriteString(subtitleStyle.Render("Your Pokedex is empty. Go catch some Pokemon!"))
		b.WriteString("\n")
	} else {
		for i, name := range m.pokedexNames {
			cursor := "  "
			style := normalItemStyle
			if i == m.pokedexCursor {
				cursor = "> "
				style = selectedStyle
			}

			// Get Pokemon type for styling
			pokemon := m.cfg.Inventory[name]
			typeStr := ""
			if len(pokemon.Types) > 0 {
				typeStr = " " + getTypeStyle(pokemon.Types[0].Type.Name).Render(pokemon.Types[0].Type.Name)
			}

			b.WriteString(cursor + style.Render(name) + typeStr + "\n")
		}
	}

	// Help
	b.WriteString("\n")
	b.WriteString(helpStyle.Render(pokedexListHelp))

	return b.String()
}

// renderStatBar creates a visual bar for a stat value (max ~255)
func renderStatBar(value int) string {
	maxWidth := 20
	filled := (value * maxWidth) / 255
	if filled > maxWidth {
		filled = maxWidth
	}

	bar := strings.Repeat("█", filled) + strings.Repeat("░", maxWidth-filled)
	return subtitleStyle.Render(bar)
}
