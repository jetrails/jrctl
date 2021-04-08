package internal

import (
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/spf13/cobra"
)

var secretCmd = &cobra.Command{
	Use:   "secret",
	Short: "Interact with one-time secret service",
	Example: text.Examples([]string{
		"jrctl secret create -h",
		"jrctl secret delete -h",
		"jrctl secret read -h",
	}),
}

func init() {
	RootCmd.AddCommand(secretCmd)
	secretCmd.Flags().SortFlags = true
}
