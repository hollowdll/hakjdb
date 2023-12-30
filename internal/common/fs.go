package common

import (
	"os"
	"path/filepath"
)

// GetExecParentDirPath gets path to the executable's parent directory.
func GetExecParentDirPath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	path := filepath.Dir(execPath)

	return path, nil
}

// FileExists returns a bool indicating if a file exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsExist(err)
}
