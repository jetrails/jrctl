package internal

import (
	"fmt"
	"strings"
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/alternative"
	"github.com/spf13/cobra"
)

var alternativeListCmd = &cobra.Command{
	Use:   "list",
	Short: "Display programs with alternative versions",
	Example: text.Examples([]string{
		"jrctl alternative list",
		"jrctl alternative list -q",
		"jrctl alternative list -t www",
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
			"Hostname",
			"Program",
			"Current Version",
			"Available Versions",
		})

		for _, context := range config.GetContexts(tags) {
			response := alternative.List(context)
			output.AddServer(
				context,
				response.GetGeneric(),
				response.GetFirstMessage(),
			)
			for _, entry := range response.Payload {
				tbl.AddUniqueQuietEntry(entry.Current)
				tbl.AddRow(Columns{
					response.Metadata["hostname"],
					entry.Name,
					entry.Current,
					strings.Join(entry.Versions, ", "),
				})
			}
		}

		if !quiet {
			fmt.Println()
			fmt.Println("WARNING: if you are trying to change php-fpm versions, checkout `jrctl website --help` instead.")
		}
		output.Print()
	},
}

func init() {
	alternativeCmd.AddCommand(alternativeListCmd)
	alternativeListCmd.Flags().SortFlags = true
	alternativeListCmd.Flags().BoolP("quiet", "q", false, "only display current versions")
	alternativeListCmd.Flags().StringArrayP("tag", "t", []string{"default"}, "filter nodes using tags")
}
