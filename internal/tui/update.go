package tui

import (
	"github.com/DerBaumann/timelog/internal/tui/cmds"
	"github.com/DerBaumann/timelog/internal/tui/screens/add"
	"github.com/DerBaumann/timelog/internal/tui/screens/export"
	"github.com/DerBaumann/timelog/internal/tui/screens/home"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case cmds.ErrMsg:
		if msg.Err != nil {
			m.err = msg.Err
			return m, nil
		}

	// default keymaps
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

	// routing
	case cmds.Route:
		switch msg {
		case cmds.Home:
			m.currentRoute = home.New(m.store)
		case cmds.Add:
			m.currentRoute = add.New(m.store)
		case cmds.Export:
			m.currentRoute = export.New(m.store)
		}

		return m, m.currentRoute.Init()
	}

	// route updates
	m.currentRoute, cmd = m.currentRoute.Update(msg)
	return m, cmd
}
