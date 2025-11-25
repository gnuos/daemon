//go:build !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd && !plan9 && !solaris
// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!plan9,!solaris

// Package daemon
package daemon

func lockFile(_ uintptr) error {
	return errNotSupported
}

func unlockFile(_ uintptr) error {
	return errNotSupported
}
