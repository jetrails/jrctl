package internal

import (
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Interact with servers in configured deployment",
	Example: text.Examples([]string{
		"jrctl server list -h",
		"jrctl server restart -h",
		"jrctl server version -h",
	}),
}

func init() {
	RootCmd.AddCommand(serverCmd)
	serverCmd.Flags().SortFlags = true
}
