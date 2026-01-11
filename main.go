package main

import (
	"log"

	"timelog/internal/store"
	"timelog/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	store        *store.Store
	home         tui.HomeUI
	add          tui.AddUI
	currentRoute tui.RouteMsg
}

func NewModel(store *store.Store) Model {
	return Model{
		store:        store,
		home:         tui.NewHomeUI(store),
		add:          tui.NewAddUI(store),
		currentRoute: tui.HomeRoute,
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
	case tui.RouteMsg:
		switch msg {
		case tui.HomeRoute:
			m.currentRoute = tui.HomeRoute
		case tui.AddRoute:
			m.currentRoute = tui.AddRoute
		}
	}

	// route updates
	var model tea.Model
	switch m.currentRoute {
	case tui.HomeRoute:
		model, cmd = m.home.Update(msg)
		if mod, ok := model.(tui.HomeUI); ok {
			m.home = mod
		}
	case tui.AddRoute:
		model, cmd = m.add.Update(msg)
		if mod, ok := model.(tui.AddUI); ok {
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
	case tui.HomeRoute:
		s += m.home.View()
	case tui.AddRoute:
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
