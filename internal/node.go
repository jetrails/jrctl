package internal

import (
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Manage configured nodes",
}

func init() {
	RootCmd.AddCommand(nodeCmd)
	nodeCmd.Flags().SortFlags = true
}
