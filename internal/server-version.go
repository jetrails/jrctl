package internal

import (
	"strings"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

var serverVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display daemon version running on configured servers",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Display daemon version running on configured servers.",
			"Specifing a server type will only display results for servers of that type.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl server version",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("type")

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

		for _, context := range server.GetContexts(tags) {
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
					strings.TrimSuffix(context.Endpoint, ":27482"),
					response.Payload,
				})
			} else {
				tbl.AddRow(Columns{
					response.Metadata["hostname"],
					strings.TrimSuffix(context.Endpoint, ":27482"),
				})
			}
		}

		output.Print()
	},
}

func init() {
	serverCmd.AddCommand(serverVersionCmd)
	serverVersionCmd.Flags().SortFlags = true
	serverVersionCmd.Flags().BoolP("quiet", "q", false, "only display versions")
	serverVersionCmd.Flags().StringArrayP("type", "t", []string{}, "filter servers using type selectors")
}
