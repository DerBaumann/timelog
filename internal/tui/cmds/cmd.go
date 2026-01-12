package cmds

import tea "github.com/charmbracelet/bubbletea"

type RefreshDataMsg struct{}

func RefreshData() tea.Msg { return RefreshDataMsg{} }
