package tui

import (
	"github.com/DerBaumann/timelog/internal/tui/cmds"
	"github.com/DerBaumann/timelog/internal/tui/screens/add"
	"github.com/DerBaumann/timelog/internal/tui/screens/export"
	"github.com/DerBaumann/timelog/internal/tui/screens/home"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.showQuit {
		model, cmd := m.quitForm.Update(msg)
		m.quitForm = model.(*huh.Form)

		if m.quitForm.State == huh.StateCompleted {
			if quitting := m.quitForm.GetBool(quittingKey); quitting {
				return m, tea.Quit
			} else {
				m.showQuit = false
				m.quitForm = newQuitForm()
			}
		}

		return m, cmd
	}

	switch msg := msg.(type) {
	case cmds.ErrMsg:
		if msg.Err != nil {
			m.err = msg.Err
			return m, nil
		}

	// default keymaps
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.showQuit = true
			return m, nil
		case "ctrl+c":
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

	var cmd tea.Cmd

	// route updates
	m.currentRoute, cmd = m.currentRoute.Update(msg)
	return m, cmd
}
