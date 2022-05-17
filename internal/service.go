package internal

import (
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Interact with services in configured deployment",
}

func init() {
	RootCmd.AddCommand(serviceCmd)
	serviceCmd.Flags().SortFlags = true
}
