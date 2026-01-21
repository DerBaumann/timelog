package add

import (
	"time"

	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/cmds"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Step int

const (
	StepProject Step = iota
	StepProjectAdd
	StepStopwatch
	StepDescription
)

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.add):
			if m.step == StepProject {
				m.step = StepProjectAdd
				m.projectAddForm.GetFocusedField().Focus()
				return m, nil
			}
		case key.Matches(msg, m.keymap.back):
			return New(m.store), cmds.GoTo(cmds.Home)
		case key.Matches(msg, m.keymap.start, m.keymap.stop):
			m.keymap.start.SetEnabled(m.stopwatch.Running())
			m.keymap.stop.SetEnabled(!m.stopwatch.Running())
			return m, m.stopwatch.Toggle()
		case key.Matches(msg, m.keymap.save):
			m.data.EndTime = store.MinutesSinceMidnight(time.Now())
			m.data.Duration = m.stopwatch.Elapsed()
			m.step = StepDescription

			m.keymap.save.SetEnabled(false)
			m.keymap.start.SetEnabled(false)
			m.keymap.stop.SetEnabled(false)

			return m, m.stopwatch.Stop()
		}
	}

	var model tea.Model
	var cmd tea.Cmd

	switch m.step {
	case StepProject:
		model, cmd = m.projectForm.Update(msg)
		m.projectForm = model.(*huh.Form)

		if m.projectForm.State == huh.StateCompleted {
			m.step = StepStopwatch
			m.keymap.add.SetEnabled(false)
			m.keymap.save.SetEnabled(true)
			m.keymap.start.SetEnabled(true)
			m.data.StartTime = store.MinutesSinceMidnight(time.Now())
			return m, tea.Batch(
				m.stopwatch.Start(),
				cmd,
			)
		}

		return m, cmd
	case StepProjectAdd:
		model, cmd = m.projectAddForm.Update(msg)
		m.projectAddForm = model.(*huh.Form)

		if m.projectAddForm.State == huh.StateCompleted {
			return New(m.store), tea.Batch(cmd, m.saveProject)
		}

		return m, cmd
	case StepStopwatch:
		m.stopwatch, cmd = m.stopwatch.Update(msg)
		return m, cmd
	case StepDescription:
		model, cmd = m.descriptionForm.Update(msg)
		m.descriptionForm = model.(*huh.Form)
		m.descriptionForm.GetFocusedField().Focus()

		if m.descriptionForm.State == huh.StateCompleted {
			return New(m.store), m.save
		}

		return m, cmd
	}

	return m, cmd
}
