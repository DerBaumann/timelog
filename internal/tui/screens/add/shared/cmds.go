package shared

import (
	"strings"
	"time"

	"github.com/DerBaumann/timelog/internal/store"
	"github.com/DerBaumann/timelog/internal/tui/cmds"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

func Save(s *store.Store, formData *FormData) tea.Cmd {
	return func() tea.Msg {
		id, err := uuid.NewRandom()
		if err != nil {
			return cmds.ErrCmd(err)()
		}

		s.Entries = append(s.Entries, store.Entry{
			ID:          id,
			ProjectID:   formData.Project,
			Date:        time.Now().Format("2006-01-02"),
			Description: formData.Description,
			StartTime:   formData.StartTime,
			EndTime:     formData.EndTime,
			CreatedAt:   time.Now(),
		})

		if err := s.Write(); err != nil {
			return cmds.ErrCmd(err)()
		}

		return cmds.GoTo(cmds.Home)()
	}
}

func SaveProject(s *store.Store, newProject string) tea.Cmd {
	return func() tea.Msg {
		pid := strings.ReplaceAll(strings.ToLower(newProject), " ", "_")
		s.Projects[pid] = store.Project{
			Name: newProject,
		}

		if err := s.Write(); err != nil {
			return cmds.ErrCmd(err)()
		}

		return nil
	}
}
