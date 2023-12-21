package client

import (
	"context"
	"fmt"

	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

func CreateClient(ctx context.Context, home string) *cosmosclient.Client {
	addressPrefix := "hid"

	client, err := cosmosclient.New(
		ctx,
		cosmosclient.WithAddressPrefix(addressPrefix),
		cosmosclient.WithHome(home),
		cosmosclient.WithGas("400000"),
		cosmosclient.WithFees("50uhid"),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to create client: %v", err.Error()))
	}

	return &client
}
