package internal

import (
	"github.com/spf13/cobra"
)

var transferCmd = &cobra.Command{
	Use:     "transfer",
	Aliases: []string{"trans", "file"},
	Short:   "Securely transfer files from one machine to another",
	Hidden:  true,
}

func init() {
	RootCmd.AddCommand(transferCmd)
	transferCmd.Flags().SortFlags = true
}
