package add

import (
	"github.com/DerBaumann/timelog/internal/tui/cmds"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/steps"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/steps/description"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/steps/project"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/steps/project_add"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/steps/stopwatch"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.showCancel {
		var model tea.Model

		model, cmd = m.cancelConfirm.Update(msg)
		m.cancelConfirm = model.(*huh.Form)

		if m.cancelConfirm.State == huh.StateCompleted {
			if *m.cancelling {
				return New(m.store), cmds.GoTo(cmds.Home)
			} else {
				m.cancelConfirm.State = huh.StateNormal
				m.showCancel = false
			}
		}

		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.back):
			if m.showCancel {
				return New(m.store), cmds.GoTo(cmds.Home)
			}
			m.showCancel = true
			return m, m.cancelConfirm.Init()
		}
	case steps.StepMsg:
		switch msg {
		case steps.StepProject:
			m.step = project.New(m.store, m.formData)
		case steps.StepProjectAdd:
			m.step = project_add.New(m.store, m.formData)
		case steps.StepStopwatch:
			m.step = stopwatch.New(m.formData)
		case steps.StepDescription:
			m.step = description.New(m.store, m.formData)
		}
	}

	m.step, cmd = m.step.Update(msg)

	return m, cmd
}
