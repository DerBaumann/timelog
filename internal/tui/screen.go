package tui

import (
	"github.com/DerBaumann/timelog/internal/tui/screens/home"
	tea "github.com/charmbracelet/bubbletea"
)

type Route int

const (
	RouteHome Route = iota
	RouteProject
	RouteAddProject
	RouteStopwatch
	RouteDescription
	RouteExport
)

func (m Model) GoTo(route Route) tea.Cmd {
	return func() tea.Msg {
		switch m.currentRoute.(type) {
		case home.Model:
			switch route {
			case RouteProject, RouteExport:
				return route
			}
		}
		return nil
	}
}
