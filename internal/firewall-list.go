package internal

import (
	"fmt"
	"strings"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/firewall"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

func extractFirewallComment(input string) string {
	commentEnd := strings.Index(input, " -- ")
	if commentEnd == -1 {
		commentEnd = len(input)
	}
	return strings.ReplaceAll(input[:commentEnd], "_", " ")
}

var firewallListCmd = &cobra.Command{
	Use:   "list",
	Short: "List firewall entries across configured servers",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"List firewall entries across configured servers.",
			"Specifing a server type will only display results for servers of that type.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl firewall list",
		"jrctl firewall list -t admin",
		"jrctl firewall list -t db",
		"jrctl firewall list -t www",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("type")

		output := NewOutput(quiet, tags)
		output.DisplayServers = true
		output.FailOnNoServers = true
		output.FailOnNoResults = true
		output.ExitCodeNoServers = 1
		output.ExitCodeNoResults = 2

		for _, context := range server.GetContexts(tags) {
			response := firewall.List(context)
			output.Servers.AddQuietEntry(
				fmt.Sprintf("%d", len(response.Payload)),
			)
			output.AddServer(
				context,
				response.GetGeneric(),
				response.Messages[0],
			)
			tbl := NewTable(Columns{
				"Hostname",
				"Action",
				"IPV4/CIDR",
				"Port(s)",
				"Protocol(s)",
				"Comment",
			})
			for _, entry := range response.Payload {
				tbl.AddRow(Columns{
					response.Metadata["hostname"],
					entry.Action,
					entry.Source,
					strings.Trim(strings.Join(strings.Fields(fmt.Sprint(entry.Ports)), ", "), "[]"),
					strings.Join(entry.Protocols, ", "),
					fmt.Sprintf("%q", extractFirewallComment(entry.Comment)),
				})
			}
			tbl.SquashOnPivot(4)
			output.AddTable(tbl)
		}

		output.Print()
	},
}

func init() {
	firewallCmd.AddCommand(firewallListCmd)
	firewallListCmd.Flags().SortFlags = true
	firewallListCmd.Flags().BoolP("quiet", "q", false, "display number of entries found for each matching server")
	firewallListCmd.Flags().StringArrayP("type", "t", []string{"localhost"}, "filter servers using type selectors")
}
