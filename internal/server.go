package internal

import (
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
)

var serverCmd = &cobra.Command {
	Use:   "server",
	Short: "Interact with configured servers",
	Example: utils.Examples ([] string {
		"jrctl server list -h",
		"jrctl server version -h",
	}),
}

func init () {
	rootCmd.AddCommand ( serverCmd )
	serverCmd.Flags ().SortFlags = true
}
