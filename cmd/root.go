package cmd

import (
	"os"

	"github.com/spf13/cobra"
)


func rootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hypersign-load-test",
		Short: "Load Testing for Hypersign Network",
		Long: "Load Testing for Hypersign Network",
	}

	cmd.AddCommand(createAccountCommand())
	cmd.AddCommand(loadTestCommand())
	cmd.AddCommand(listAccountsCommand())

	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd := rootCommand()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}


