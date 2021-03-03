package internal

import (
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
)

var serviceCmd = &cobra.Command {
	Use:   "service",
	Short: "Interact with system services",
	Example: utils.Examples ([] string {
		"jrctl service restart -h",
	}),
}

func init () {
	rootCmd.AddCommand ( serviceCmd )
	serviceCmd.Flags ().SortFlags = false
}
