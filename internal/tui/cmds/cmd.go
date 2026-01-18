package cmds

import tea "github.com/charmbracelet/bubbletea"

type RefreshDataMsg struct{}

func RefreshData() tea.Msg { return RefreshDataMsg{} }

type Route int

const (
	Home Route = iota
	Add
	Export
)

func GoTo(to Route) tea.Cmd {
	return func() tea.Msg {
		return to
	}
}
