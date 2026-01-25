package add

func (m Model) View() string {
	s := m.step.View()

	if m.showCancel {
		s = m.cancelConfirm.View()
	}

	s += m.helpView()

	return s
}
