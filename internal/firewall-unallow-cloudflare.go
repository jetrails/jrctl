package internal

import (
	"fmt"
	"strings"

	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/firewall"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

var firewallUnAllowCloudflareCmd = &cobra.Command{
	Use:   "cloudflare",
	Short: "Remove allow entries for Cloudflare IP addresses",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Remove allow entries for Cloudflare IP addresses.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl firewall unallow cloudflare",
		"jrctl firewall unallow cloudflare -t www",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		selector, _ := cmd.Flags().GetString("type")
		rows := [][]string{{"Server", "Response"}}
		runner := func(index, total int, context server.Context) {
			response := firewall.UnAllowCloudflare(context)
			row := []string{
				strings.TrimSuffix(context.Endpoint, ":27482"),
				response.Messages[0],
			}
			rows = append(rows, row)
		}
		server.FilterForEach([]string{selector}, runner)
		if !quiet {
			if len(rows) > 1 {
				fmt.Printf("\nExecuted only on %q server(s):\n", selector)
			}
			text.TablePrint(fmt.Sprintf("No configured %q server(s) found.", selector), rows, 1)
		}
	},
}

func init() {
	firewallUnAllowCmd.AddCommand(firewallUnAllowCloudflareCmd)
	firewallUnAllowCloudflareCmd.Flags().SortFlags = true
	firewallUnAllowCloudflareCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
	firewallUnAllowCloudflareCmd.Flags().StringP("type", "t", "localhost", "specify server type, useful for cluster")
}
