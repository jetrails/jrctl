package internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jetrails/jrctl/pkg/array"
	"github.com/jetrails/jrctl/pkg/input"
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NormalizeEndpoint(endpoint string) string {
	if strings.Contains(endpoint, ":") {
		return endpoint
	}
	return fmt.Sprintf("%s:27482", endpoint)
}

func NormalizeEndpoints(endpoints []string) []string {
	for i, endpoint := range endpoints {
		endpoints[i] = NormalizeEndpoint(endpoint)
	}
	return endpoints
}

var awsAutoscaleIngestCmd = &cobra.Command{
	Use:     "autoscale-ingest",
	Short:   "Display databases in deployment",
	Example: text.Examples([]string{}),
	Args:    cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !input.HasDataInPipe() {
			return errors.New("must pipe endpoints to stdin")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tag, _ := cmd.Flags().GetString("type")
		endpoints := NormalizeEndpoints(strings.Fields(input.GetPipeData()))
		tags := []string{tag}

		output := NewOutput(quiet, tags)

		if len(endpoints) < 1 {
			output.ExitWithMessage(3, "\nmust pass at least one endpoint\n")
		}

		contexts := config.GetContexts(tags)
		servers := []config.Entry{}
		viper.UnmarshalKey("servers", &servers)

		if !config.ContextsHaveSameToken(contexts) {
			output.PrintTags()
			output.ExitWithMessage(4, "\nfound differing tokens, autoscale requires same tokens\n")
		}

		tbl := output.CreateTable(Columns{
			"Endpoint",
			"Action",
			"Type(s)",
		})

		for _, context := range contexts {
			var action string
			if array.ContainsString(endpoints, context.Endpoint) {
				action = "Skipped"
			} else {
				action = "Deleted"
				filtered := []config.Entry{}
				for _, s := range servers {
					if s.Endpoint != context.Endpoint {
						filtered = append(filtered, s)
					}
				}
				servers = filtered
			}
			tbl.AddRow(Columns{
				context.Endpoint,
				action,
				strings.Join(tags, ", "),
			})
		}

		for _, endpoint := range endpoints {
			if !config.ContextsHaveSomeEndpoint(contexts, []string{endpoint}) {
				tbl.AddRow(Columns{
					endpoint,
					"Created",
					strings.Join(tags, ", "),
				})
				entry := config.Entry{
					Endpoint: endpoint,
					Token:    contexts[0].Token,
					Types:    tags,
				}
				servers = append(servers, entry)
			}
		}

		viper.Set("servers", servers)
		viper.WriteConfig()

		output.Print()
	},
}

func init() {
	OnlyRunOnAWS(awsAutoscaleIngestCmd)
	awsCmd.AddCommand(awsAutoscaleIngestCmd)
	awsAutoscaleIngestCmd.Flags().SortFlags = true
	awsAutoscaleIngestCmd.Flags().BoolP("quiet", "q", false, "output only errors")
	awsAutoscaleIngestCmd.Flags().StringP("type", "t", "www", "filter servers using type selectors, only one selector allowed")
}
