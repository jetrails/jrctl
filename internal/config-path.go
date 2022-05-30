package internal

import (
	"fmt"

	"github.com/jetrails/jrctl/pkg/text"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configLocationCmd = &cobra.Command{
	Use:     "path",
	Short:   "Print path to used config file",
	Aliases: []string{"loc", "location"},
	Example: text.Examples([]string{
		"jrctl config path",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		filepath := viper.ConfigFileUsed()
		fmt.Println(filepath)
	},
}

func init() {
	configCmd.AddCommand(configLocationCmd)
	configLocationCmd.Flags().SortFlags = true
}
