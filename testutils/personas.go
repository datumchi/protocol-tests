package testutils

import (
	"crypto/rand"
	"github.com/datumchi/protocol-tests/crypto/ed25519"
	"github.com/datumchi/protocol-tests/generated/protocol"
	"github.com/datumchi/protocol-tests/testutils/logger"
	"github.com/datumchi/protocol-tests/testutils/p"

	"os"
)

func EstablishValidStandardHumanPersona(identityServiceUrl string, identityDomain string) p.Persona {

	pub, pri := ed25519.GenerateKeypair(rand.Reader)
	persona := p.Persona {
		Title:            "Standard Human Persona",
		DevicePublicKey:  pub,
		DevicePrivateKey: pri,
		IdentityDomain:   identityDomain,
		AuthTokens:make(map[string]protocol.AuthenticationToken),
	}

	identityClient, err := CreateIdentityServicesClient(identityServiceUrl)
	if err != nil {
		logger.Fatalf("Unable to create identity services client:  %v", err)
		os.Exit(1)
	}

	err = RegisterDevice(identityClient, &persona)
	if err != nil {
		logger.Fatalf("Unable to register device:  %v", err)
		os.Exit(1)
	}

	err = AuthenticateDevice(identityClient, &persona)
	if err != nil {
		logger.Fatalf("Unable to authenticate device:  %v", err)
		os.Exit(1)
	}

	ident, err := GenerateValidIdentityUsingStandardHumanAttributes(identityClient, &persona)
	if err != nil {
		logger.Fatalf("Unable to create valid identity using standard human attributes: %v", err)
		os.Exit(1)
	}
	persona.Identity = ident

	return persona

}
