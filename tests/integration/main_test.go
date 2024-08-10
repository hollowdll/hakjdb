package integration

import (
	"os"
	"testing"
	"time"
)

const ctxTimeout = time.Second * 5

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
