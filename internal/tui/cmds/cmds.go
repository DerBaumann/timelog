package cmds

import tea "github.com/charmbracelet/bubbletea"

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

type ErrMsg struct {
	Err error
}

func ErrCmd(err error) tea.Cmd {
	return func() tea.Msg {
		return ErrMsg{
			Err: err,
		}
	}
}
