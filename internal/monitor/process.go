package monitor

import (
	"os/exec"
	"strconv"
	"strings"
)

// ProcessInfo holds CPU and memory usage for a process.
type ProcessInfo struct {
	PID    string
	CPU    float64
	Memory float64
}

// GetProcessInfo returns CPU and memory usage for a given PID.
func GetProcessInfo(pid string) ProcessInfo {
	info := ProcessInfo{PID: pid}
	if pid == "" {
		return info
	}

	cmd := exec.Command("ps", "-p", pid, "-o", "%cpu,%mem")
	out, err := cmd.Output()
	if err != nil {
		return info
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) < 2 {
		return info
	}

	fields := strings.Fields(lines[1])
	if len(fields) >= 2 {
		info.CPU, _ = strconv.ParseFloat(fields[0], 64)
		info.Memory, _ = strconv.ParseFloat(fields[1], 64)
	}

	return info
}

// GetChildProcessInfo returns aggregated CPU/memory for a PID and all children.
func GetChildProcessInfo(pid string) ProcessInfo {
	info := GetProcessInfo(pid)
	if pid == "" {
		return info
	}

	// Find child processes
	cmd := exec.Command("pgrep", "-P", pid)
	out, err := cmd.Output()
	if err != nil {
		return info
	}

	children := strings.Split(strings.TrimSpace(string(out)), "\n")
	for _, childPID := range children {
		childPID = strings.TrimSpace(childPID)
		if childPID == "" {
			continue
		}
		childInfo := GetProcessInfo(childPID)
		info.CPU += childInfo.CPU
		info.Memory += childInfo.Memory
	}

	return info
}
