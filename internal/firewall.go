package internal

import (
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
)

var firewallCmd = &cobra.Command {
	Use:   "firewall",
	Short: "Interact with firewall to whitelist IP addresses/ports",
	Example: utils.Examples ([] string {
		"jrctl firewall list",
		"jrctl firewall allow -a 1.1.1.1 -p 80 -p 443",
		"jrctl firewall allow -a 1.1.1.1 -p 80,443 -b me",
		"jrctl firewall allow -a 1.1.1.1 -p 80,443 -b me -c 'Office'",
	}),
}

func init () {
	rootCmd.AddCommand ( firewallCmd )
}
