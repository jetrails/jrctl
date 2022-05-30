package internal

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jetrails/jrctl/pkg/text"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configPrintCmd = &cobra.Command{
	Use:     "print",
	Short:   "Print used config file",
	Aliases: []string{"show"},
	Example: text.Examples([]string{
		"jrctl config print",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		filepath := viper.ConfigFileUsed()
		if bytes, err := ioutil.ReadFile(filepath); err == nil {
			fmt.Print(string(bytes))
		} else {
			fmt.Println("\ncould not read contents of file %q\n\n", filepath)
			os.Exit(2)
		}
	},
}

func init() {
	configCmd.AddCommand(configPrintCmd)
	configPrintCmd.Flags().SortFlags = true
}
