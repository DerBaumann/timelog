package export

import (
	"os"
	"slices"
	"strings"
	"text/template"

	"github.com/DerBaumann/timelog/internal/store"
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

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		var dates []string
		for _, entry := range m.store.Entries {
			if !slices.Contains(dates, entry.Date) {
				dates = append(dates, entry.Date)
			}
		}

		var days []DocDay
		for _, date := range dates {
			var entries []DocEntry
			for _, e := range m.store.Entries {
				if e.Date == date {
					entries = append(entries, DocEntry{
						Project:     m.store.Projects[e.ProjectID].Name,
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

		tmplStr := `# Arbeitsjournal

		{{ range .Dates }}
		## {{ .Date }}

		| Projekt | Beschreibung | Zeit |
		| ------- | ------------ | ---- |
		{{ range .Entries }}
		| {{ .Project }} | {{ .Description }} | {{ .Start }} - {{ .End }} |
		{{ end }}
		{{ end }}`

		tmpl, err := template.New("mddoc").Parse(tmplStr)
		if err != nil {
			panic(err)
		}

		f, err := os.Create("./export.md")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if err := tmpl.Execute(f, docData); err != nil {
			panic(err)
		}

		return DoneMsg{}
	}
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
