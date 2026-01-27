package project_add

import (
	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/shared"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/steps"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	store          *store.Store
	formData       *shared.FormData
	help           help.Model
	projectAddForm *huh.Form
}

func New(store *store.Store, formData *shared.FormData) Model {
	return Model{
		store:    store,
		formData: formData,
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

	model, cmd = m.projectAddForm.Update(msg)
	m.projectAddForm = model.(*huh.Form)

	if m.projectAddForm.State == huh.StateCompleted {
		project := m.projectAddForm.GetString("project")
		return m, tea.Batch(
			cmd,
			shared.SaveProject(m.store, project),
			steps.ChangeStep(steps.StepProject),
		)
	} else {
		return m, cmd
	}
}

func (m Model) View() string {
	return m.projectAddForm.View()
}
