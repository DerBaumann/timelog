package main

import (
	"log"

	"timelog/internal/store"
	"timelog/internal/tui/add"
	"timelog/internal/tui/home"
	"timelog/internal/tui/routes"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	store        *store.Store
	home         home.Model
	add          add.Model
	currentRoute routes.Route
}

func NewModel(store *store.Store) Model {
	return Model{
		store:        store,
		home:         home.New(store),
		add:          add.New(store),
		currentRoute: routes.Home,
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
	case routes.Route:
		// switch msg {
		// case routes.Home:
		// 	m.home = home.New(m.store)
		// case routes.Add:
		// 	m.add = add.New(m.store)
		// }
		m.currentRoute = msg
	}

	// route updates
	var model tea.Model
	switch m.currentRoute {
	case routes.Home:
		model, cmd = m.home.Update(msg)
		if mod, ok := model.(home.Model); ok {
			m.home = mod
		}
	case routes.Add:
		model, cmd = m.add.Update(msg)
		if mod, ok := model.(add.Model); ok {
			m.add = mod
		}
	}

	return m, cmd
}

func (m Model) View() string {
	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("184")).
		Bold(true)

	s := headerStyle.Render("Timelog")

	s += "\n\n"

	switch m.currentRoute {
	case routes.Home:
		s += m.home.View()
	case routes.Add:
		s += m.add.View()
	default:
		return "View not found!"
	}

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
