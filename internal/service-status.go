package internal

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jetrails/jrctl/pkg/array"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

func convertToAbility(status bool) string {
	if status {
		return "✔"
	}
	return " " // "✗"
}

var serviceStatusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"list"},
	Short:   "List services with their statuses and abilities.",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"List services with their statuses and abilities..",
			"Specifing a server type will only display results for servers of that type.",
			"Specifing the service will filter the list of services to include those services.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl service status",
		"jrctl service status -t admin",
		"jrctl service status -t localhost",
		"jrctl service status -t www",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		selectors, _ := cmd.Flags().GetStringSlice("type")
		serviceSelectors, _ := cmd.Flags().GetStringSlice("service")
		responseRows := [][]string{{"Hostname", "Server", "Type(s)", "Response"}}
		entryRows := [][]string{{"Hostname", "Service", "Status", "Restart", "Reload", "Enable", "Disable"}}
		emptyMsg := fmt.Sprintf("No configured %s server(s) found.", text.QuotedList(selectors))
		runner := func(index, total int, context server.Context) {
			response := server.ListServices(context)
			responseRow := []string{
				response.Metadata["hostname"],
				strings.TrimSuffix(context.Endpoint, ":27482"),
				strings.Join(context.Types, ", "),
				response.Messages[0],
			}
			responseRows = append(responseRows, responseRow)
			keys := make([]string, len(response.Payload))
			i := 0
			for key := range response.Payload {
				keys[i] = key
				i++
			}
			sort.Strings(keys)
			for _, service := range keys {
				properties := response.Payload[service]
				if len(serviceSelectors) == 0 || array.ContainsString(serviceSelectors, service) {
					entryRow := []string{
						response.Metadata["hostname"],
						service,
						properties.Status,
						convertToAbility(properties.Restart),
						convertToAbility(properties.Reload),
						convertToAbility(properties.Enable),
						convertToAbility(properties.Disable),
					}
					entryRows = append(entryRows, entryRow)
				}
			}
		}
		server.FilterForEach(selectors, runner)
		if len(responseRows) > 1 {
			fmt.Printf("\nDisplaying results for %s server(s):\n", text.QuotedList(selectors))
		}
		fmt.Println()
		text.TablePrint(emptyMsg, responseRows, 0)
		fmt.Println()
		text.TablePrint("No services found.", entryRows, 0)
		fmt.Println()
	},
}

func init() {
	serviceCmd.AddCommand(serviceStatusCmd)
	serviceStatusCmd.Flags().SortFlags = true
	serviceStatusCmd.Flags().StringSliceP("type", "t", []string{"localhost"}, "specify server type(s) selector")
	serviceStatusCmd.Flags().StringSliceP("service", "s", []string{}, "filter by service")
}
