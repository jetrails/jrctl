package internal

import (
	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate reports for deployment",
}

func init() {
	RootCmd.AddCommand(reportCmd)
	reportCmd.Flags().SortFlags = true
}
