package main

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sonataruby/smart-blockchain/cli"
)

func main() {
	
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
	
}

func walletCmd() *cobra.Command {
	var walletCmd = &cobra.Command{
		Use:   "wallet",
		Short: "Manages blockchain accounts and keys.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return walletCmd
}

func cliCmd() *cobra.Command {
	var cliCmd = &cobra.Command{
		Use:   "cli",
		Short: "Manages blockchain accounts and keys.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
			//cmd := cli.CommandLine{}
			//cmd.Run()
		},
	}
	return cliCmd
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}