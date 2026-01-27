package stopwatch

import (
	"time"

	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/shared"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/steps"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
)

type keymap struct {
	start,
	stop,
	save key.Binding
}

func newKeymap() keymap {
	return keymap{
		start: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "start clock"),
			key.WithDisabled(),
		),
		stop: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "stop clock"),
		),
		save: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "save time"),
		),
	}
}

func (m Model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.save,
		m.keymap.start,
		m.keymap.stop,
	})
}

type Model struct {
	formData  *shared.FormData
	keymap    keymap
	help      help.Model
	stopwatch stopwatch.Model
}

func New(formData *shared.FormData) Model {
	formData.StartTime = store.MinutesSinceMidnight(time.Now())

	return Model{
		formData:  formData,
		keymap:    newKeymap(),
		help:      help.New(),
		stopwatch: stopwatch.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return m.stopwatch.Start()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.start, m.keymap.stop):
			m.keymap.start.SetEnabled(m.stopwatch.Running())
			m.keymap.stop.SetEnabled(!m.stopwatch.Running())
			return m, m.stopwatch.Toggle()
		case key.Matches(msg, m.keymap.save):
			m.formData.EndTime = store.MinutesSinceMidnight(time.Now())
			m.formData.Duration = m.stopwatch.Elapsed()

			return m, tea.Batch(
				m.stopwatch.Stop(),
				steps.ChangeStep(steps.StepDescription),
			)
		}
	}

	m.stopwatch, cmd = m.stopwatch.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	s := m.stopwatch.View()

	s += m.helpView()

	return s
}
