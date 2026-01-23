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
	store      *store.Store
	step       Step
	showCancel bool

	data       *FormData // should be changed to value
	newProject *string
	cancelling *bool

	cancelConfirm   *huh.Form
	projectAddForm  *huh.Form
	projectForm     *huh.Form
	descriptionForm *huh.Form
	stopwatch       stopwatch.Model

	keymap keymap
	help   help.Model
}

func New(store *store.Store) Model {
	data := &FormData{}
	var newProject string
	var cancelling bool

	return Model{
		store:      store,
		step:       StepProject,
		showCancel: false,

		data:       data,
		newProject: &newProject,
		cancelling: &cancelling,

		cancelConfirm:   newCancelConfirm(&cancelling),
		projectAddForm:  newProjectAddForm(&newProject),
		projectForm:     newProjectForm(store.Projects, &data.Project),
		descriptionForm: newDescriptionForm(&data.Description),
		stopwatch:       stopwatch.New(),

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

func newProjectAddForm(value *string) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project Name").
				Prompt("> ").
				Value(value),
		),
	)
}

func newProjectForm(projects map[string]store.Project, value *string) *huh.Form {
	var options []huh.Option[string]
	for k, p := range projects {
		options = append(options, huh.NewOption(p.Name, k))
	}

	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Project").
				Options(options...).
				Value(value),
		),
	)
}

func newDescriptionForm(value *string) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("Description").
				Value(value),
		),
	)
}
