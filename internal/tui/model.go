package tui

import (
	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/screens/home"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

const quittingKey = "quitting"

type Model struct {
	store        *store.Store
	err          error
	currentRoute tea.Model
	quitForm     *huh.Form
	showQuit     bool
}

func New(store *store.Store) Model {
	return Model{
		store:        store,
		err:          nil,
		currentRoute: home.New(store),
		quitForm:     newQuitForm(),
		showQuit:     false,
	}
}

func newQuitForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Are you sure?").
				Affirmative("Yes").
				Negative("No").
				Key(quittingKey),
		),
	)
}
