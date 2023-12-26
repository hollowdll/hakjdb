package kvdb

import (
	"os"
	"path/filepath"
)

// Data directory name
const dataDirName string = "data"

// Gets path to the data directory.
// Data directory is a subdir in the executable's parent dir.
func getDataDirectoryPath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}

	parentDirPath := filepath.Dir(execPath)
	dataDirPath := filepath.Join(parentDirPath, dataDirName)

	return dataDirPath, nil
}

// Creates data directory if it doesn't exist.
func createDataDirectoryIfNotExist() error {
	dataDirPath, err := getDataDirectoryPath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(dataDirPath); os.IsNotExist(err) {
		// create directory
		err := os.Mkdir(dataDirPath, os.ModePerm)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
