package main

import (
	"log"

	"timelog/internal/components"
	"timelog/internal/store"

	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	store           *store.Store
	table           table.Model
	currentView     string
	projectSelect   *huh.Select[string]
	descriptionText *huh.Text
	stopwatch       stopwatch.Model
}

func NewModel(store *store.Store) Model {
	options := make([]huh.Option[string], len(store.Projects))
	for k, p := range store.Projects {
		options = append(options, huh.NewOption(p.Name, k))
	}

	return Model{
		store:       store,
		table:       components.NewTable(store),
		currentView: "main",
		projectSelect: huh.NewSelect[string]().
			Key("project").
			Options(options...).
			Title("Project"),
		descriptionText: huh.NewText().
			Title("Description").
			Key("description"),
		stopwatch: stopwatch.New(),
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			switch m.currentView {
			case "main":
				if m.table.Focused() {
					m.table.Blur()
				} else {
					m.table.Focus()
				}
			default:
				m.currentView = "main"
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			// return m, tea.Batch(
			// 	tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			// )
			return m, tea.Quit
		case "a", "+":
			m.currentView = "add"
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("184")).
		Bold(true)

	s := headerStyle.Render("Timelog")

	s += "\n\n"

	switch m.currentView {
	case "main":
		s += m.table.View()
	case "add":
		s += "Add View"
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
