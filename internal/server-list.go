package internal

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

var serverListCmd = &cobra.Command{
	Use:   "list",
	Short: "List servers in configured deployment",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"List servers in configured deployment.",
			"Specifing a server type will only display results for servers of that type.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl server list",
		"jrctl server list -t admin",
		"jrctl server list -t localhost",
		"jrctl server list -t www",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		selector, _ := cmd.Flags().GetString("type")
		filter := []string{}
		emptyMsg := "No configured servers found."
		if selector != "" {
			filter = []string{selector}
			emptyMsg = fmt.Sprintf("No configured %q server(s) found.", selector)
		}
		rows := [][]string{[]string{"Server", "Type(s)", "Service(s)"}}
		runner := func(index, total int, context server.Context) {
			sort.Strings(context.Types)
			response := server.ListServices(context)
			var row []string
			if response.Code != 200 {
				row = []string{
					strings.TrimSuffix(context.Endpoint, ":27482"),
					strings.Join(context.Types, ", "),
					response.Messages[0],
				}
			} else {
				row = []string{
					strings.TrimSuffix(context.Endpoint, ":27482"),
					strings.Join(context.Types, ", "),
					strings.Join(response.Payload, ", "),
				}
			}
			if row[1] == "" {
				row[1] = "-"
			}
			if row[2] == "" {
				row[2] = "-"
			}
			rows = append(rows, row)
		}
		server.FilterForEach(filter, runner)
		if selector != "" && len(rows) > 1 {
			fmt.Printf("\nDisplaying results for %q server(s):\n", selector)
		}
		text.TablePrint(emptyMsg, rows, 1)
	},
}

func init() {
	serverCmd.AddCommand(serverListCmd)
	serverListCmd.Flags().SortFlags = true
	serverListCmd.Flags().StringP("type", "t", "", "specify server type selector")
}
