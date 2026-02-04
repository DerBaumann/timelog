package home

import "github.com/charmbracelet/bubbles/key"

type keymap struct {
	add,
	export,
	delete key.Binding
}

func newKeymap() keymap {
	return keymap{
		add: key.NewBinding(
			key.WithKeys("a", "+"),
			key.WithHelp("a", "add new entry"),
		),
		export: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "export data"),
		),
		delete: key.NewBinding(
			key.WithKeys("D"),
			key.WithHelp("D", "delete selected item"),
		),
	}
}

func (m Model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.add,
		m.keymap.export,
		m.keymap.delete,
		m.table.KeyMap.LineDown,
		m.table.KeyMap.LineUp,
	})
}
