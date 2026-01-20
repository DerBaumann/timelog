package tui

import (
	"fmt"

	"github.com/DerBaumann/timelog/internal/styles"
)

func (m Model) View() string {
	if m.err != nil {
		s := fmt.Sprintf("There was an Error:\n%s\n\nPress (q/ctrl+c) to quit\n", m.err.Error())
		return styles.ErrorStyle.Render(s)
	}

	s := styles.HeaderStyle.Render("Timelog")

	s += "\n\n"

	s += m.currentRoute.View()

	return "\n" + s + "\n\n"
}
