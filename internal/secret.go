package internal

import "github.com/spf13/cobra"

var secretCmd = &cobra.Command {
	Use:   "secret",
	Short: "Interact with our one-time secret service",
}

func init () {
	rootCmd.AddCommand ( secretCmd )
}
