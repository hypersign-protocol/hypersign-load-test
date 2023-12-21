package ssi

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"

	ldcontext "github.com/hypersign-protocol/hid-node/x/ssi/ld-context"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/btcsuite/btcutil/base58"
	"github.com/multiformats/go-multibase"
)

func getSignature(privateKey string, message []byte) (string, error) {
	// Decode key into bytes
	privKeyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", err
	}

	// Sign Message
	signatureBytes := ed25519.Sign(privKeyBytes, message)

	return multibase.Encode(multibase.Base58BTC, signatureBytes)
}

func getEd25519Signature(doc types.SsiMsg, docProof *types.DocumentProof, privateKey string) string {
	docBytes, err := ldcontext.Ed25519Signature2020Normalize(doc, docProof)
	if err != nil {
		panic(err)
	}

	signature, err := getSignature(privateKey, docBytes)
	if err != nil {
		panic(err)
	}

	return signature
}

func generateEd25519Key() (string, string) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	// A 2-byte header must be prefixed before Ed25519 public key based on
	// W3C's Ed25519VerificationKey2020 Specification for the attribute `publicKeyMultibase`
	// Read more: https://www.w3.org/community/reports/credentials/CG-FINAL-di-eddsa-2020-20220724/#ed25519verificationkey2020
	var publicKeyWithHeader []byte
	publicKeyWithHeader = append(publicKeyWithHeader, append([]byte{0xed, 0x01}, pubKey...)...)

	pubKeyMultibase := "z" + base58.Encode(publicKeyWithHeader)
	// W3C's Ed25519VerificationKey2020 Specification has not explicitly mentioned about the encoding of the private key
	// or whether it should be prefixed similar to the public key. For now, the encoding of private key remains Base64
	// with no prefix
	privKeyBase64 := base64.StdEncoding.EncodeToString(privKey)

	return pubKeyMultibase, privKeyBase64
}
