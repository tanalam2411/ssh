package ssh

import "errors"


// ErrServerClosed is returned by the Server's Serve, ListenAndServe,
// and ListenAndServeTLS methods after a call to Shutdown or Close.
var ErrServerClosed = errors.New("ssh: Server closed")


type Server struct {
	Addr string // TCP address to listen on, ":22" if empty
	Handler Handler
}
