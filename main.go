package main

import (
	"log"

	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	store, err := store.ReadFile()
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(tui.New(store))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
