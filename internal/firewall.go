package internal

import (
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
)

var firewallCmd = &cobra.Command {
	Use:   "firewall",
	Short: "Interact with firewall to whitelist IP addresses/ports",
	Example: utils.Examples ([] string {
		"jrctl firewall list -h",
		"jrctl firewall allow -h",
	}),
}

func init () {
	rootCmd.AddCommand ( firewallCmd )
	firewallCmd.Flags ().SortFlags = false
}
