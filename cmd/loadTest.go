package cmd

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"

	localclient "github.com/hypersign-protocol/hypersign-load-test/client"
	"github.com/hypersign-protocol/hypersign-load-test/ssi"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

type UTCFormatter struct {
    log.Formatter
}

func (u UTCFormatter) Format(e *log.Entry) ([]byte, error) {
    e.Time = e.Time.UTC()
    return u.Formatter.Format(e)
}

func executeTransactions(
	ctx context.Context,
	client *cosmosclient.Client,
	account *cosmosaccount.Account,
	credStatusIterations uint64,
) error {
	logger := log.New()
	logger.SetFormatter(UTCFormatter{&log.TextFormatter{FullTimestamp: true}})
	
	// Generate Ed25519 Key Pair
	ed25519KeyPair := ssi.NewKeyPair()
	txAuthor, err := account.Address("hid")
	if err != nil {
		return err
	}

	// Generate a DID Document
	logger.Printf("Load Test has begun\n")
	msgRegisterDID := ssi.GetDidDocument(ed25519KeyPair, txAuthor)
	didID := msgRegisterDID.DidDocument.Id
	vmId := msgRegisterDID.DidDocument.VerificationMethod[0].Id
	logger.Printf("DID Document %v has been generated\n", didID)

	// Submit a Tx to register this DID Document
	if err := localclient.PerformRegisterDIDTx(ctx, *client, *account, msgRegisterDID); err != nil {
		logger.Errorln(err)
		os.Exit(1)
	}
	logger.Printf("DID Document %v has been registered successfully\n", didID)

	// Loop through (credStatusIterations) times to execute RegisterCredentialStatus Txs
	logger.Printf("Starting the process of submitting %v RegisterCredentialStatus transactions....\n", credStatusIterations)

	for iteration := 1; iteration <= int(credStatusIterations); iteration++ {
		msgRegisterCredentialStatus := ssi.GetCredentialStatusDocument(
			ed25519KeyPair,
			txAuthor,
			didID,
			vmId,
		)

		if err := localclient.PerformRegisterCredentialStatusTx(ctx, *client, *account, msgRegisterCredentialStatus); err != nil {
			logger.Errorf("Iteration %v has failed | Error: %v\n",  iteration, err)
			os.Exit(1)		
		}
		
		logger.Infof("Iteration %v successful | Credential Status %v is registered successfully\n", iteration, msgRegisterCredentialStatus.CredentialStatusDocument.Id)
	}
	return nil
}