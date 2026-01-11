package tui

import (
	"time"

	"timelog/internal/store"

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

type AddUI struct {
	store         *store.Store
	data          *FormData
	preTimerForm  *huh.Form
	stopwatch     stopwatch.Model
	postTimerForm *huh.Form
	step          FormStep
}

func NewAddUI(store *store.Store) AddUI {
	var options []huh.Option[string]
	for k, p := range store.Projects {
		options = append(options, huh.NewOption(p.Name, k))
	}

	data := &FormData{}

	return AddUI{
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
		step: Project,
	}
}

func (m AddUI) Init() tea.Cmd { return nil }

func (m AddUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, Route(HomeRoute)
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
		}
	case Stopwatch:
		m.stopwatch, cmd = m.stopwatch.Update(msg)
	case Description:
		model, cmd = m.postTimerForm.Update(msg)
		m.postTimerForm = model.(*huh.Form)
	}

	return m, cmd
}

func (m AddUI) View() string {
	var s string

	switch m.step {
	case Project:
		s = m.preTimerForm.View()
	case Stopwatch:
		s = m.stopwatch.View()
	case Description:
		s = m.postTimerForm.View()

	}
	return s
}
