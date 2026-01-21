package add

import "github.com/charmbracelet/bubbles/key"

type keymap struct {
	add,
	start,
	stop,
	save,
	back key.Binding
}

func newKeymap() keymap {
	return keymap{
		add: key.NewBinding(
			key.WithKeys("a", "+"),
			key.WithHelp("a", "add project"),
		),
		start: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "start clock"),
			key.WithDisabled(),
		),
		stop: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "stop clock"),
			key.WithDisabled(),
		),
		save: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "save time"),
			key.WithDisabled(),
		),
		back: key.NewBinding(
			key.WithKeys("esc", "delete"),
			key.WithHelp("esc", "cancel"),
		),
	}
}

func (m Model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.add,
		m.keymap.save,
		m.keymap.start,
		m.keymap.stop,
		m.keymap.back,
	})
}
