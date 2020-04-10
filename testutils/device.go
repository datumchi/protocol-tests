package testutils

import (
	"context"
	"errors"
	"github.com/datumchi/protocol-tests/crypto/ed25519"
	"github.com/datumchi/protocol-tests/crypto/sha256"
	"github.com/datumchi/protocol-tests/encoding"
	"github.com/datumchi/protocol-tests/generated/protocol"
	"github.com/datumchi/protocol-tests/testutils/p"
)

func RegisterDevice(client protocol.IdentityServicesClient, persona *p.Persona) error {

	hashedData := sha256.HashData([]byte(persona.IdentityDomain))
	sig := ed25519.Sign(persona.DevicePrivateKey, hashedData)

	deviceInfo := protocol.DeviceInfo{
		DevicePublicKey:encoding.Encode(persona.DevicePublicKey),
		Signature:encoding.Encode(sig),
	}

	commonResponse, err := client.RegisterDevice(context.Background(), &deviceInfo)
	if err != nil {
		return err
	}

	if !commonResponse.IsOk {
		return errors.New("Unable to register device: " + commonResponse.ExtraInformation)
	}

	return nil

}
