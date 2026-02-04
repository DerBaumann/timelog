package home

import (
	"fmt"

	"github.com/DerBaumann/timelog/internal/store"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/huh"
)

const deleteFormKey = "delete-form"

type Model struct {
	table      table.Model
	store      *store.Store
	keymap     keymap
	help       help.Model
	deleteForm *huh.Form
	deleting   bool
}

func New(s *store.Store) Model {
	return Model{
		table:      newTable(s),
		store:      s,
		keymap:     newKeymap(),
		help:       help.New(),
		deleteForm: newDeleteForm(),
		deleting:   false,
	}
}

func newDeleteForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Are you sure, you want to delete this?").
				Affirmative("Yes!").
				Negative("No").
				Key(deleteFormKey),
		),
	)
}

func newTable(s *store.Store) table.Model {
	cols := []table.Column{
		{Title: "ID", Width: 0},
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
			e.ID.String(),
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
