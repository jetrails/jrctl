package internal

import (
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/website"
	"github.com/spf13/cobra"
)

var websitePhpSwitchCmd = &cobra.Command{
	Use:     "php-switch WEBSITE_NAME PHP_VERSION",
	Aliases: []string{"switch-php"},
	Short:   "Switch php-fpm version for website",
	Args:    cobra.ExactArgs(2),
	Example: text.Examples([]string{
		"jrctl website php-switch example.com php-fpm-7.4",
		"jrctl website php-switch example.com php-fpm-7.4 -q",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("tag")
		name := args[0]
		version := args[1]

		output := NewOutput(quiet, tags)
		contexts := config.GetContexts(tags)

		output.PrintTags()
		output.PrintDivider()

		if len(contexts) < 1 {
			output.ExitWithMessage(1, ErrMsgNoServers+"\n")
		}

		if len(contexts) > 1 {
			output.ExitWithMessage(5, ErrMsgRequiresOneServer+"\n")
		}

		request := website.PhpSwitchRequest{Name: name, Version: version}
		response := website.SwitchPHP(contexts[0], request)
		generic := response.GetGeneric()

		output.PrintResponse(generic)
		output.PrintDivider()
		output.ExitCodeFromResponse(generic)
	},
}

func init() {
	websiteCmd.AddCommand(websitePhpSwitchCmd)
	websitePhpSwitchCmd.Flags().SortFlags = true
	websitePhpSwitchCmd.Flags().BoolP("quiet", "q", false, "display no output")
	websitePhpSwitchCmd.Flags().StringArrayP("tag", "t", []string{"localhost"}, "filter nodes using tags")
}
