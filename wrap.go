package ssh

import gossh "golang.org/x/crypto/ssh"

// PublicKey is an abstraction of different types of public keys.
type PublicKey interface {
	gossh.PublicKey
}


// The Permissions type holds fine-grained permissions that are specific to a user
// or a specific authentication method for a user.
// Permissions, except for "source-address", must be enforced in the server application layer,
// after successful authentication.
type Permissions struct {
	gossh.Permissions
}
