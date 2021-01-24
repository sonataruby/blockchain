package main

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	//"github.com/sonataruby/smart-blockchain/cli"
)

func main() {
	var smartCmd = &cobra.Command{
		Use:   "smart",
		Short: "The SMART Blockchain CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	/*
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
	*/
	err := smartCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
