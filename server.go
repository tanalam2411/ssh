package ssh

import (
	"errors"
	gossh "golang.org/x/crypto/ssh"
	"net"
	"sync"
	"time"
)

// ErrServerClosed is returned by the Server's Serve, ListenAndServe,
// and ListenAndServeTLS methods after a call to Shutdown or Close.
var ErrServerClosed = errors.New("ssh: Server closed")

type RequestHandler func(ctx Context, srv *Server, req *gossh.Request) (ok bool, payload []byte)

type ChannelHandler func(srv *Server, conn *gossh.ServerConn, newChan gossh.NewChannel, ctx Context)

type Server struct {
	Addr        string   // TCP address to listen on, ":22" if empty
	Handler     Handler  // handler to invoke, ssh.DefaultHandler if nil
	HostSigners []Signer // private keys for the host key, must have at least one
	Version     string   // server version to be sent before the initial handshake

	KeyboardInteractiveHandler    KeyboardInteractiveHandler    // keyboard-interactive authentication handler
	PasswordHandler               PasswordHandler               // password authentication handler
	PublicKeyHandler              PublicKeyHandler              // public key authentication handler
	PtyCallback                   PtyCallback                   // callback for allowing PTY sessions, allows all if nil
	ConnCallback                  ConnCallback                  // optional callback for wrapping net.Conn before handling
	LocalPortForwardingCallback   LocalPortForwardingCallback   // callback for allowing local port forwarding, denies all if nil
	ReversePortForwardingCallback ReversePortForwardingCallback // callback for allowing reverse port forwarding, denies all if nil
	ServerConfigCallback          ServerConfigCallback          // callback for configuring detailed SSH options.
	SessionRequestCallback        SessionRequestCallback        // callback for allowing or denying SSH sessions

	IdleTimeout time.Duration // connection timeout when no activity, none if empty
	MaxTimeout  time.Duration // absolute connection timeout, none if empty

	// ChannelHandlers allow overriding the built-in session handlers or provide
	// extensions to the protocol, such as tcpip forwarding.
	// By default only the "session" handler is enabled.
	ChannelHandlers map[string]ChannelHandler

	// RequestHandlers allow overriding the several-level request handlers or
	// provide extensions to the protocol, such as tcpip forwarding.
	// By default no handlers are enabled.
	RequestHandlers map[string]RequestHandler

	listenerWg sync.WaitGroup
	mu         sync.RWMutex
	listeners  map[net.Listener]struct{}
	conns      map[*gossh.ServerConn]struct{}
	connWg     sync.WaitGroup
	doneChan   chan struct{}
}
