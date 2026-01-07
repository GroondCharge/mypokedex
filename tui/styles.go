package tui

import "github.com/charmbracelet/lipgloss"

// Colors
var (
	primaryColor   = lipgloss.Color("#FF6B6B")
	secondaryColor = lipgloss.Color("#4ECDC4")
	accentColor    = lipgloss.Color("#FFE66D")
	subtleColor    = lipgloss.Color("#666666")
	errorColor     = lipgloss.Color("#FF4757")
	successColor   = lipgloss.Color("#2ED573")
)

// Pokemon type colors
var typeColors = map[string]lipgloss.Color{
	"normal":   lipgloss.Color("#A8A878"),
	"fire":     lipgloss.Color("#F08030"),
	"water":    lipgloss.Color("#6890F0"),
	"electric": lipgloss.Color("#F8D030"),
	"grass":    lipgloss.Color("#78C850"),
	"ice":      lipgloss.Color("#98D8D8"),
	"fighting": lipgloss.Color("#C03028"),
	"poison":   lipgloss.Color("#A040A0"),
	"ground":   lipgloss.Color("#E0C068"),
	"flying":   lipgloss.Color("#A890F0"),
	"psychic":  lipgloss.Color("#F85888"),
	"bug":      lipgloss.Color("#A8B820"),
	"rock":     lipgloss.Color("#B8A038"),
	"ghost":    lipgloss.Color("#705898"),
	"dragon":   lipgloss.Color("#7038F8"),
	"dark":     lipgloss.Color("#705848"),
	"steel":    lipgloss.Color("#B8B8D0"),
	"fairy":    lipgloss.Color("#EE99AC"),
}

// UI Styles
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primaryColor).
			MarginBottom(1)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(subtleColor).
			Italic(true)

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(secondaryColor).
			Bold(true).
			Padding(0, 1)

	normalItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(0, 1)

	caughtBadgeStyle = lipgloss.NewStyle().
				Foreground(successColor).
				Bold(true)

	messageStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true).
			MarginTop(1)

	errorStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true)

	helpStyle = lipgloss.NewStyle().
			Foreground(subtleColor).
			MarginTop(1)

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(secondaryColor).
			Padding(1, 2)

	statLabelStyle = lipgloss.NewStyle().
			Width(15).
			Foreground(subtleColor)

	statValueStyle = lipgloss.NewStyle().
			Bold(true)
)

// getTypeStyle returns a styled string for a Pokemon type
func getTypeStyle(typeName string) lipgloss.Style {
	color, ok := typeColors[typeName]
	if !ok {
		color = lipgloss.Color("#FFFFFF")
	}
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")).
		Background(color).
		Bold(true).
		Padding(0, 1)
}
