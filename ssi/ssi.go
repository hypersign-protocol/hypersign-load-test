package ssi

import (
	ldcontext "github.com/hypersign-protocol/hid-node/x/ssi/ld-context"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func GetCredentialStatusDocument(kp *KeyPair, txAuthor string, didId string, vmId string) *types.MsgRegisterCredentialStatus {
	credStatusId := "vc:hid:testnet:" + generateRandomId()

	credStatusDocument := &types.CredentialStatusDocument{
		Context: []string{
			ldcontext.CredentialStatusContext,
			ldcontext.Ed25519Context2020,
		},
		Id:                       credStatusId,
		Revoked:                  false,
		Suspended:                false,
		Remarks:                  "Lorem Ipsum",
		Issuer:                   didId,
		IssuanceDate:             "2023-08-16T09:37:12Z",
		CredentialMerkleRootHash: "f35c3a4e3f1b8ba54ee3cf59d3de91b8b357f707fdb72a46473b65b46f92f80b",
	}

	credentialStatusProof := &types.DocumentProof{
		Type:               types.Ed25519Signature2020,
		Created:            "2023-08-16T09:37:12Z",
		VerificationMethod: vmId,
		ProofPurpose:       "assertionMethod",
		ClientSpecType:     types.CLIENT_SPEC_TYPE_NONE,
	}

	signature := getEd25519Signature(credStatusDocument, credentialStatusProof, kp.PrivateKey)
	credentialStatusProof.ProofValue = signature

	return &types.MsgRegisterCredentialStatus{
		CredentialStatusDocument: credStatusDocument,
		CredentialStatusProof:    credentialStatusProof,
		TxAuthor:                 txAuthor,
	}
}

func GetDidDocument(kp *KeyPair, creator string) *types.MsgRegisterDID {
	didId := "did:hid:testnet:" + kp.PublicKeyMultibase

	didDoc := &types.DidDocument{
		Context: []string{
			ldcontext.DidContext,
			ldcontext.Ed25519Context2020,
		},
		Id:         didId,
		Controller: []string{didId},
		VerificationMethod: []*types.VerificationMethod{
			{
				Id:                 didId + "#k1",
				Type:               types.Ed25519VerificationKey2020,
				Controller:         didId,
				PublicKeyMultibase: kp.PublicKeyMultibase,
			},
		},
	}

	didDocProof := &types.DocumentProof{
		Type:               types.Ed25519Signature2020,
		Created:            "2022-03-02T15:02:00Z",
		VerificationMethod: didDoc.VerificationMethod[0].Id,
		ProofPurpose:       "assertionMethod",
		ClientSpecType:     types.CLIENT_SPEC_TYPE_NONE,
	}

	signature := getEd25519Signature(didDoc, didDocProof, kp.PrivateKey)
	didDocProof.ProofValue = signature

	return &types.MsgRegisterDID{
		DidDocument: didDoc,
		DidDocumentProofs: []*types.DocumentProof{
			didDocProof,
		},
		TxAuthor: creator,
	}
}
