package kvdb

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const (
	LogLevelDebug   LogLevel = 0
	LogLevelInfo    LogLevel = 1
	LogLevelWarning LogLevel = 2
	LogLevelError   LogLevel = 3
	LogLevelFatal   LogLevel = 4

	LogLevelDebugStr   string = "debug"
	LogLevelInfoStr    string = "info"
	LogLevelWarningStr string = "warning"
	LogLevelErrorStr   string = "error"
	LogLevelFatalStr   string = "fatal"

	DefaultLogLevel    LogLevel = LogLevelInfo
	DefaultLogLevelStr string   = LogLevelInfoStr
)

type LogLevel uint8

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

	// SetLogLevel sets the log level.
	SetLogLevel(level LogLevel)

	// LogLevel returns the log level and its string equivalent.
	LogLevel() (LogLevel, string)

	// EnableLogFile enables log file.
	EnableLogFile(filePath string) error

	// CloseLogFile closes the log file if it is open.
	CloseLogFile() error

	// Disable disables all log outputs.
	Disable()
}

// DefaultLogger is a default implementation of the Logger interface.
// Log output defaults to standard error stream.
// Debug logs are disabled by default. Call EnableDebug to enable them.
type DefaultLogger struct {
	logger         *log.Logger
	fileLogger     *log.Logger
	logFile        *os.File
	logLevel       LogLevel
	logFileEnabled bool
}

func NewDefaultLogger() *DefaultLogger {
	lg := &DefaultLogger{
		logger:         log.New(os.Stderr, "", 0),
		fileLogger:     log.New(io.Discard, "", 0),
		logFile:        nil,
		logLevel:       LogLevelInfo,
		logFileEnabled: false,
	}
	lg.clearFlags()
	return lg
}

func (l *DefaultLogger) SetLogLevel(level LogLevel) {
	l.logLevel = level
}

func (l *DefaultLogger) LogLevel() (LogLevel, string) {
	switch l.logLevel {
	case LogLevelDebug:
		return LogLevelDebug, LogLevelDebugStr
	case LogLevelInfo:
		return LogLevelInfo, LogLevelInfoStr
	case LogLevelWarning:
		return LogLevelWarning, LogLevelWarningStr
	case LogLevelError:
		return LogLevelError, LogLevelErrorStr
	case LogLevelFatal:
		return LogLevelFatal, LogLevelFatalStr
	default:
		return DefaultLogLevel, DefaultLogLevelStr
	}
}

func (l *DefaultLogger) EnableLogFile(filepath string) error {
	file, err := openLogFile(filepath)
	if err != nil {
		return err
	}

	l.fileLogger.SetOutput(file)
	l.logFile = file
	l.logFileEnabled = true

	return nil
}

func openLogFile(filepath string) (*os.File, error) {
	return os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
}

func (l *DefaultLogger) CloseLogFile() error {
	if l.logFile != nil {
		return l.logFile.Close()
	}
	return nil
}

func (l *DefaultLogger) clearFlags() {
	l.logger.SetFlags(0)
}

func (l *DefaultLogger) Disable() {
	l.logger.SetOutput(io.Discard)
	l.fileLogger.SetOutput(io.Discard)
}

func (l *DefaultLogger) writeToFile(logMsg string) {
	if l.logFileEnabled {
		l.fileLogger.Print(logMsg)
	}
}

func (l *DefaultLogger) Debug(v ...any) {
	if l.logLevel <= LogLevelDebug {
		logMsg := fmt.Sprintf("%s [Debug] %s", timestampPrefix(), fmt.Sprint(v...))
		l.logger.Print(logMsg)
		l.writeToFile(logMsg)
	}
}

func (l *DefaultLogger) Debugf(format string, v ...any) {
	if l.logLevel <= LogLevelDebug {
		logMsg := fmt.Sprintf("%s [Debug] %s", timestampPrefix(), fmt.Sprintf(format, v...))
		l.logger.Print(logMsg)
		l.writeToFile(logMsg)
	}
}

func (l *DefaultLogger) Info(v ...any) {
	if l.logLevel <= LogLevelInfo {
		logMsg := fmt.Sprintf("%s [Info] %s", timestampPrefix(), fmt.Sprint(v...))
		l.logger.Print(logMsg)
		l.writeToFile(logMsg)
	}
}

func (l *DefaultLogger) Infof(format string, v ...any) {
	if l.logLevel <= LogLevelInfo {
		logMsg := fmt.Sprintf("%s [Info] %s", timestampPrefix(), fmt.Sprintf(format, v...))
		l.logger.Print(logMsg)
		l.writeToFile(logMsg)
	}
}

func (l *DefaultLogger) Warning(v ...any) {
	if l.logLevel <= LogLevelWarning {
		logMsg := fmt.Sprintf("%s [Warning] %s", timestampPrefix(), fmt.Sprint(v...))
		l.logger.Print(logMsg)
		l.writeToFile(logMsg)
	}
}

func (l *DefaultLogger) Warningf(format string, v ...any) {
	if l.logLevel <= LogLevelWarning {
		logMsg := fmt.Sprintf("%s [Warning] %s", timestampPrefix(), fmt.Sprintf(format, v...))
		l.logger.Print(logMsg)
		l.writeToFile(logMsg)
	}
}

func (l *DefaultLogger) Error(v ...any) {
	if l.logLevel <= LogLevelError {
		logMsg := fmt.Sprintf("%s [Error] %s", timestampPrefix(), fmt.Sprint(v...))
		l.logger.Print(logMsg)
		l.writeToFile(logMsg)
	}
}

func (l *DefaultLogger) Errorf(format string, v ...any) {
	if l.logLevel <= LogLevelError {
		logMsg := fmt.Sprintf("%s [Error] %s", timestampPrefix(), fmt.Sprintf(format, v...))
		l.logger.Print(logMsg)
		l.writeToFile(logMsg)
	}
}

func (l *DefaultLogger) Fatal(v ...any) {
	if l.logLevel == LogLevelFatal {
		logMsg := fmt.Sprintf("%s [Fatal] %s", timestampPrefix(), fmt.Sprint(v...))
		l.logger.Fatal(logMsg)
		l.writeToFile(logMsg)
	}
}

func (l *DefaultLogger) Fatalf(format string, v ...any) {
	if l.logLevel == LogLevelFatal {
		logMsg := fmt.Sprintf("%s [Fatal] %s", timestampPrefix(), fmt.Sprintf(format, v...))
		l.logger.Fatal(logMsg)
		l.writeToFile(logMsg)
	}
}

func timestampPrefix() string {
	return time.Now().Format("2006-01-02T15:04:05.999Z07:00")
}

// GetLogLevelFromStr returns the log level that matches its string equivalent.
// If invalid log level string is passed, this function returns the default log level.
// The returned string is the log level's string equivalent in lowercase.
// The returned bool is true if the passed string is valid log level.
func GetLogLevelFromStr(levelStr string) (LogLevel, string, bool) {
	switch strings.ToLower(levelStr) {
	case LogLevelDebugStr:
		return LogLevelDebug, LogLevelDebugStr, true
	case LogLevelInfoStr:
		return LogLevelInfo, LogLevelInfoStr, true
	case LogLevelWarningStr:
		return LogLevelWarning, LogLevelWarningStr, true
	case LogLevelErrorStr:
		return LogLevelError, LogLevelErrorStr, true
	case LogLevelFatalStr:
		return LogLevelFatal, LogLevelFatalStr, true
	default:
		// return default log level if invalid
		return DefaultLogLevel, DefaultLogLevelStr, false
	}
}
