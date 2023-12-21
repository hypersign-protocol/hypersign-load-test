package client

import (
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
)

func GetDefaultRegistry(home string) (*cosmosaccount.Registry, error) {
	registry, err := cosmosaccount.New(
		cosmosaccount.WithHome(home), 
		cosmosaccount.WithKeyringBackend(cosmosaccount.KeyringTest),
	)
	if err != nil {
		return nil, err
	}

	return &registry, nil
}

func CreateNewAccount(home, accountName string) (*cosmosaccount.Account, error) {
	registry, err := GetDefaultRegistry(home)
	if err != nil {
		return nil, err
	}

	account, _, err := registry.Create(accountName)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func ListAccounts(home string) ([]cosmosaccount.Account, error) {
	registry, err := GetDefaultRegistry(home)
	if err != nil {
		return nil, err
	}

	accounts, err := registry.List()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
