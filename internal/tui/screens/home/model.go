package home

import (
	"fmt"

	"github.com/DerBaumann/timelog/internal/store"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
)

type Model struct {
	table  table.Model
	store  *store.Store
	keymap keymap
	help   help.Model
}

func New(s *store.Store) Model {
	return Model{
		table:  newTable(s),
		store:  s,
		keymap: newKeymap(),
		help:   help.New(),
	}
}

func newTable(s *store.Store) table.Model {
	cols := []table.Column{
		{Title: "Day", Width: 15},
		{Title: "Project", Width: 15},
		{Title: "Description", Width: 15},
		{Title: "Duration", Width: 15},
	}

	rows := []table.Row{}

	for _, e := range s.Entries {
		var project *store.Project
		for _, p := range s.Projects {
			if p.ID == e.ProjectID {
				project = &p
			}
		}
		rows = append(rows, table.Row{
			e.Date,
			project.Name,
			e.Description,
			fmt.Sprintf("%s - %s", e.StartTime.Format(), e.EndTime.Format()),
		})
	}

	t := table.New(
		table.WithColumns(cols),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	t.SetStyles(newTableStyles())

	return t
}
