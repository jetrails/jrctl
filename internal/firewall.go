package internal

import (
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/spf13/cobra"
)

var firewallCmd = &cobra.Command{
	Use:   "firewall",
	Short: "Interact with firewall to whitelist IP addresses/ports",
	Example: text.Examples([]string{
		"jrctl firewall allow -h",
		"jrctl firewall list -h",
	}),
}

func init() {
	RootCmd.AddCommand(firewallCmd)
	firewallCmd.Flags().SortFlags = true
}
