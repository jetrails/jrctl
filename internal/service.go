package internal

import (
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:     "service",
	Aliases: []string{"server"},
	Short:   "Interact with services in deployment",
}

func init() {
	RootCmd.AddCommand(serviceCmd)
	serviceCmd.Flags().SortFlags = true
}
