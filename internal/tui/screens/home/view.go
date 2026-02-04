package home

func (m Model) View() string {
	var s string

	if m.deleting {
		s = m.deleteForm.View() + "\n"
	} else {
		s = m.table.View() + "\n"
	}

	s += m.helpView()

	return s
}
