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
	EnableLogFile(filePath string) error
	// CloseLogFile closes the log file for I/O operations.
	CloseLogFile() error
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
	FileLogger     *log.Logger
	logFile        *os.File
	debug          bool
	logFileEnabled bool
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{
		Logger:         log.New(os.Stderr, "", 0),
		FileLogger:     log.New(io.Discard, "", 0),
		logFile:        nil,
		debug:          false,
		logFileEnabled: false,
	}
}

func (l *DefaultLogger) EnableDebug() {
	l.debug = true
}

func (l *DefaultLogger) EnableLogFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	l.FileLogger = log.New(file, "", 0)
	l.logFile = file
	l.logFileEnabled = true

	return nil
}

func (l *DefaultLogger) CloseLogFile() error {
	if l.logFile != nil {
		return l.logFile.Close()
	}
	return nil
}

func (l *DefaultLogger) ClearFlags() {
	l.Logger.SetFlags(0)
}

func (l *DefaultLogger) Disable() {
	l.Logger.SetOutput(io.Discard)
}

func (l *DefaultLogger) writeToFile(logMsg string) {
	if l.logFileEnabled {
		l.FileLogger.Print(logMsg)
	}
}

func (l *DefaultLogger) Debug(v ...any) {
	if l.debug {
		logMsg := fmt.Sprintf("%s [Debug] %s", timestampPrefix(), fmt.Sprint(v...))
		l.Logger.Print(logMsg)
		l.writeToFile(logMsg)
	}
}

func (l *DefaultLogger) Debugf(format string, v ...any) {
	if l.debug {
		logMsg := fmt.Sprintf("%s [Debug] %s", timestampPrefix(), fmt.Sprintf(format, v...))
		l.Logger.Print(logMsg)
		l.writeToFile(logMsg)
	}
}

func (l *DefaultLogger) Info(v ...any) {
	logMsg := fmt.Sprintf("%s [Info] %s", timestampPrefix(), fmt.Sprint(v...))
	l.Logger.Print(logMsg)
	l.writeToFile(logMsg)
}

func (l *DefaultLogger) Infof(format string, v ...any) {
	logMsg := fmt.Sprintf("%s [Info] %s", timestampPrefix(), fmt.Sprintf(format, v...))
	l.Logger.Print(logMsg)
	l.writeToFile(logMsg)
}

func (l *DefaultLogger) Error(v ...any) {
	logMsg := fmt.Sprintf("%s [Error] %s", timestampPrefix(), fmt.Sprint(v...))
	l.Logger.Print(logMsg)
	l.writeToFile(logMsg)
}

func (l *DefaultLogger) Errorf(format string, v ...any) {
	logMsg := fmt.Sprintf("%s [Error] %s", timestampPrefix(), fmt.Sprintf(format, v...))
	l.Logger.Print(logMsg)
	l.writeToFile(logMsg)
}

func (l *DefaultLogger) Warning(v ...any) {
	logMsg := fmt.Sprintf("%s [Warning] %s", timestampPrefix(), fmt.Sprint(v...))
	l.Logger.Print(logMsg)
	l.writeToFile(logMsg)
}

func (l *DefaultLogger) Warningf(format string, v ...any) {
	logMsg := fmt.Sprintf("%s [Warning] %s", timestampPrefix(), fmt.Sprintf(format, v...))
	l.Logger.Print(logMsg)
	l.writeToFile(logMsg)
}

func (l *DefaultLogger) Fatal(v ...any) {
	logMsg := fmt.Sprintf("%s [Fatal] %s", timestampPrefix(), fmt.Sprint(v...))
	l.Logger.Fatal(logMsg)
	l.writeToFile(logMsg)
}

func (l *DefaultLogger) Fatalf(format string, v ...any) {
	logMsg := fmt.Sprintf("%s [Fatal] %s", timestampPrefix(), fmt.Sprintf(format, v...))
	l.Logger.Fatal(logMsg)
	l.writeToFile(logMsg)
}

func timestampPrefix() string {
	return time.Now().Format("2006-01-02T15:04:05.999Z07:00")
}
