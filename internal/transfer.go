package internal

import (
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/spf13/cobra"
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Securely transfer files from one machine to another",
	Example: text.Examples([]string{
		"jrctl transfer send private.png",
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo",
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -o private.png",
	}),
}

func init() {
	RootCmd.AddCommand(transferCmd)
	transferCmd.Flags().SortFlags = true
}
