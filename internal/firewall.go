package internal

import (
	"github.com/spf13/cobra"
)

var firewallCmd = &cobra.Command{
	Use:   "firewall",
	Short: "Interact with firewall to whitelist IP addresses/ports",
}

func init() {
	RootCmd.AddCommand(firewallCmd)
	firewallCmd.Flags().SortFlags = true
}
