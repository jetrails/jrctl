package internal

import (
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
)

var daemonCmd = &cobra.Command {
	Use:   "daemon",
	Short: "Interact with configured daemons",
	Example: utils.Examples ([] string {
		"jrctl daemon list -h",
		"jrctl daemon version -h",
	}),
}

func init () {
	rootCmd.AddCommand ( daemonCmd )
	daemonCmd.Flags ().SortFlags = true
}
