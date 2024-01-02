package kvdb

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const logsDirName string = "logs"
const allLogsFile string = "all_logs.log"

// Logger manages read and write operations to log files.
type Logger struct{}

type logEntry struct {
	createdAt time.Time
	logType   LogType
	content   string
}

// NewLogger returns a new logger.
func NewLogger() *Logger {
	return &Logger{}
}

func newLogEntry(logType LogType, content string) *logEntry {
	return &logEntry{
		createdAt: time.Now(),
		logType:   logType,
		content:   content,
	}
}

func (l *logEntry) String() string {
	return fmt.Sprintf("[%s] [%s] %s\n", l.createdAt.Format(time.RFC3339), l.logType, l.content)
}

// LogType represents the type of log.
type LogType int

const (
	// LogTypeInfo is log type for informative log messages.
	LogTypeInfo LogType = iota
)

func (l LogType) String() string {
	switch l {
	case LogTypeInfo:
		return "Info"
	default:
		return "Invalid LogType"
	}
}

// LogMessage writes message to log file.
func (l Logger) LogMessage(logType LogType, message string) error {
	dirPath, err := createDataDirSubDirIfNotExist(logsDirName)
	if err != nil {
		return err
	}
	filePath := filepath.Join(dirPath, allLogsFile)
	logEntry := newLogEntry(logType, message)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(logEntry.String())
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
