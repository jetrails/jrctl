package internal

import (
	"fmt"
	"strings"

	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/firewall"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

var firewallAllowCloudflareCmd = &cobra.Command{
	Use:   "cloudflare",
	Short: "Whitelist Cloudflare IP addresses",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Whitelist Cloudflare IP addresses.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl firewall allow cloudflare",
		"jrctl firewall allow cloudflare -t www",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		selector, _ := cmd.Flags().GetString("type")
		rows := [][]string{[]string{"Server", "Response"}}
		runner := func(index, total int, context server.Context) {
			response := firewall.AllowCloudflare(context)
			row := []string{
				strings.TrimSuffix(context.Endpoint, ":27482"),
				response.Messages[0],
			}
			rows = append(rows, row)
		}
		server.FilterForEach([]string{selector}, runner)
		if len(rows) > 1 {
			fmt.Printf("\nExecuted only on %q server(s):\n", selector)
		}
		text.TablePrint(fmt.Sprintf("No configured %q server(s) found.", selector), rows, 1)
	},
}

func init() {
	firewallAllowCmd.AddCommand(firewallAllowCloudflareCmd)
	firewallAllowCloudflareCmd.Flags().SortFlags = true
	firewallAllowCloudflareCmd.Flags().StringP("type", "t", "localhost", "specify server type, useful for cluster")
}
