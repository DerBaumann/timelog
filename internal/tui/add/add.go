package add

import (
	"time"

	"timelog/internal/store"
	"timelog/internal/tui/routes"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type FormStep int

const (
	Project FormStep = iota
	Stopwatch
	Description
)

type FormData struct {
	Project     string
	StartTime   time.Time
	EndTime     time.Time
	Duration    time.Duration
	Description string
}

type keymap struct {
	start key.Binding
	stop  key.Binding
	save  key.Binding
	back  key.Binding
}

type Model struct {
	store         *store.Store
	data          *FormData
	preTimerForm  *huh.Form
	stopwatch     stopwatch.Model
	postTimerForm *huh.Form
	step          FormStep
	keymap        keymap
	help          help.Model
}

func New(store *store.Store) Model {
	var options []huh.Option[string]
	for k, p := range store.Projects {
		options = append(options, huh.NewOption(p.Name, k))
	}

	data := &FormData{}

	k := keymap{
		start: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "start clock"),
			key.WithDisabled(),
		),
		stop: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "stop clock"),
			key.WithDisabled(),
		),
		save: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "save time"),
			key.WithDisabled(),
		),
		back: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "cancel"),
		),
	}

	return Model{
		store: store,
		data:  data,
		preTimerForm: huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Project").
					Options(options...).
					Value(&data.Project),
			),
		),
		stopwatch: stopwatch.New(),
		postTimerForm: huh.NewForm(
			huh.NewGroup(
				huh.NewText().
					Title("Description").
					Value(&data.Description),
			),
		),
		step:   Project,
		keymap: k,
		help:   help.New(),
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.back):
			m.step = Project
			m.keymap.save.SetEnabled(false)
			m.keymap.start.SetEnabled(false)
			m.keymap.stop.SetEnabled(false)
			m.data = &FormData{}

			return m, routes.GoTo(routes.Home)
		}
	}

	var model tea.Model
	var cmd tea.Cmd

	switch m.step {
	case Project:
		model, cmd = m.preTimerForm.Update(msg)
		m.preTimerForm = model.(*huh.Form)

		if m.preTimerForm.State == huh.StateCompleted {
			m.step = Stopwatch
			m.keymap.save.SetEnabled(true)
			m.keymap.start.SetEnabled(true)
			m.keymap.stop.SetEnabled(true)
		}
	case Stopwatch:
		m.stopwatch, cmd = m.stopwatch.Update(msg)
	case Description:
		model, cmd = m.postTimerForm.Update(msg)
		m.postTimerForm = model.(*huh.Form)
	}

	return m, cmd
}

func (m Model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.save,
		m.keymap.start,
		m.keymap.stop,
		m.keymap.back,
	})
}

func (m Model) View() string {
	var s string

	switch m.step {
	case Project:
		s = m.preTimerForm.View()
	case Stopwatch:
		s = m.stopwatch.View()
	case Description:
		s = m.postTimerForm.View()
	}

	s += m.helpView()

	return s
}
