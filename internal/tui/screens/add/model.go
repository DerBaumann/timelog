package add

import (
	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/shared"
	"github.com/DerBaumann/timelog/internal/tui/screens/add/steps/project"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	store      *store.Store
	step       tea.Model
	showCancel bool

	formData   *shared.FormData // should be changed to value
	cancelling *bool

	cancelConfirm *huh.Form

	keymap keymap
	help   help.Model
}

func New(store *store.Store) Model {
	data := &shared.FormData{}
	var cancelling bool

	return Model{
		store:      store,
		step:       project.New(store, data),
		showCancel: false,

		formData:   data,
		cancelling: &cancelling,

		cancelConfirm: newCancelConfirm(&cancelling),

		keymap: newKeymap(),
		help:   help.New(),
	}
}

func newCancelConfirm(value *bool) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Cancel entry?").
				Affirmative("Yes").
				Negative("No").
				Value(value),
		),
	)
}
