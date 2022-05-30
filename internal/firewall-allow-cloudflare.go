package internal

import (
	"strings"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/firewall"
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
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("type")

		output := NewOutput(quiet, tags)
		output.DisplayServers = true
		output.FailOnNoServers = true
		output.ExitCodeNoServers = 1

		for _, context := range config.GetContexts(tags) {
			response := firewall.AllowCloudflare(context)
			output.AddServer(
				context,
				response.GetGeneric(),
				strings.Join(response.Messages, ","),
			)
		}

		output.Print()
	},
}

func init() {
	firewallAllowCmd.AddCommand(firewallAllowCloudflareCmd)
	firewallAllowCloudflareCmd.Flags().SortFlags = true
	firewallAllowCloudflareCmd.Flags().BoolP("quiet", "q", false, "display no output")
	firewallAllowCloudflareCmd.Flags().StringArrayP("type", "t", []string{"localhost"}, "filter servers using type selectors")
}
