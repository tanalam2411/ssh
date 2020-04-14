package ssh

import (
	gossh "golang.org/x/crypto/ssh"
	"testing"
)

func newTestSessionWithOptions(t *testing.T, srv *Server, cfg *gossh.ClientConfig, options ...Option) (*gossh.Session, *gossh.Client, func()) {

}
