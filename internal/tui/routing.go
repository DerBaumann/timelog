package tui

import tea "github.com/charmbracelet/bubbletea"

type RouteMsg int

const (
	HomeRoute RouteMsg = iota
	AddRoute
)

func Route(to RouteMsg) tea.Cmd {
	return func() tea.Msg {
		return to
	}
}
