package export

import (
	"os"
	"slices"
	"strings"
	"text/template"

	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/cmds"
	tea "github.com/charmbracelet/bubbletea"
)

type DoneMsg struct{}

type DocEntry struct {
	Project,
	Description,
	Start,
	End string
}

type DocDay struct {
	Date    string
	Entries []DocEntry
}

type DocumentData struct {
	Days []DocDay
}

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

const docTmpl = `# Arbeitsjournal

{{ range .Days }}
## {{ .Date }}


| Projekt | Beschreibung | Zeit |
| ------- | ------------ | ---- |
{{ range .Entries }}| {{ .Project }} | {{ .Description }} | {{ .Start }} - {{ .End }} |
{{ end }}
{{ end }}`

func exportCmd(s *store.Store) tea.Cmd {
	return func() tea.Msg {
		var dates []string
		for _, entry := range s.Entries {
			if !slices.Contains(dates, entry.Date) {
				dates = append(dates, entry.Date)
			}
		}

		var days []DocDay
		for _, date := range dates {
			var entries []DocEntry
			for _, e := range s.Entries {
				if e.Date == date {
					entries = append(entries, DocEntry{
						Project:     s.Projects[e.ProjectID].Name,
						Description: strings.ReplaceAll(e.Description, "\n", "<br />"),
						Start:       e.StartTime.Format(),
						End:         e.EndTime.Format(),
					})
				}
			}

			days = append(days, DocDay{
				Date:    date,
				Entries: entries,
			})
		}
		docData := DocumentData{
			Days: days,
		}

		tmpl, err := template.New("mddoc").Parse(docTmpl)
		if err != nil {
			return cmds.ErrMsg{Err: err}
		}

		f, err := os.Create("./export.md")
		if err != nil {
			return cmds.ErrMsg{Err: err}
		}
		defer f.Close()

		if err := tmpl.Execute(f, docData); err != nil {
			return cmds.ErrMsg{Err: err}
		}

		return DoneMsg{}
	}
}

func (m Model) Init() tea.Cmd {
	return exportCmd(m.store)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case DoneMsg:
		m.loading = false
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
