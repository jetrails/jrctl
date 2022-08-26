package internal

import (
	"github.com/spf13/cobra"
)

var databaseUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage database users in deployment",
}

func init() {
	// OnlyRunOnNonAWS(databaseUserCmd)
	databaseCmd.AddCommand(databaseUserCmd)
	databaseUserCmd.Flags().SortFlags = true
}
