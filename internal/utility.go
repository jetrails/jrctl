package internal

import (
	"github.com/spf13/cobra"
)

var utilityCmd = &cobra.Command{
	Use:     "utility",
	Aliases: []string{"utilities", "util", "utils"},
	Short:   "Auxiliary command-line tools",
}

func init() {
	RootCmd.AddCommand(utilityCmd)
	utilityCmd.Flags().SortFlags = true
}
