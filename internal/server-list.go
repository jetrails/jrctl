package internal

import (
	"strings"

	"github.com/jetrails/jrctl/pkg/array"
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

var serverListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"token"},
	Short:   "Displays token information for all configured servers",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Displays token information for all configured servers.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl server token",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("type")
		identityFilters, _ := cmd.Flags().GetStringSlice("identity")
		tokenIDFilters, _ := cmd.Flags().GetStringSlice("token-id")

		output := NewOutput(quiet, tags)
		output.DisplayServers = false
		output.FailOnNoServers = true
		output.FailOnNoResults = true
		output.ExitCodeNoServers = 1
		output.ExitCodeNoResults = 2

		tbl := output.CreateTable(Columns{
			"Hostname",
			"Server",
			"Type(s)",
			"Token ID",
			"Identity",
			"Allowed Client IPs",
		})

		for _, context := range config.GetContexts(tags) {
			response := server.TokenInfo(context)
			output.AddServer(
				context,
				response.GetGeneric(),
				response.Status,
			)
			if response.IsOkay() {
				if len(identityFilters) > 0 && !array.ContainsString(identityFilters, response.Payload.Identity) {
					continue
				}
				if len(tokenIDFilters) > 0 && !array.ContainsString(tokenIDFilters, response.Payload.TokenID) {
					continue
				}
				tbl.AddQuietEntry(response.Payload.TokenID)
				tbl.AddRow(Columns{
					response.Metadata["hostname"],
					strings.TrimSuffix(context.Endpoint, ":27482"),
					strings.Join(context.Types, ", "),
					response.Payload.TokenID,
					response.Payload.Identity,
					strings.Join(response.Payload.AllowedClientIPs, ", "),
				})
			} else if len(identityFilters) == 0 && len(tokenIDFilters) == 0 {
				tbl.AddRow(Columns{
					response.Metadata["hostname"],
					strings.TrimSuffix(context.Endpoint, ":27482"),
					strings.Join(context.Types, ", "),
				})
			}
		}

		output.Print()
	},
}

func init() {
	serverCmd.AddCommand(serverListCmd)
	serverListCmd.Flags().SortFlags = true
	serverListCmd.Flags().BoolP("quiet", "q", false, "only display versions")
	serverListCmd.Flags().StringArrayP("type", "t", []string{}, "filter servers using type selectors")
	serverListCmd.Flags().StringSliceP("identity", "i", []string{}, "filter with identity")
	serverListCmd.Flags().StringSliceP("token-id", "I", []string{}, "filter with token id")
}
