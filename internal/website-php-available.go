package internal

import (
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/website"
	"github.com/spf13/cobra"
)

var websitePhpAvailableCmd = &cobra.Command{
	Use:     "php-available",
	Aliases: []string{"available-php", "list-php"},
	Short:   "List available php-fpm versions that are available for websites to use",
	Args:    cobra.ExactArgs(0),
	Example: text.Examples([]string{
		"jrctl website php-available",
		"jrctl website php-available -q",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("tag")

		output := NewOutput(quiet, tags)
		output.FailOnNoResults = true
		output.ExitCodeNoResults = 2

		tbl := output.CreateTable(Columns{
			"Version",
			"Installed",
			"Enabled",
			"Configured",
			"Available",
		})

		contexts := config.GetContexts(tags)

		if len(contexts) < 1 {
			output.PrintDivider()
			output.ExitWithMessage(1, ErrMsgNoServers+"\n")
		}

		if len(contexts) > 1 {
			output.PrintTags()
			output.PrintDivider()
			output.ExitWithMessage(3, ErrMsgRequiresOneServer+"\n")
		}

		response := website.AvailablePHP(contexts[0])
		generic := response.GetGeneric()

		if !response.IsOkay() {
			output.PrintTags()
			output.PrintDivider()
			output.PrintResponse(generic)
			output.PrintDivider()
			output.ExitCodeFromResponse(generic)
		}

		output.AddServer(
			contexts[0],
			generic,
			response.GetFirstMessage(),
		)
		for _, availability := range response.Payload {
			boolMapping := map[bool]string{true: "✔", false: " "}
			availableMapping := map[bool]string{true: "Yes", false: "No"}
			available := availability.Installed && availability.Enabled && availability.Configured
			tbl.AddRow(Columns{
				availability.Name,
				boolMapping[availability.Installed],
				boolMapping[availability.Enabled],
				boolMapping[availability.Configured],
				availableMapping[available],
			})
			if available {
				tbl.AddUniqueQuietEntry(availability.Name)
			}
		}

		output.Print()
	},
}

func init() {
	websiteCmd.AddCommand(websitePhpAvailableCmd)
	websitePhpAvailableCmd.Flags().SortFlags = true
	websitePhpAvailableCmd.Flags().BoolP("quiet", "q", false, "display only available php-fpm versions")
	websitePhpAvailableCmd.Flags().StringArrayP("tag", "t", []string{"default"}, "filter nodes using tags")
}
