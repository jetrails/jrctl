package internal

import (
	"strings"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/firewall"
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
		tags, _ := cmd.Flags().GetStringArray("type")

		output := NewOutput(quiet, tags)
		output.DisplayServers = true
		output.FailOnNoServers = true
		output.ExitCodeNoServers = 1

		for _, context := range config.GetContexts(tags) {
			response := firewall.UnAllowCloudflare(context)
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
	firewallUnAllowCmd.AddCommand(firewallUnAllowCloudflareCmd)
	firewallUnAllowCloudflareCmd.Flags().SortFlags = true
	firewallUnAllowCloudflareCmd.Flags().BoolP("quiet", "q", false, "display no output")
	firewallUnAllowCloudflareCmd.Flags().StringArrayP("type", "t", []string{"localhost"}, "filter servers using type selectors")
}
