package internal

import (
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/spf13/cobra"
)

var firewallCmd = &cobra.Command{
	Use:   "firewall",
	Short: "Interact with firewall to whitelist IP addresses/ports",
	Example: utils.Examples([]string{
		"jrctl firewall allow -h",
		"jrctl firewall list -h",
	}),
}

func init() {
	rootCmd.AddCommand(firewallCmd)
	firewallCmd.Flags().SortFlags = true
}
