package common

import (
	"bufio"
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

// GetDirPath gets path to a directory and returns it.
// parent is the parent directory and dirName the name of the directory.
// Creates the directory if it does not exist.
func GetDirPath(parent string, dirName string) (string, error) {
	path := filepath.Join(parent, dirName)
	if err := createDirIfNotExist(path); err != nil {
		return "", err
	}

	return path, nil
}

// createDirIfNotExist creates a directory if it doesn't exist.
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

// CreateFileIfNotExist creates a file if it doesn't exist.
func CreateFileIfNotExist(filePath string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
	} else if err != nil {
		return err
	}

	return nil
}

// ReadFileLines reads a file and returns its lines
func ReadFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
