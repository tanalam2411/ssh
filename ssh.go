package ssh

import (
	gossh "golang.org/x/crypto/ssh"
	"net"
)

type Signal string

// POSIX signals as listed in RFC 4245 Section 6.10.
const (
	SIGABRT Signal = "ABRT"
	SIGALRM Signal = "ALRM"
	SIGFPE  Signal = "FPE"
	SIGHUP  Signal = "HUP"
	SIGILL  Signal = "ILL"
	SIGINT  Signal = "INT"
	SIGKILL Signal = "KILL"
	SIGPIPE Signal = "PIPE"
	SIGQUIT Signal = "QUIT"
	SIGSEGV Signal = "SEGV"
	SIGTERM Signal = "TERM"
	SIGUSR1 Signal = "USR1"
	SIGUSR2 Signal = "USR2"
)

// Handler is a callback fir handling established SSH sessions.
type Handler func(Session)

// PublicKeyHandler is a callback for performing public key authentication.
type PublicKeyHandler func(ctx Context, key PublicKey) bool

// PasswordHandler is a callback for performing password authentication.
type PasswordHandler func(ctx Context, password string) bool

// KeyboardInteractiveHandler is a callback for performing keyboard-interactive authentication.
type KeyboardInteractiveHandler func(ctx Context, challenger gossh.KeyboardInteractiveChallenge) bool

// PtyCallback is a hook for allowing PTY sessions.
type PtyCallback func(ctx Context, pty Pty) bool

// SessionRequestCallback is a callback for allowing or denying SSH sessions.
type SessionRequestCallback func(sess Session, requestType string) bool

// ConnCallback is a hook for new connections before handling.
// It allows wrapping for timeouts and limiting by returning
// the net.Conn that will be used as the underlying connection.
type ConnCallback func(ctx Context, conn net.Conn) net.Conn

// LocalPortForwardingCallback is a hook for allowing port forwarding
type LocalPortForwardingCallback func(ctx Context, destinationHost string, destinationPort uint32) bool

// ReversePortForwardingCallback is a hook for allowing reverse port forwarding.
type ReversePortForwardingCallback func(ctx Context, bindHost string, bindPort uint32) bool

// ServerConfigCallback is a hook for creating custom default server configs
type ServerConfigCallback func(ctx Context) *gossh.ServerConfig

// Window represents the size of a PTY window.
type Window struct {
	Width  int
	Height int
}

// http://man7.org/linux/man-pages/man7/pty.7.html
// Pty represents a PTY request and configuration.
type Pty struct {
	Term   string
	Window Window
}
