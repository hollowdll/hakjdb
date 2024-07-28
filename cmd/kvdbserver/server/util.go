package server

import (
	"os/exec"
	"runtime"
	"strings"
)

// getOsInfo returns information about the server's operating system.
func getOsInfo() (string, error) {
	osInfo := runtime.GOOS

	switch osInfo {
	case "linux":
		cmd := exec.Command("uname", "-r", "-m")
		output, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return "Linux " + strings.TrimSpace(string(output)), nil
	case "windows":
		cmd := exec.Command("cmd", "/c", "ver")
		output, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(output)), nil
	default:
		return osInfo, nil
	}
}
