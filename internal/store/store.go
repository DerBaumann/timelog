package store

import (
	"encoding/json"
	"errors"
	"os"
)

type Store struct {
	Version  int       `json:"version"`
	Entries  []Entry   `json:"entries"`
	Projects []Project `json:"projects"`
}

func New() *Store {
	return &Store{
		Version:  1,
		Entries:  []Entry{},
		Projects: []Project{},
	}
}

func Read() (*Store, error) {
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
