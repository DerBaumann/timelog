package project_add

import (
	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/shared"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/steps"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type keymap struct{}

func newKeymap() keymap {
	return keymap{}
}

func (m Model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{})
}

type Model struct {
	store          *store.Store
	formData       *shared.FormData
	keymap         keymap
	help           help.Model
	projectAddForm *huh.Form
}

func New(store *store.Store, formData *shared.FormData) Model {
	return Model{
		store:    store,
		formData: formData,
		keymap:   newKeymap(),
		help:     help.New(),
		projectAddForm: huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Project Name").
					Prompt("> ").
					Key("project"),
			),
		),
	}
}

func (m Model) Init() tea.Cmd {
	return m.projectAddForm.GetFocusedField().Focus()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var model tea.Model
	var cmd tea.Cmd

	switch msg.(type) {
	case tea.KeyMsg:
		switch {
		}
	}

	model, cmd = m.projectAddForm.Update(msg)
	m.projectAddForm = model.(*huh.Form)

	if m.projectAddForm.State == huh.StateCompleted {
		project := m.projectAddForm.GetString("project")
		return m, tea.Batch(
			cmd,
			shared.SaveProject(m.store, project),
			steps.ChangeStep(steps.StepProject),
		)
	}

	return m, cmd
}

func (m Model) View() string {
	s := m.projectAddForm.View()

	s += m.helpView()

	return s
}
