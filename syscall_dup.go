//go:build (!linux || !arm64) && !windows
// +build !linux !arm64
// +build !windows

// Package daemon
package daemon

import (
	"syscall"
)

func syscallDup(oldfd int, newfd int) (err error) {
	return syscall.Dup2(oldfd, newfd)
}
