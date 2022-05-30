package internal

import (
	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:    "aws",
	Short:  "Auxiliary tools for aws deployments",
	Hidden: true,
}

func init() {
	RootCmd.AddCommand(awsCmd)
	awsCmd.Flags().SortFlags = true
}
