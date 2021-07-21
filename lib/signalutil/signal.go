package signalutil

import (
	"os"
	"os/signal"
	"syscall"
)

// NewSighupChan Creates a new channel for
// receiving SIGHUP system calls
func NewSighupChan() chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP)
	return ch
}
