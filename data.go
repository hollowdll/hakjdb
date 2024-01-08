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

// Gets path to a sub directory in the data directory.
func getDataDirSubDirPath(dirName string) (string, error) {
	dataDirPath, err := GetDataDirPath()
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
		return "", err
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
//
// Returns the path to the sub directory.
func createDataDirSubDirIfNotExist(dirName string) (string, error) {
	dataDirPath, err := GetDataDirPath()
	if err != nil {
		return "", err
	}
	dirPath := filepath.Join(dataDirPath, dirName)
	err = createDirIfNotExist(dataDirPath)
	if err != nil {
		return "", err
	}
	err = createDirIfNotExist(dirPath)
	if err != nil {
		return "", err
	}

	return dirPath, nil
}
