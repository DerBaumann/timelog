package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const permissions = 0o644

type Minutes int

func MinutesSinceMidnight(t time.Time) Minutes {
	return Minutes(t.Hour()*60 + t.Minute())
}

func (m Minutes) Format() string {
	return fmt.Sprintf("%02d:%02d", m/60, m%60)
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
