package internal

import (
	"github.com/spf13/cobra"
)

var alternativeCmd = &cobra.Command{
	Use:     "alternative",
	Aliases: []string{"alt","update-alternatives"},
	Short:   "Manage alternative programs",
}

func init() {
	RootCmd.AddCommand(alternativeCmd)
	alternativeCmd.Flags().SortFlags = true
}
