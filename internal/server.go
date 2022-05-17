package internal

import (
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Manage servers in deployment",
}

func init() {
	RootCmd.AddCommand(serverCmd)
	serverCmd.Flags().SortFlags = true
}
