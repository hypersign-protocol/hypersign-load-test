package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hypersign-protocol/hypersign-load-test/client"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/spf13/cobra"
)

var defaultHome = os.ExpandEnv("$HOME/.hid-node-load-test")

func createAccountCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-account [name]",
		Short: "Create an account to send transactions from",
		Long:  "Create an account to send transactions from",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			accountName := args[0]
			homeDir, err := cmd.Flags().GetString("home")
			if err != nil {
				return err
			}

			account, err := client.CreateNewAccount(homeDir, accountName)
			if err != nil {
				if err == cosmosaccount.ErrAccountExists {
					fmt.Fprintf(cmd.OutOrStdout(), "Account with name '%v' already exists at %v\n", accountName, homeDir)
					os.Exit(1)
				} else {
					return err
				}
			}

			address, err := account.Address("hid")
			if err != nil {
				return err
			}

			_, err = fmt.Fprintf(
				cmd.OutOrStdout(), 
				"\n\nAccount %v (%v) has been created at %v.\nYou need to fund this account with atleast 2.6 HID for the Load test to work (Recommended: 6 HID)\n\n", 
				accountName, 
				address, 
				homeDir,
			)
			return err
		},
	}

	cmd.Flags().String("home", defaultHome, "home directory")

	return cmd
}

func listAccountsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-accounts",
		Short: "List created accounts",
		Long:  "List created accounts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			homeDir, err := cmd.Flags().GetString("home")
			if err != nil {
				return err
			}

			accounts, err :=  client.ListAccounts(homeDir)
			if err != nil {
				return err
			}

			fmt.Fprintln(cmd.OutOrStdout(), "Accounts:")
			for i := 0; i < len(accounts); i++ {
				accountAddress, err := accounts[i].Address("hid")
				if err != nil {
					return err
				}

				output := struct {
					Name string `json:"account_name"`
					Address string `json:"account_address"`
				} {
					Name: accounts[i].Name,
					Address: accountAddress,
				}

				outputJson, err := json.MarshalIndent(output, "", " ")
				if err != nil {
					return err
				}
				
				_, err = fmt.Fprintln(cmd.OutOrStdout(), string(outputJson))
				if err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmd.Flags().String("home", defaultHome, "home directory")

	return cmd
}

func loadTestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start --account <name of account> --cred-status-iter (Optional) <no. of RegisterCredentialStatus Tx to execute>",
		Short: "Start the load test",
		Long:  "Start the load test",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			
			homeDir, err := cmd.Flags().GetString("home")
			if err != nil {
				return err
			}

			accountName, err := cmd.Flags().GetString("account")
			if err != nil {
				return err
			}
			if accountName == "" {
				return fmt.Errorf("--account flag must not be empty")
			}
			
			credStatusIter, err := cmd.Flags().GetUint64("cred-status-iter")
			if err != nil {
				return err
			}

			nodeURL, err := cmd.Flags().GetString("node")
			if err != nil {
				return err
			}

			registry, err := client.GetDefaultRegistry(homeDir)
			if err != nil {
				return err
			}

			account, err := registry.GetByName(accountName)
			if err != nil {
				return err
			}

			cmdContext := cmd.Context()
			cosmosClient := client.CreateClient(cmdContext, homeDir, nodeURL)

			if err := executeTransactions(cmdContext, cosmosClient, &account, credStatusIter); err != nil {
				return nil
			}
			return nil
		},
	}

	cmd.Flags().String("home", defaultHome, "home directory")
	cmd.Flags().String("account", "", "account name (not address) to send transaction from")
	cmd.Flags().Uint64("cred-status-iter", 50000, "number of RegisterCredentialStatus transactions to execute")
	cmd.Flags().String("node", "http://localhost:26657", "Node RPC Interface")

	return cmd
}
