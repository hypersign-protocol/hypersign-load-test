package ssi

type KeyPair struct {
	PublicKeyMultibase string
	PrivateKey string
}

func NewKeyPair() *KeyPair {
	pubKeyMultibase, privKey := generateEd25519Key()

	return &KeyPair{
		PublicKeyMultibase: pubKeyMultibase,
		PrivateKey: privKey,
	}
}