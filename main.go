package main

import (
	"fmt"
	"log"

	"timelog/internal/store"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	store *store.Store
}

func NewModel(store *store.Store) Model {
	return Model{store}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("184")).
		Background(lipgloss.Color("26")).
		Bold(true).
		Padding(1, 4)

	s := headerStyle.Render("Timelog")

	s += fmt.Sprintf("\n\n%+v", *m.store)

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
