package home

func (m Model) View() string {
	s := m.table.View() + "\n"

	s += m.helpView()

	return s
}
