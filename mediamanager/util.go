package mediamanager

import (
	"errors"
	"os"
)

// fileExists returns true if file exists, false if not found
func fileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
