package ssh

import (
	"context"
	"net"
	"sync"
)

// contextKey is a value for use with context.WithValue.
// It's used as a pointer so it fits in an interface{} without allocation.
type contextKey struct {
	name string
}

var (
	// ContextKeyUser is a context key for use with Contexts in this package.
	// The associated value will be of type string.
	ContextKeyUser = &contextKey{"user"}

	// ContextKeySessionID is a context key for use with Contexts in this package.
	// The associated value will be of type string.
	ContextKeySessionID = &contextKey{"session-id"}

	// ContextKeyPermissions is context key for use with Contexts in this package.
	// The associated value will be of type *Permissions.
	ContextKeyPermissions = &contextKey{"permissions"}

	// ContextKeyClientVersion is a context key for use with Contexts in this package.
	// The associated value will be of type string.
	ContextKeyClientVersion = &contextKey{"client-version"}

	// ContextKeyServerVersion is a context key for use with Contexts in this package.
	ContextKetServerVersion = &contextKey{"server-version"}

	// ContextKeyLocalAddr is a context key for use with Contexts in this package.
	// The associated value will be of type "net.addr"
	ContextKeyLocalAddr = &contextKey{"local-addr"}

	// ContextKeyRemoteAddr is a context key for use with Contexts in this package.
	// The associated value will be of type `net.addr`
	ContextKeyRemoteAddr = &contextKey{"remote-addr"}

	// ContextKeyServer is a context key for use with Contexts in this package.
	// The associated value will be of type `*Server`.
	ContextKeyServer = &contextKey{"ssh-server"}

	// ContextKeyConn is a context key for use with Contexts in this package.
	// The associated value will be of type `gossh.ServerConn`.
	ContextKeyConn = &contextKey{"ssh-conn"}

	// ContextKeyPublicKey is a context key for use with Contexts in this package.
	// The associated value will be of type `PublicKey`.
	ContextKeyPublicKey = &contextKey{"public-key"}
)

// Context is a package specific context interface. It exposes connection
// metadata and allows new values to be easily written to it.
// It's used in authentication handlers and callbacks, and its underlying
// context.Context is exposed on Session in the session Handler.
// A connection-scoped lock is also embedded in the context to make it easier
// to limit operations per-connection.
type Context interface {
	context.Context
	sync.Locker

	// User returns the username used when establishing the SSH connection.
	User() string

	// SessionID returns the session hash.
	SessionID() string

	// ClientVersion returns the version reported by the client.
	ClientVersion() string

	// ServerVersion returns the version reported by the server.
	ServerVersion() string

	// RemoteAddr returns the remote address for this connection.
	RemoteAddr() net.Addr

	// Permissions returns the Permissions object used for this connection
	Permissions() *Permissions

	// SetValue allows you to easily write new values into the underlying context.
	SetValue(key, value interface{})
}


type sshContext struct {
	context.Context
	*sync.Mutex
}


func (ctx *sshContext) SetValue(key, value interface{}) {
	ctx.Context = context.WithValue(ctx.Context, key, value)
}


func (ctx *sshContext) User() string {
	return ctx.Value(ContextKeyUser).(string)
}


func (ctx *sshContext) SessionID() string {
	return ctx.Value(ContextKeySessionID).(string)
}

func (ctx *sshContext) ClientVersion() string {
	return ctx.Value(ContextKeyClientVersion).(string)
}

func (ctx *sshContext) ServerVersion() string {
	return ctx.Value(ContextKetServerVersion).(string)
}

func (ctx *sshContext) RemoteAddr() net.Addr {
	return ctx.Value(ContextKeyRemoteAddr).(net.Addr)
}


func (ctx *sshContext) LocalAddr() net.Addr {
	return ctx.Value(ContextKeyLocalAddr).(net.Addr)
}

func (ctx *sshContext) Permissions() *Permissions {
	return ctx.Value(ContextKeyPermissions).(*Permissions)
}


// ToDo - newContext(srv *Server) (*sshContext, context.CancelFunc)

// ToDo - applyConnMetadata(ctx Context, conn gossh.ConnMetadata)



