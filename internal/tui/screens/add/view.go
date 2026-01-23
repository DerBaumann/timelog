package add

func (m Model) View() string {
	var s string

	if m.showCancel {
		s = m.cancelConfirm.View()
	} else {
		switch m.step {
		case StepProject:
			s = m.projectForm.View()
		case StepProjectAdd:
			s = m.projectAddForm.View()
		case StepStopwatch:
			s = m.stopwatch.View()
		case StepDescription:
			s = m.descriptionForm.View()
		}
	}

	s += m.helpView()

	return s
}
