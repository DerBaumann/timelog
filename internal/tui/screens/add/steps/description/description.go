package description

import (
	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/shared"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	store           *store.Store
	formData        *shared.FormData
	descriptionForm *huh.Form
}

func New(store *store.Store, formData *shared.FormData) Model {
	return Model{
		store:    store,
		formData: formData,
		descriptionForm: huh.NewForm(
			huh.NewGroup(
				huh.NewText().
					Title("Description").
					Value(&formData.Description),
			),
		),
	}
}

func (m Model) Init() tea.Cmd {
	return m.descriptionForm.GetFocusedField().Focus()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var model tea.Model
	var cmd tea.Cmd

	model, cmd = m.descriptionForm.Update(msg)
	m.descriptionForm = model.(*huh.Form)

	if m.descriptionForm.State == huh.StateCompleted {
		return m, shared.Save(m.store, m.formData)
	} else {
		return m, cmd
	}
}

func (m Model) View() string {
	return m.descriptionForm.View()
}
