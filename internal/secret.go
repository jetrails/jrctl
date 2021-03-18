package internal

import (
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
)

var secretCmd = &cobra.Command {
	Use:   "secret",
	Short: "Interact with one-time secret service",
	Example: utils.Examples ([] string {
		"jrctl secret create -h",
		"jrctl secret delete -h",
		"jrctl secret read -h",
	}),
}

func init () {
	rootCmd.AddCommand ( secretCmd )
	secretCmd.Flags ().SortFlags = true
}
