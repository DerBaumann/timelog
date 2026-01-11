package routes

import tea "github.com/charmbracelet/bubbletea"

type Route int

const (
	Home Route = iota
	Add
)

func GoTo(to Route) tea.Cmd {
	return func() tea.Msg {
		return to
	}
}
