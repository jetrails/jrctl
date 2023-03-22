package internal

import (
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/website"
	"github.com/spf13/cobra"
)

var websiteListCmd = &cobra.Command{
	Use:   "list",
	Short: "Display configured websites",
	Example: text.Examples([]string{
		"jrctl website list",
		"jrctl website list -q",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("tag")

		output := NewOutput(quiet, tags)
		output.DisplayServers = true
		output.FailOnNoServers = true
		output.FailOnNoResults = true
		output.ExitCodeNoServers = 1
		output.ExitCodeNoResults = 2

		tbl := output.CreateTable(Columns{
			"Name",
			"Compatible",
			"PHP-FPM Version",
		})

		compatibilityMap := map[bool]string{true: "Yes", false: "No"}

		for _, context := range config.GetContexts(tags) {
			response := website.ListWebsites(context)
			output.AddServer(
				context,
				response.GetGeneric(),
				response.GetFirstMessage(),
			)
			for _, entry := range response.Payload {
				tbl.AddUniqueQuietEntry(entry.Name)
				tbl.AddRow(Columns{
					entry.Name,
					compatibilityMap[entry.Compatible],
					entry.PHPVersion,
				})
			}
		}

		output.Print()
	},
}

func init() {
	websiteCmd.AddCommand(websiteListCmd)
	websiteListCmd.Flags().SortFlags = true
	websiteListCmd.Flags().BoolP("quiet", "q", false, "only display site names")
	websiteListCmd.Flags().StringArrayP("tag", "t", []string{"localhost"}, "filter nodes using tags")
}
