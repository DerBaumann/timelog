package styles

import "github.com/charmbracelet/lipgloss"

var ErrorStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#ff0000"))

var HeaderStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("184")).
	Bold(true)
