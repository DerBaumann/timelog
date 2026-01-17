package export

import (
	"slices"

	"github.com/DerBaumann/timelog/internal/store"
	tea "github.com/charmbracelet/bubbletea"
)

type DoneMsg struct{}

type Model struct {
	loading bool
	store   *store.Store
}

func New(store *store.Store) Model {
	return Model{
		loading: true,
		store:   store,
	}
}

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		var dates []string
		for _, entry := range m.store.Entries {
			if !slices.Contains(dates, entry.Date) {
				dates = append(dates, entry.Date)
			}
		}

		return DoneMsg{}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case DoneMsg:
		return m, tea.Quit
	}

	return m, nil
}

func (m Model) View() string {
	if m.loading {
		return "Loading"
	} else {
		return "Done!"
	}
}
