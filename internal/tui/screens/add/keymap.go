package add

import "github.com/charmbracelet/bubbles/key"

type keymap struct {
	back key.Binding
}

func newKeymap() keymap {
	return keymap{
		back: key.NewBinding(
			key.WithKeys("esc", "delete"),
			key.WithHelp("esc", "cancel"),
		),
	}
}

func (m Model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.back,
	})
}
