package store

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type Store struct {
	Version  int                `json:"version"`
	Entries  []Entry            `json:"entries"`
	Projects map[string]Project `json:"projects"`
}

// Both times are the amount of minutes since midnight
type Entry struct {
	ID          uuid.UUID `json:"id"`
	ProjectKey  string    `json:"project_key"`
	Date        string    `json:"date"`
	Description string    `json:"description"`
	StartTime   int       `json:"start_time"`
	EndTime     int       `json:"end_time"`
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
		return "", nil
	}

	path := filepath.Join(configDir, "timelog", "store.json")

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return "", err
	}

	return path, nil
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

	store := New()

	file, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {

			data, err := json.Marshal(store)
			if err != nil {
				return nil, err
			}

			file, err = os.Create(path)
			if err != nil {
				return nil, err
			}

			if _, err := file.Write(data); err != nil {
				file.Close()
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	defer file.Close()

	json.NewDecoder(file).Decode(store)

	return store, nil
}
