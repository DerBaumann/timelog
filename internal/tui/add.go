package tui

import (
	"timelog/internal/store"

	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type AddUI struct {
	store           *store.Store
	projectSelect   *huh.Select[string]
	descriptionText *huh.Text
	stopwatch       stopwatch.Model
}

func NewAddUI(store *store.Store) AddUI {
	options := make([]huh.Option[string], len(store.Projects))
	for k, p := range store.Projects {
		options = append(options, huh.NewOption(p.Name, k))
	}

	return AddUI{
		store: store,
		projectSelect: huh.NewSelect[string]().
			Key("project").
			Options(options...).
			Title("Project"),
		descriptionText: huh.NewText().
			Title("Description").
			Key("description"),
		stopwatch: stopwatch.New(),
	}
}

func (m AddUI) Init() tea.Cmd { return nil }

func (m AddUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, Route(HomeRoute)
		}
	}

	return m, cmd
}

func (m AddUI) View() string {
	s := "Add View"
	return s
}
