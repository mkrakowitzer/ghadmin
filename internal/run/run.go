package run

import (
	"os/exec"
)

// Runnable is typically an exec.Cmd or its stub in tests
type Runnable interface {
	Output() ([]byte, error)
	Run() error
}

// PrepareCmd extends exec.Cmd with extra error reporting features and provides a
// hook to stub command execution in tests
var PrepareCmd = func(cmd *exec.Cmd) Runnable {
	return &cmdWithStderr{cmd}
}

// SetPrepareCmd overrides PrepareCmd and returns a func to revert it back
func SetPrepareCmd(fn func(*exec.Cmd) Runnable) func() {
	origPrepare := PrepareCmd
	PrepareCmd = fn
	return func() {
		PrepareCmd = origPrepare
	}
}

// cmdWithStderr augments exec.Cmd by adding stderr to the error message
type cmdWithStderr struct {
	*exec.Cmd
}
