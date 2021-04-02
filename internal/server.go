package internal

import (
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Interact with servers in configured deployment",
	Example: utils.Examples([]string{
		"jrctl server list -h",
		"jrctl server restart -h",
		"jrctl server version -h",
	}),
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().SortFlags = true
}
