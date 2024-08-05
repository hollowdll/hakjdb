package logging

import (
	"strings"

	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/viper"
)

const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
	LogLevelFatal
)

func getLogLevelStr() string {
	return viper.GetString(common.ConfigKeyLogLevel)
}

func getLogLevel(ready bool) int {
	if !ready {
		return LogLevelInfo
	}
	logLevelStr := getLogLevelStr()
	switch strings.ToLower(logLevelStr) {
	case "debug":
		return LogLevelDebug
	case "info":
		return LogLevelInfo
	case "warning":
		return LogLevelWarning
	case "error":
		return LogLevelError
	case "fatal":
		return LogLevelFatal
	default:
		// default log level is info
		return LogLevelInfo
	}
}
