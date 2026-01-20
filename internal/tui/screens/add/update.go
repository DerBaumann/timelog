package add

import (
	"strings"
	"time"

	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/cmds"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/google/uuid"
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
			project := store.Project{
				Name: *m.newProject,
			}
			pid := strings.ReplaceAll(strings.ToLower(*m.newProject), " ", "_")
			m.store.Projects[pid] = project

			if err := m.store.Write(); err != nil {
				return m, cmds.ErrCmd(err)
			}

			return New(m.store), cmd
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
			// save data
			id, err := uuid.NewRandom()
			if err != nil {
				return m, cmds.ErrCmd(err)
			}

			m.store.Entries = append(m.store.Entries, store.Entry{
				ID:          id,
				ProjectID:   m.data.Project,
				Date:        time.Now().Format("2006-01-02"),
				Description: m.data.Description,
				StartTime:   m.data.StartTime,
				EndTime:     m.data.EndTime,
				CreatedAt:   time.Now(),
			})

			if err := m.store.Write(); err != nil {
				return m, cmds.ErrCmd(err)
			}

			return New(m.store), cmds.GoTo(cmds.Home)
		}

		return m, cmd
	}

	return m, cmd
}
