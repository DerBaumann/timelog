package components

import (
	"fmt"

	"timelog/internal/store"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func NewTable(store *store.Store) table.Model {
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

	return t
}
