package testutils

import (
	"context"
	"github.com/datumchi/protocol-tests/crypto/ed25519"
	"github.com/datumchi/protocol-tests/crypto/sha256"
	"github.com/datumchi/protocol-tests/encoding"
	"github.com/datumchi/protocol-tests/generated/protocol"
	"github.com/datumchi/protocol-tests/testutils/p"
)

func AuthenticateDevice(client protocol.IdentityServicesClient, persona *p.Persona) error {

	deviceInfo := protocol.DeviceInfo{
		DevicePublicKey: encoding.Encode(persona.DevicePublicKey),
	}

	challengeResponse, err := client.AuthenticateGetChallenge(context.Background(), &deviceInfo)
	if err != nil {
		return err
	}

	dataToSign := encoding.Encode(persona.DevicePublicKey) + persona.IdentityDomain + challengeResponse.ExtraInformation
	hashedChallenge := sha256.HashData([]byte(dataToSign))
	challengeSig := ed25519.Sign(persona.DevicePrivateKey.Seed(), hashedChallenge)
	challengeSigEncoded := encoding.Encode(challengeSig)
	deviceInfo.Signature = challengeSigEncoded
	authToken, err := client.AuthenticateDevice(context.Background(), &deviceInfo)

	persona.AuthTokens[persona.IdentityDomain] = *authToken
	return nil

}
