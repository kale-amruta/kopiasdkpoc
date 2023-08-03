package osexec

import (
	"os/exec"
	"syscall"
)

// DisableInterruptSignal modifies child process attributes so that parent Ctrl-C is not propagated to a child.
func DisableInterruptSignal(c *exec.Cmd) {
	c.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
}
