package internal

import (
	"github.com/spf13/cobra"
)

var secretCmd = &cobra.Command{
	Use:   "secret",
	Short: "Interact with one-time secret service",
}

func init() {
	RootCmd.AddCommand(secretCmd)
	secretCmd.Flags().SortFlags = true
}
