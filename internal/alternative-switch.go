package internal

import (
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/alternative"
	"github.com/spf13/cobra"
)

var alternativeSwitchCmd = &cobra.Command{
	Use:     "switch PROGRAM",
	Aliases: []string{"update"},
	Short:   "Switch current version of program",
	Args:    cobra.ExactArgs(1),
	Example: text.Examples([]string{
		"jrctl alternative switch php-cli -v php-cli-8.0",
		"jrctl alternative switch php-cli -v php-cli-8.0 -q",
		"jrctl alternative switch php-cli -v php-cli-8.0 -t www",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("tag")
		version, _ := cmd.Flags().GetString("version")
		name := args[0]

		output := NewOutput(quiet, tags)
		output.DisplayServers = true
		contexts := config.GetContexts(tags)

		if len(contexts) < 1 {
			output.PrintDivider()
			output.ExitWithMessage(1, ErrMsgNoServers+"\n")
		}

		for _, context := range config.GetContexts(tags) {
			request := alternative.SwitchRequest{Name: name, Version: version}
			response := alternative.Switch(context, request)
			output.AddServer(
				context,
				response.GetGeneric(),
				response.GetFirstMessage(),
			)
		}

		output.Print()
	},
}

func init() {
	alternativeCmd.AddCommand(alternativeSwitchCmd)
	alternativeSwitchCmd.Flags().SortFlags = true
	alternativeSwitchCmd.Flags().BoolP("quiet", "q", false, "display no output")
	alternativeSwitchCmd.Flags().StringArrayP("tag", "t", []string{"default"}, "filter nodes using tags")
	alternativeSwitchCmd.Flags().StringP("version", "v", "", "version to switch to")
	alternativeSwitchCmd.MarkFlagRequired("version")
}
