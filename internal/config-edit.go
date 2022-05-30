package internal

import (
	"fmt"

	"github.com/jetrails/jrctl/pkg/input"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configEditCmd = &cobra.Command{
	Use:     "edit",
	Short:   "Print edit to used config file",
	Aliases: []string{"mod", "modify"},
	Example: text.Examples([]string{
		"jrctl config edit",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		filepath := viper.ConfigFileUsed()
		if err := input.OpenFileInEditor(filepath); err != nil {
			fmt.Println("\nfailed to edit config file, check EDITOR env var\n")
		}
	},
}

func init() {
	configCmd.AddCommand(configEditCmd)
	configEditCmd.Flags().SortFlags = true
}
