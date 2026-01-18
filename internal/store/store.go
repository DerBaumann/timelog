package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

const permissions = 0o644

type Minutes int

func MinutesSinceMidnight(t time.Time) Minutes {
	return Minutes(t.Hour()*60 + t.Minute())
}

func (m Minutes) Format() string {
	return fmt.Sprintf("%02d:%02d", m/60, m%60)
}

type Store struct {
	Version  int                `json:"version"`
	Entries  []Entry            `json:"entries"`
	Projects map[string]Project `json:"projects"`
}

// Both times store the timestamps as minutes since midnight
type Entry struct {
	ID          uuid.UUID `json:"id"`
	ProjectID   string    `json:"project_id"`
	Date        string    `json:"date"`
	Description string    `json:"description"`
	StartTime   Minutes   `json:"start_time"`
	EndTime     Minutes   `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
}

type Project struct {
	Name string `json:"name"`
}

func getPath() (string, error) {
	if path, ok := os.LookupEnv("TIMELOG_STOREPATH"); ok {
		return path, nil
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(configDir, "timelog", "store.json")

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return "", err
	}

	return path, nil
}

func writeEmptyStore(path string) ([]byte, error) {
	data, err := json.Marshal(New())
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(path, data, permissions); err != nil {
		return nil, err
	}

	return data, nil
}

func New() *Store {
	return &Store{
		Version:  1,
		Entries:  []Entry{},
		Projects: map[string]Project{},
	}
}

func ReadFile() (*Store, error) {
	path, err := getPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			data, err = writeEmptyStore(path)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	store := &Store{}
	if err := json.Unmarshal(data, &store); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *Store) Write() error {
	path, err := getPath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(*s)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, data, permissions); err != nil {
		return err
	}

	return nil
}
