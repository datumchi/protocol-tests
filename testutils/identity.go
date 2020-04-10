package testutils

import (
	"context"
	"errors"
	"github.com/datumchi/protocol-tests/encoding"
	"github.com/datumchi/protocol-tests/generated/protocol"
	"github.com/datumchi/protocol-tests/testutils/p"
)

func GenerateValidIdentityUsingStandardHumanAttributes(client protocol.IdentityServicesClient, persona *p.Persona) (protocol.Identity, error) {

	validIdentityAddress := protocol.Address{
		Domain:              persona.IdentityDomain,
		DescriptorReference: persona.Title + encoding.Encode(persona.DevicePublicKey),
	}

	validIdentity := protocol.Identity{
		Address: &validIdentityAddress,
		Attributes: make(map[string]*protocol.Identity_Attribute),
		IdentityPublicKey:encoding.Encode(persona.IdentityPublicKey),
	}

	givenNameAttr, givenNameAttest := utils.GenerateValidSelfAttestedAttribute(&persona.DevicePublicKey, &persona.DevicePrivateKey, "given_name", "Given Name/First Name", "TestFirstName", validIdentityAddress)
	validIdentity.Attributes["given_name"] = &givenNameAttr
	validIdentity.Attestations = append(validIdentity.Attestations, &givenNameAttest)

	surNameAttr, surNameAttest := utils.GenerateValidSelfAttestedAttribute(&persona.DevicePublicKey, &persona.DevicePrivateKey, "sur_name", "Surname/Family Name", "TestSurFamilyName", validIdentityAddress)
	validIdentity.Attributes["sur_name"] = &surNameAttr
	validIdentity.Attestations = append(validIdentity.Attestations, &surNameAttest)

	authToken := persona.AuthTokens[persona.IdentityDomain]
	request := protocol.EstablishIdentityAuthenticatedRequest{
		AuthToken:&authToken,
		Identity:&validIdentity,
	}

	commonResponse, err := client.EstablishIdentity(context.Background(), &request)
	if err != nil {
		return protocol.Identity{}, err
	}
	if !commonResponse.IsOk {
		return protocol.Identity{}, errors.New("Error during ID create: " + commonResponse.ExtraInformation)
	}

	return validIdentity, nil

}

