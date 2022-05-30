package internal

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:    "config",
	Short:  "Information about config file",
	Hidden: true,
}

func init() {
	RootCmd.AddCommand(configCmd)
	configCmd.Flags().SortFlags = true
}
