package tui

import (
	"fmt"

	"timelog/internal/store"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HomeUI struct {
	table table.Model
	store *store.Store
}

func NewHomeUI(store *store.Store) HomeUI {
	cols := []table.Column{
		{Title: "Day", Width: 15},
		{Title: "Project", Width: 15},
		{Title: "Description", Width: 15},
		{Title: "Duration", Width: 15},
	}

	rows := []table.Row{}

	for _, e := range store.Entries {
		rows = append(rows, table.Row{
			e.Date,
			store.Projects[e.ProjectID].Name,
			e.Description,
			fmt.Sprintf("%s - %s", e.StartTime, e.EndTime),
		})
	}

	t := table.New(
		table.WithColumns(cols),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return HomeUI{
		table: t,
		store: store,
	}
}

func (m HomeUI) Init() tea.Cmd { return nil }

func (m HomeUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			// return m, tea.Batch(
			// 	tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			// )
			return m, tea.Quit
		case "a", "+":
			return m, Route(AddRoute)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m HomeUI) View() string { return m.table.View() }
