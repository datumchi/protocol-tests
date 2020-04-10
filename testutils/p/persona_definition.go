package p

import (
	"github.com/datumchi/protocol-tests/generated/protocol"
	"golang.org/x/crypto/ed25519"
)

type Persona struct {

	Title            string
	DevicePublicKey  ed25519.PublicKey
	DevicePrivateKey ed25519.PrivateKey

	IdentityDomain     string
	Identity           protocol.Identity
	IdentityPublicKey  ed25519.PublicKey
	IdentityPrivateKey ed25519.PrivateKey

	AuthTokens map[string]protocol.AuthenticationToken

}

