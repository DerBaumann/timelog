package utils

import (
	"os"
	"path/filepath"
)

func GetStoreFile() (string, error) {
	if path, ok := os.LookupEnv("TIMELOG_STOREPATH"); ok {
		return path, nil
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", nil
	}

	return filepath.Join(configDir, "timelog", "store.json"), nil
}
