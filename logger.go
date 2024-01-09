package kvdb

import (
	"fmt"
	"log"
	"time"
)

const logsDirName string = "logs"
const allLogsFile string = "all_logs.log"

type Logger interface {
	Debug(v ...any)
	Debugf(format string, v ...any)

	Info(v ...any)
	Infof(format string, v ...any)

	Error(v ...any)
	Errorf(format string, v ...any)

	Warning(v ...any)
	Warningf(format string, v ...any)

	Fatal(v ...any)
	Fatalf(format string, v ...any)

	EnableDebug()
}

type DefaultLogger struct {
	Logger *log.Logger
	debug  bool
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{
		Logger: log.Default(),
		debug:  false,
	}
}

func (l *DefaultLogger) EnableDebug() {
	l.debug = true
}

type logEntry struct {
	createdAt time.Time
	logType   LogType
	content   string
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

// Disabled
/*
func AppendLogFile(logType LogType, message string) error {
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
*/

func (l *DefaultLogger) Debug(v ...any) {
	if l.debug {
		l.Logger.Printf("[Debug] %s", fmt.Sprint(v...))
	}
}

func (l *DefaultLogger) Debugf(format string, v ...any) {
	if l.debug {
		l.Logger.Printf("[Debug] %s", fmt.Sprintf(format, v...))
	}
}

func (l *DefaultLogger) Info(v ...any) {
	l.Logger.Printf("[Info] %s", fmt.Sprint(v...))
}

func (l *DefaultLogger) Infof(format string, v ...any) {
	l.Logger.Printf("[Info] %s", fmt.Sprintf(format, v...))
}

func (l *DefaultLogger) Error(v ...any) {
	l.Logger.Printf("[Error] %s", fmt.Sprint(v...))
}

func (l *DefaultLogger) Errorf(format string, v ...any) {
	l.Logger.Printf("[Error] %s", fmt.Sprintf(format, v...))
}

func (l *DefaultLogger) Warning(v ...any) {
	l.Logger.Printf("[Warning] %s", fmt.Sprint(v...))
}

func (l *DefaultLogger) Warningf(format string, v ...any) {
	l.Logger.Printf("[Warning] %s", fmt.Sprintf(format, v...))
}

func (l *DefaultLogger) Fatal(v ...any) {
	l.Logger.Fatalf("[Fatal] %s", fmt.Sprint(v...))
}

func (l *DefaultLogger) Fatalf(format string, v ...any) {
	l.Logger.Fatalf("[Fatal] %s", fmt.Sprintf(format, v...))
}
