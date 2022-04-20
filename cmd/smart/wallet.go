package main
import (
	//"fmt"
	"github.com/spf13/cobra"
)

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

	//walletCmd.AddCommand(walletNewAccountCmd())
	//walletCmd.AddCommand(walletPrintPrivKeyCmd())

	return walletCmd
}

