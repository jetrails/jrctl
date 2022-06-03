package internal

import (
	"github.com/spf13/cobra"
)

var websiteCmd = &cobra.Command{
	Use:     "website",
	Aliases: []string{"site"},
	Short:   "Manage websites in deployment",
}

func init() {
	RootCmd.AddCommand(websiteCmd)
	websiteCmd.Flags().SortFlags = true
}
