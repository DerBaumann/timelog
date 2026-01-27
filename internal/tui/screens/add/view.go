package add

func (m Model) View() string {
	s := m.step.View()

	if m.showCancel {
		s = m.cancelConfirm.View()
	}

	s += "\n\n"

	s += m.formData.Project + "\n"
	s += m.formData.StartTime.Format() + "\n"
	s += m.formData.Duration.String() + "\n"
	s += m.formData.EndTime.Format() + "\n"
	s += m.formData.Description + "\n"

	s += m.helpView()

	return s
}
