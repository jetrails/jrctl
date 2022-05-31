package internal

import (
	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Auxiliary tools for aws deployments",
}

func init() {
	OnlyRunOnAWS(awsCmd)
	RootCmd.AddCommand(awsCmd)
	awsCmd.Flags().SortFlags = true
}
