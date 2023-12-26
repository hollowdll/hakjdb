package kvdb

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const logsDirName string = "logs"
const allLogsFile string = "all_logs.log"

// Logger manages read and write operations to log files.
type Logger struct{}

type logEntry struct {
	createdAt time.Time
	logType   logType
	content   string
}

func (l *logEntry) String() string {
	return fmt.Sprintf("[%s] [%s] %s", l.createdAt.UTC(), l.logType, l.content)
}

type logType int

const (
	logTypeInfo logType = iota
)

func (l logType) String() string {
	switch l {
	case logTypeInfo:
		return "Info"
	default:
		return "Invalid LogType"
	}
}

// Logs info to log file.
func (l Logger) LogInfo(info string) error {
	err := createDataDirSubDirIfNotExist(logsDirName)
	if err != nil {
		return nil
	}
	filePath, err := getDataDirSubDirFilePath(logsDirName, allLogsFile)
	if err != nil {
		return nil
	}
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(info)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
