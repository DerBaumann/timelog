package home

import (
	"slices"

	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/cmds"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.deleting {
		var model tea.Model

		model, cmd = m.deleteForm.Update(msg)
		m.deleteForm = model.(*huh.Form)

		if m.deleteForm.State == huh.StateCompleted {
			// Get selected row
			row := m.table.SelectedRow()

			// Get ID from row
			id := row[0]

			// Remove project from store.Projects ( filter (p -> p.id != id) )
			m.store.Projects = slices.DeleteFunc(m.store.Projects, func(p store.Project) bool {
				return p.ID == id
			})

			if err := m.store.Write(); err != nil {
				return m, cmds.ErrCmd(err)
			}

			m.deleting = false
		}

		return m, cmd
	} else {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, m.keymap.add):
				return m, cmds.GoTo(cmds.Add)
			case key.Matches(msg, m.keymap.export):
				return m, cmds.GoTo(cmds.Export)
			}
		}
		m.table, cmd = m.table.Update(msg)
		return m, cmd
	}
}
