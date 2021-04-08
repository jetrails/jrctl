package internal

import (
	"fmt"
	"strings"

	"github.com/jetrails/jrctl/sdk/firewall"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/spf13/cobra"
)

var firewallUnAllowCloudflareCmd = &cobra.Command{
	Use:   "cloudflare",
	Short: "Remove allow entries for Cloudflare IP addresses",
	Long: utils.Combine([]string{
		utils.Paragraph([]string{
			"Remove allow entries for Cloudflare IP addresses.",
		}),
	}),
	Example: utils.Examples([]string{
		"jrctl firewall unallow cloudflare",
		"jrctl firewall unallow cloudflare -t www",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		selector, _ := cmd.Flags().GetString("type")
		rows := [][]string{[]string{"Server", "Response"}}
		runner := func(index, total int, context server.Context) {
			response := firewall.UnAllowCloudflare(context)
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
		utils.TablePrint(fmt.Sprintf("No configured %q server(s) found.", selector), rows, 1)
	},
}

func init() {
	firewallUnAllowCmd.AddCommand(firewallUnAllowCloudflareCmd)
	firewallUnAllowCloudflareCmd.Flags().SortFlags = true
	firewallUnAllowCloudflareCmd.Flags().StringP("type", "t", "localhost", "specify server type, useful for cluster")
}
