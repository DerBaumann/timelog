package project

import (
	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/shared"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/steps"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type keymap struct {
	add key.Binding
}

func newKeymap() keymap {
	return keymap{
		add: key.NewBinding(
			key.WithKeys("a", "+"),
			key.WithHelp("a", "add project"),
		),
	}
}

func (m Model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.add,
	})
}

type Model struct {
	formData    *shared.FormData
	projectForm *huh.Form
	keymap      keymap
	help        help.Model
}

func New(s *store.Store, formData *shared.FormData) Model {
	var options []huh.Option[string]
	for _, p := range s.Projects {
		options = append(options, huh.NewOption(p.Name, p.ID))
	}

	return Model{
		formData: formData,
		projectForm: huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Project").
					Options(options...).
					Key("project"),
			),
		),
		keymap: newKeymap(),
		help:   help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var model tea.Model
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.add):
			return m, steps.ChangeStep(steps.StepProjectAdd)
		}
	}

	model, cmd = m.projectForm.Update(msg)
	m.projectForm = model.(*huh.Form)

	if m.projectForm.State == huh.StateCompleted {
		m.formData.Project = m.projectForm.GetString("project")

		return m, steps.ChangeStep(steps.StepStopwatch)
	}

	return m, cmd
}

func (m Model) View() string {
	s := m.projectForm.View()

	s += m.helpView()

	return s
}
