package testutil

import (
	"github.com/hollowdll/kvdb"
)

func DisabledLogger() kvdb.Logger {
	lg := kvdb.NewDefaultLogger()
	lg.Disable()
	return lg
}
