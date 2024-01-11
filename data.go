package kvdb

import (
	"os"
	"path/filepath"

	"github.com/hollowdll/kvdb/internal/common"
)

// Data directory name
const dataDirName string = "data"

// GetDataDirPath gets path to the data directory.
// Data directory is a subdir in the executable's parent dir.
// Creates the directory if it does not exist.
func GetDataDirPath() (string, error) {
	parentDirPath, err := common.GetExecParentDirPath()
	if err != nil {
		return "", err
	}
	path := filepath.Join(parentDirPath, dataDirName)
	if err = createDirIfNotExist(path); err != nil {
		return "", err
	}

	return path, nil
}

// Creates a directory if it doesn't exist.
// Does nothing if the directory exists.
func createDirIfNotExist(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
