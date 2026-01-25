package description

import (
	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/shared"
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
	store           *store.Store
	formData        *shared.FormData
	keymap          keymap
	help            help.Model
	descriptionForm *huh.Form
}

func New(store *store.Store, formData *shared.FormData) Model {
	return Model{
		store:    store,
		formData: formData,
		keymap:   newKeymap(),
		help:     help.New(),
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
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var model tea.Model
	var cmd tea.Cmd

	switch msg.(type) {
	case tea.KeyMsg:
		switch {
		}
	}

	model, cmd = m.descriptionForm.Update(msg)
	m.descriptionForm = model.(*huh.Form)
	m.descriptionForm.GetFocusedField().Focus()

	if m.descriptionForm.State == huh.StateCompleted {
		return m, shared.Save(m.store, m.formData)
	}

	return m, cmd
}

func (m Model) View() string {
	s := m.descriptionForm.View()

	s += m.helpView()

	return s
}
