package kvdb

import (
	"os"
	"path/filepath"
)

// Data directory name
const dataDirName string = "data"

// Gets path to the data directory.
// Data directory is a subdir in the executable's parent dir.
func getDataDirPath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}

	parentDirPath := filepath.Dir(execPath)
	path := filepath.Join(parentDirPath, dataDirName)

	return path, nil
}

// Gets path to a sub directory in the data directory.
func getDataDirSubDirPath(dirName string) (string, error) {
	dataDirPath, err := getDataDirPath()
	if err != nil {
		return "", err
	}

	path := filepath.Join(dataDirPath, dirName)

	return path, nil
}

// Gets path to a file in a data directory sub directory.
func getDataDirSubDirFilePath(dirName string, fileName string) (string, error) {
	dirPath, err := getDataDirSubDirPath(dirName)
	if err != nil {
		return "", nil
	}

	filePath := filepath.Join(dirPath, fileName)

	return filePath, nil
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

// Creates a sub directory to the data directory if it doesn't exist.
// Creates the data directory if it doesn't exist.
// Does nothing if the directory exists.
func createDataDirSubDirIfNotExist(dirName string) error {
	dataDirPath, err := getDataDirPath()
	if err != nil {
		return err
	}
	dirPath, err := getDataDirSubDirPath(dirName)
	if err != nil {
		return err
	}
	err = createDirIfNotExist(dataDirPath)
	if err != nil {
		return err
	}
	err = createDirIfNotExist(dirPath)
	if err != nil {
		return err
	}

	return nil
}
