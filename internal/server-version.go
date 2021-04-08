package internal

import (
	"fmt"
	"strings"

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
		selector, _ := cmd.Flags().GetString("type")
		filter := []string{}
		emptyMsg := "No configured servers found."
		if selector != "" {
			filter = []string{selector}
			emptyMsg = fmt.Sprintf("No configured %q server(s) found.", selector)
		}
		rows := [][]string{[]string{"Server", "Version"}}
		runner := func(index, total int, context server.Context) {
			response := server.Version(context)
			var row []string
			if response.Code != 200 {
				row = []string{
					strings.TrimSuffix(context.Endpoint, ":27482"),
					response.Messages[0],
				}
			} else {
				row = []string{
					strings.TrimSuffix(context.Endpoint, ":27482"),
					response.Payload,
				}
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
	serverCmd.AddCommand(serverVersionCmd)
	serverVersionCmd.Flags().SortFlags = true
	serverVersionCmd.Flags().StringP("type", "t", "", "specify server type selector")
}
