package home

import (
	"slices"

	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/cmds"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/google/uuid"
)

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.deleting {
		var model tea.Model

		model, cmd = m.deleteForm.Update(msg)
		m.deleteForm = model.(*huh.Form)

		if m.deleteForm.State == huh.StateCompleted {
			if m.deleteForm.GetBool(deleteFormKey) {
				id := uuid.MustParse(m.table.SelectedRow()[0])

				m.store.Entries = slices.DeleteFunc(m.store.Entries, func(e store.Entry) bool {
					return e.ID == id
				})

				if err := m.store.Write(); err != nil {
					return m, cmds.ErrCmd(err)
				}

				m.table = newTable(m.store)
			}

			m.deleting = false
			m.deleteForm = newDeleteForm()
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
			case key.Matches(msg, m.keymap.delete):
				m.deleting = true
				return m, nil
			}
		}
		m.table, cmd = m.table.Update(msg)
		return m, cmd
	}
}
