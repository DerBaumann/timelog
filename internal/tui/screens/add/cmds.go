package add

import (
	"strings"
	"time"

	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/cmds"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func (m Model) save() tea.Msg {
	// save data
	id, err := uuid.NewRandom()
	if err != nil {
		return cmds.ErrCmd(err)()
	}

	m.store.Entries = append(m.store.Entries, store.Entry{
		ID:          id,
		ProjectID:   m.data.Project,
		Date:        time.Now().Format("2006-01-02"),
		Description: m.data.Description,
		StartTime:   m.data.StartTime,
		EndTime:     m.data.EndTime,
		CreatedAt:   time.Now(),
	})

	if err := m.store.Write(); err != nil {
		return cmds.ErrCmd(err)()
	}

	return cmds.GoTo(cmds.Home)()
}

func (m Model) saveProject() tea.Msg {
	pid := strings.ReplaceAll(strings.ToLower(*m.newProject), " ", "_")
	m.store.Projects[pid] = store.Project{
		Name: *m.newProject,
	}

	if err := m.store.Write(); err != nil {
		return cmds.ErrCmd(err)()
	}

	return nil
}
