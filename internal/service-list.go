package internal

import (
	"github.com/jetrails/jrctl/pkg/array"
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/service"
	"github.com/spf13/cobra"
)

var serviceListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"status"},
	Short:   "List services with their statuses and abilities.",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"List services with their statuses and abilities..",
			"Specifing a tag will display nodes that have that tag.",
			"Specifing the service will filter the list of services to include those services.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl service list",
		"jrctl service list -t admin",
		"jrctl service list -t default",
		"jrctl service list -t www",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("tag")
		serviceSelectors, _ := cmd.Flags().GetStringSlice("service")

		output := NewOutput(quiet, tags)
		output.DisplayServers = true
		output.FailOnNoServers = true
		output.FailOnNoResults = true
		output.ExitCodeNoServers = 1
		output.ExitCodeNoResults = 2

		tbl := NewTable(Columns{
			"Hostname",
			"Service",
			"Status",
			"Restart",
			"Reload",
			"Enable",
			"Disable",
		})

		statusChar := map[bool]string{
			true:  "✔",
			false: " ",
		}

		for _, context := range config.GetContexts(tags) {
			response := service.ListServices(context)
			output.AddServer(
				context,
				response.GetGeneric(),
				response.GetFirstMessage(),
			)
			for serviceName, properties := range response.Payload {
				if len(serviceSelectors) == 0 || array.ContainsString(serviceSelectors, serviceName) {
					output.Servers.AddUniqueQuietEntry(serviceName)
					tbl.AddRow(Columns{
						response.Metadata["hostname"],
						serviceName,
						properties.Status,
						statusChar[properties.Restart],
						statusChar[properties.Reload],
						statusChar[properties.Enable],
						statusChar[properties.Disable],
					})
				}
			}
		}

		tbl.Sort(1)
		output.AddTable(tbl)
		output.Print()
	},
}

func init() {
	serviceCmd.AddCommand(serviceListCmd)
	serviceListCmd.Flags().SortFlags = true
	serviceListCmd.Flags().BoolP("quiet", "q", false, "display unique list of found services")
	serviceListCmd.Flags().StringArrayP("tag", "t", []string{"default"}, "filter nodes using tags")
	serviceListCmd.Flags().StringSliceP("service", "s", []string{}, "filter by service")
}
