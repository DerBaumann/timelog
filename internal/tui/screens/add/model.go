package add

import (
	"time"

	"github.com/DerBaumann/timelog/internal/store"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/charmbracelet/huh"
)

type FormData struct {
	Project     string
	StartTime   store.Minutes
	EndTime     store.Minutes
	Duration    time.Duration
	Description string
}

type Model struct {
	store *store.Store
	step  Step

	data *FormData // should be changed to value

	newProject      *string
	projectAddForm  *huh.Form
	projectForm     *huh.Form
	descriptionForm *huh.Form
	stopwatch       stopwatch.Model

	keymap keymap
	help   help.Model
}

func New(store *store.Store) Model {
	var options []huh.Option[string]
	for k, p := range store.Projects {
		options = append(options, huh.NewOption(p.Name, k))
	}

	data := &FormData{}
	var newProject string

	k := newKeymap()

	return Model{
		store:      store,
		data:       data,
		newProject: &newProject,
		projectAddForm: huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Project Name").
					Prompt("> ").
					Value(&newProject),
			),
		),
		projectForm: huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Project").
					Options(options...).
					Value(&data.Project),
			),
		),
		stopwatch: stopwatch.New(),
		descriptionForm: huh.NewForm(
			huh.NewGroup(
				huh.NewText().
					Title("Description").
					Value(&data.Description),
			),
		),
		step:   StepProject,
		keymap: k,
		help:   help.New(),
	}
}
