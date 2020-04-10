package testutils

import (
	"crypto/rand"
	"github.com/datumchi/protocol-tests/crypto/ed25519"
	"github.com/datumchi/protocol-tests/crypto/sha256"
	"github.com/datumchi/protocol-tests/encoding"
	"github.com/datumchi/protocol-tests/generated/protocol"
)

func GenerateValidIdentity() protocol.Identity {

	publicKey, privateKey := ed25519.GenerateKeypair(rand.Reader)
	return GenerateValidIdentityWithKeys(publicKey, privateKey)

}



func GenerateValidSelfAttestedAttribute(publicKey []byte, privateKey []byte, attributeName string, attributeDescription string, attributeValue string, attestorAddress protocol.Address) (protocol.Identity_Attribute, protocol.Identity_Attestation) {

	publicKeyEncoded := encoding.Encode(publicKey)
	attributeValueHashEncoded := encoding.Encode(sha256.HashData([]byte(attributeValue)))
	identityAttribute := protocol.Identity_Attribute{
		Name:        attributeName,
		ValueHash:   attributeValueHashEncoded,
		Description: attributeDescription,
	}

	attributeAttestationSignature := ed25519.Sign(privateKey, sha256.HashData([]byte(publicKeyEncoded + attributeName + attributeValueHashEncoded)))
	identityAttributeAttestation := protocol.Identity_Attestation{
		AttributeName:      attributeName,
		Attestor: &attestorAddress,
		Attestation:        encoding.Encode(attributeAttestationSignature),
	}

	return identityAttribute, identityAttributeAttestation


}


func GenerateValidIdentityWithKeys(publicKey []byte, privateKey []byte) protocol.Identity {

	var identityAttestationList []*protocol.Identity_Attestation
	encodedPublicKey := encoding.Encode(publicKey)

	address := protocol.Address{
		Domain: "alpha.datumchi.com",
		DescriptorReference:"test" + encodedPublicKey + "@datumchi.com",
	}

	givenNameAttr, givenNameAttestation := GenerateValidSelfAttestedAttribute(publicKey, privateKey, "given_name", "Given Name/First Name", "Test", address)
	surNameAttr, surNameAttestation := GenerateValidSelfAttestedAttribute(publicKey, privateKey, "sur_name", "Surname/Last Name/Family Name", "User", address)

	attributeMap := map[string]*protocol.Identity_Attribute {
		"given_name": &givenNameAttr,
		"sur_name": &surNameAttr,
	}

	identityAttestationList = append(identityAttestationList, &givenNameAttestation)
	identityAttestationList = append(identityAttestationList, &surNameAttestation)

	identity := protocol.Identity{
		Address:           &address,
		IdentityPublicKey: encodedPublicKey,
		Attributes:        attributeMap,
		Attestations:      identityAttestationList,
	}

	return identity

}


func GenerateIdentityWithInvalidAttestation() protocol.Identity {

	var identityAttestationList []*protocol.Identity_Attestation
	publicKey, privateKey := ed25519.GenerateKeypair(rand.Reader)
	encodedPublicKey := encoding.Encode(publicKey)

	address := protocol.Address{
		Domain: "alpha.datumchi.com",
		DescriptorReference:"test@datumchi.com",
	}

	givenNameHash := encoding.Encode(sha256.HashData([]byte("Test")))
	givenNameAttr := protocol.Identity_Attribute{
		Name:        "given_name",
		ValueHash:   givenNameHash,
		Description: "Given Name",
	}
	// !!!!!!!!! Given Name is the invalid attestation
	givenNameAttestationSignature := ed25519.Sign(privateKey, sha256.HashData([]byte(encodedPublicKey + "given_name1" + givenNameHash)))
	givenNameAttestation := protocol.Identity_Attestation{
		AttributeName:      "given_name",
		Attestor: &address,
		Attestation:        encoding.Encode(givenNameAttestationSignature),
	}

	surNameHash := encoding.Encode(sha256.HashData([]byte("User")))
	surNameAttr := protocol.Identity_Attribute{
		Name:        "sur_name",
		ValueHash:   surNameHash,
		Description: "Surname/Last Name/Family Name",
	}
	surNameAttestationSignature := ed25519.Sign(privateKey, sha256.HashData([]byte(encodedPublicKey + "sur_name" + surNameHash)))
	surNameAttestation := protocol.Identity_Attestation{
		AttributeName:      "sur_name",
		Attestor: &address,
		Attestation:        encoding.Encode(surNameAttestationSignature),
	}


	attributeMap := map[string]*protocol.Identity_Attribute {
		"given_name": &givenNameAttr,
		"sur_name": &surNameAttr,
	}

	identityAttestationList = append(identityAttestationList, &givenNameAttestation)
	identityAttestationList = append(identityAttestationList, &surNameAttestation)

	identity := protocol.Identity{
		Address:           &address,
		IdentityPublicKey: encodedPublicKey,
		Attributes:        attributeMap,
		Attestations:      identityAttestationList,
	}

	return identity

}
