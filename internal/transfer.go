package internal

import (
	"github.com/spf13/cobra"
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Aliases: []string{"trans"}
	Short: "Securely transfer files from one machine to another",
}

func init() {
	RootCmd.AddCommand(transferCmd)
	transferCmd.Flags().SortFlags = true
}
