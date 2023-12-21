package client

import (
	"context"
	"fmt"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

func PerformRegisterDIDTx(
	ctx context.Context, 
	client cosmosclient.Client,
	account cosmosaccount.Account, 
	msg *types.MsgRegisterDID,
) error {
	response, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		return err
	}
	if response.Code != 0 {
		return fmt.Errorf("failed to execute RegisterDID Tx: %v", response.RawLog)
	}

	return nil
}

func PerformRegisterCredentialStatusTx(
	ctx context.Context, 
	client cosmosclient.Client,
	account cosmosaccount.Account, 
	msg *types.MsgRegisterCredentialStatus,
) error {
	response, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		return err
	}
	if response.Code != 0 {
		return fmt.Errorf("failed to execute RegisterCredentialStatus Tx: %v", response.RawLog)
	}
	return nil
}
