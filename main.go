package main

import (
	"log"

	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/add"
	"github.com/DerBaumann/timelog/internal/tui/cmds"
	"github.com/DerBaumann/timelog/internal/tui/export"
	"github.com/DerBaumann/timelog/internal/tui/home"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	store        *store.Store
	currentRoute tea.Model
}

func NewModel(store *store.Store) Model {
	return Model{
		store:        store,
		currentRoute: home.New(store),
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

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
	}

	// route updates
	m.currentRoute, cmd = m.currentRoute.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("184")).
		Bold(true)

	s := headerStyle.Render("Timelog")

	s += "\n\n"

	s += m.currentRoute.View()

	return "\n" + s + "\n\n"
}

func main() {
	store, err := store.ReadFile()
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(NewModel(store))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
