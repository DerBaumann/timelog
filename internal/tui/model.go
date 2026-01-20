package tui

import (
	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/screens/home"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	store        *store.Store
	err          error
	currentRoute tea.Model
}

func New(store *store.Store) Model {
	return Model{
		store:        store,
		err:          nil,
		currentRoute: home.New(store),
	}
}
