package internal

import (
	"fmt"
	"strings"

	"github.com/jetrails/jrctl/sdk/firewall"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/spf13/cobra"
)

var firewallListCmd = &cobra.Command{
	Use:   "list",
	Short: "List firewall entries across configured servers",
	Long: utils.Combine([]string{
		utils.Paragraph([]string{
			"List firewall entries across configured servers.",
			"Specifing a server type will only display results for servers of that type.",
		}),
	}),
	Example: utils.Examples([]string{
		"jrctl firewall list",
		"jrctl firewall list -t admin",
		"jrctl firewall list -t db",
		"jrctl firewall list -t www",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		selector, _ := cmd.Flags().GetString("type")
		responseRows := [][]string{[]string{"Server", "Response"}}
		entryRows := [][]string{[]string{"Server", "Action", "IPV4/CIDR", "Port(s)", "Protocol(s)"}}
		filter := []string{}
		emptyMsg := "No configured servers found."
		if selector != "" {
			filter = []string{selector}
			emptyMsg = fmt.Sprintf("No configured %q server(s) found.", selector)
		}
		runner := func(index, total int, context server.Context) {
			response := firewall.List(context)
			responseRow := []string{
				strings.TrimSuffix(context.Endpoint, ":27482"),
				response.Messages[0],
			}
			responseRows = append(responseRows, responseRow)
			for _, entry := range response.Payload {
				entryRow := []string{
					strings.TrimSuffix(context.Endpoint, ":27482"),
					entry.Action,
					entry.Source,
					strings.Trim(strings.Join(strings.Fields(fmt.Sprint(entry.Ports)), ", "), "[]"),
					strings.Join(entry.Protocols, ", "),
				}
				entryRows = append(entryRows, entryRow)
			}
		}
		server.FilterForEach(filter, runner)
		if selector != "" && len(responseRows) > 1 {
			fmt.Printf("\nDisplaying results for %q server(s):\n", selector)
		}
		fmt.Println()
		utils.TablePrint(emptyMsg, responseRows, 0)
		fmt.Println()
		if len(responseRows) > 1 {
			utils.TablePrint("No firewall entries found.", entryRows, 0)
			fmt.Println()
		}
	},
}

func init() {
	firewallCmd.AddCommand(firewallListCmd)
	firewallListCmd.Flags().SortFlags = true
	firewallListCmd.Flags().StringP("type", "t", "", "specify server type selector")
}
