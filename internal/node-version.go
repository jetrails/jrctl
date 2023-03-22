package internal

import (
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

var nodeVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display daemon version running on configured nodes",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Display daemon version running on configured nodes.",
			"Specifing a node type will only display results for nodes of that type.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl node version",
		"jrctl node version -t www",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("tag")

		output := NewOutput(quiet, tags)
		output.DisplayServers = false
		output.FailOnNoServers = true
		output.FailOnNoResults = true
		output.ExitCodeNoServers = 1
		output.ExitCodeNoResults = 2

		tbl := output.CreateTable(Columns{
			"Hostname",
			"Server",
			"Version",
		})

		for _, context := range config.GetContexts(tags) {
			response := server.Version(context)
			output.AddServer(
				context,
				response.GetGeneric(),
				response.Status,
			)
			if response.IsOkay() {
				tbl.AddQuietEntry(response.Payload)
				tbl.AddRow(Columns{
					response.Metadata["hostname"],
					context.Endpoint,
					response.Payload,
				})
			} else {
				tbl.AddRow(Columns{
					response.Metadata["hostname"],
					context.Endpoint,
				})
			}
		}

		output.Print()
	},
}

func init() {
	nodeCmd.AddCommand(nodeVersionCmd)
	nodeVersionCmd.Flags().SortFlags = true
	nodeVersionCmd.Flags().BoolP("quiet", "q", false, "only display versions")
	nodeVersionCmd.Flags().StringArrayP("tag", "t", []string{}, "filter nodes using tags")
}
