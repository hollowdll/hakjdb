package kvdb

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

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

	// EnableDebug enables debug logs.
	EnableDebug()
	// EnableLogFile enables log file.
	EnableLogFile(filePath string)
	// ClearFlags clears all default logger output flags.
	ClearFlags()
	// Disable disables all log outputs.
	Disable()
}

// DefaultLogger is a default implementation of the Logger interface.
// Log output defaults to standard error stream.
// Debug logs are disabled by default. Call EnableDebug to enable them.
type DefaultLogger struct {
	Logger         *log.Logger
	debug          bool
	logFileEnabled bool
	logFilePath    string
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{
		Logger:         log.New(os.Stderr, "", 0),
		debug:          false,
		logFileEnabled: false,
		logFilePath:    "",
	}
}

func (l *DefaultLogger) EnableDebug() {
	l.debug = true
}

func (l *DefaultLogger) EnableLogFile(filePath string) {
	l.logFileEnabled = true
	l.logFilePath = filePath
}

func (l *DefaultLogger) ClearFlags() {
	l.Logger.SetFlags(0)
}

func (l *DefaultLogger) Disable() {
	l.Logger.SetOutput(io.Discard)
}

func (l *DefaultLogger) Debug(v ...any) {
	if l.debug {
		l.Logger.Printf("%s [Debug] %s", timestampPrefix(), fmt.Sprint(v...))
	}
}

func (l *DefaultLogger) Debugf(format string, v ...any) {
	if l.debug {
		l.Logger.Printf("%s [Debug] %s", timestampPrefix(), fmt.Sprintf(format, v...))
	}
}

func (l *DefaultLogger) Info(v ...any) {
	l.Logger.Printf("%s [Info] %s", timestampPrefix(), fmt.Sprint(v...))
}

func (l *DefaultLogger) Infof(format string, v ...any) {
	l.Logger.Printf("%s [Info] %s", timestampPrefix(), fmt.Sprintf(format, v...))
}

func (l *DefaultLogger) Error(v ...any) {
	l.Logger.Printf("%s [Error] %s", timestampPrefix(), fmt.Sprint(v...))
}

func (l *DefaultLogger) Errorf(format string, v ...any) {
	l.Logger.Printf("%s [Error] %s", timestampPrefix(), fmt.Sprintf(format, v...))
}

func (l *DefaultLogger) Warning(v ...any) {
	l.Logger.Printf("%s [Warning] %s", timestampPrefix(), fmt.Sprint(v...))
}

func (l *DefaultLogger) Warningf(format string, v ...any) {
	l.Logger.Printf("%s [Warning] %s", timestampPrefix(), fmt.Sprintf(format, v...))
}

func (l *DefaultLogger) Fatal(v ...any) {
	l.Logger.Fatalf("%s [Fatal] %s", timestampPrefix(), fmt.Sprint(v...))
}

func (l *DefaultLogger) Fatalf(format string, v ...any) {
	l.Logger.Fatalf("%s [Fatal] %s", timestampPrefix(), fmt.Sprintf(format, v...))
}

func timestampPrefix() string {
	return time.Now().Format("2006-01-02T15:04:05.999Z07:00")
}
