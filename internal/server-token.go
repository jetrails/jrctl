package internal

import (
	"fmt"
	"strings"

	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

var serverTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Displays token information for all configured servers",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Displays token information for all configured servers.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl server token",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		selector, _ := cmd.Flags().GetString("type")
		filter := []string{}
		emptyMsg := "No configured servers found."
		if selector != "" {
			filter = []string{selector}
			emptyMsg = fmt.Sprintf("No configured %q server(s) found.", selector)
		}
		rows := [][]string{[]string{"Server", "Token ID", "Identity", "Allowed Client IPs"}}
		runner := func(index, total int, context server.Context) {
			response := server.TokenInfo(context)
			var row []string
			if response.Code != 200 {
				row = []string{
					strings.TrimSuffix(context.Endpoint, ":27482"),
					response.Messages[0],
				}
			} else {
				row = []string{
					strings.TrimSuffix(context.Endpoint, ":27482"),
					response.Payload.TokenID,
					response.Payload.Identity,
					strings.Join(response.Payload.AllowedClientIPs, ", "),
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
	serverCmd.AddCommand(serverTokenCmd)
	serverTokenCmd.Flags().SortFlags = true
	serverTokenCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
	serverTokenCmd.Flags().StringP("type", "t", "", "specify server type selector")
}
