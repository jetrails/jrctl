package internal

import (
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate server reports",
	Example: text.Examples([]string{
		"jrctl report audit -h",
	}),
}

func init() {
	RootCmd.AddCommand(reportCmd)
	reportCmd.Flags().SortFlags = true
}
