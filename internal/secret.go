package internal

import (
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
)

var secretCmd = &cobra.Command {
	Use:   "secret",
	Short: "Interact with our one-time secret service",
	Example: utils.Examples ([] string {
		"jrctl secret read -h",
		"jrctl secret delete -h",
		"jrctl secret create -h",
	}),
}

func init () {
	rootCmd.AddCommand ( secretCmd )
	secretCmd.Flags ().SortFlags = true
}
