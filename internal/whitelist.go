package internal

import (
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
)

var whitelistCmd = &cobra.Command {
	Use:   "whitelist",
	Short: "Whitelist IP addresses & ports",
	Example: utils.Examples ([] string {
		"jrctl whitelist list",
		"jrctl whitelist add -a 1.1.1.1 -p 80 -p 443",
		"jrctl whitelist add -a 1.1.1.1 -p 80,443 -b me",
		"jrctl whitelist add -a 1.1.1.1 -p 80,443 -b me -c 'Office'",
	}),
}

func init () {
	rootCmd.AddCommand ( whitelistCmd )
}
