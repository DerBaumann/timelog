package steps

import (
	tea "github.com/charmbracelet/bubbletea"
)

type StepMsg int

const (
	StepProject StepMsg = iota
	StepProjectAdd
	StepStopwatch
	StepDescription
	StepCancel
)

func ChangeStep(step StepMsg) tea.Cmd {
	return func() tea.Msg {
		return step
	}
}
