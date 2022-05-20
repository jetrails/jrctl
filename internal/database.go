package internal

import (
	"github.com/spf13/cobra"
)

var databaseCmd = &cobra.Command{
	Use:     "database",
	Aliases: []string{"db"},
	Short:   "Manage databases and database users in deployment",
}

func init() {
	RootCmd.AddCommand(databaseCmd)
	databaseCmd.Flags().SortFlags = true
}
