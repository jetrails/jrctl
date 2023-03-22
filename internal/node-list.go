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

var nodeListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"token"},
	Short:   "Displays token information for all configured nodes",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Displays token information for all configured nodes.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl node token",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("tag")
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
			"Tag(s)",
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
					strings.Join(context.Tags, ", "),
					response.Payload.TokenID,
					response.Payload.Identity,
					strings.Join(response.Payload.AllowedClientIPs, ", "),
				})
			} else if len(identityFilters) == 0 && len(tokenIDFilters) == 0 {
				tbl.AddRow(Columns{
					response.Metadata["hostname"],
					strings.TrimSuffix(context.Endpoint, ":27482"),
					strings.Join(context.Tags, ", "),
				})
			}
		}

		output.Print()
	},
}

func init() {
	nodeCmd.AddCommand(nodeListCmd)
	nodeListCmd.Flags().SortFlags = true
	nodeListCmd.Flags().BoolP("quiet", "q", false, "only display versions")
	nodeListCmd.Flags().StringArrayP("tag", "t", []string{}, "filter nodes using tags")
	nodeListCmd.Flags().StringSliceP("identity", "i", []string{}, "filter with identity")
	nodeListCmd.Flags().StringSliceP("token-id", "I", []string{}, "filter with token id")
}
