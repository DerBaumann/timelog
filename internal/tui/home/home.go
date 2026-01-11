package home

import (
	"fmt"

	"timelog/internal/store"
	"timelog/internal/tui/routes"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keymap struct {
	add key.Binding
}

type Model struct {
	table  table.Model
	store  *store.Store
	keymap keymap
	help   help.Model
}

func New(store *store.Store) Model {
	cols := []table.Column{
		{Title: "Day", Width: 15},
		{Title: "Project", Width: 15},
		{Title: "Description", Width: 15},
		{Title: "Duration", Width: 15},
	}

	rows := []table.Row{}

	for _, e := range store.Entries {
		rows = append(rows, table.Row{
			e.Date,
			store.Projects[e.ProjectID].Name,
			e.Description,
			fmt.Sprintf("%s - %s", e.StartTime, e.EndTime),
		})
	}

	t := table.New(
		table.WithColumns(cols),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	k := keymap{
		add: key.NewBinding(
			key.WithKeys("a", "+"),
			key.WithHelp("a", "add new entry"),
		),
	}

	return Model{
		table:  t,
		store:  store,
		keymap: k,
		help:   help.New(),
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.add):
			return m, routes.GoTo(routes.Add)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.add,
		m.table.KeyMap.LineDown,
		m.table.KeyMap.LineUp,
	})
}

func (m Model) View() string {
	s := m.table.View() + "\n"

	s += m.helpView()

	return s
}
