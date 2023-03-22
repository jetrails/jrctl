package internal

import (
	"errors"
	"strings"

	"github.com/jetrails/jrctl/pkg/input"
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var nodeIngestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "Ingest node token and save it to config",
	Example: text.Examples([]string{
		"echo -n TOKEN | jrctl node ingest -t localhost",
		"echo -n TOKEN | jrctl node ingest -t jump -e 10.10.10.7",
		"echo -n TOKEN | jrctl node ingest -t web -e 10.10.10.6 -f",
	}),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !input.HasDataInPipe() {
			return errors.New("must pipe token to stdin")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringSlice("type")
		force, _ := cmd.Flags().GetBool("force")
		endpoint, _ := cmd.Flags().GetString("endpoint")
		tokenValue := input.GetPipeData()

		output := NewOutput(quiet, tags)
		output.DisplayServers = false

		tbl := NewTable(Columns{
			"Endpoint",
			"Type(s)",
			"Action",
		})

		savedServers := []config.Entry{}
		if err := viper.UnmarshalKey("servers", &savedServers); err != nil {
			output.ExitWithMessage(4, "\nfailed to parse current config file\n")
		}
		for i, savedServer := range savedServers {
			if savedServer.Endpoint == endpoint {
				savedServers[i].Token = tokenValue
				savedServers[i].Types = tags
				tbl.AddRow(Columns{
					strings.TrimSuffix(savedServer.Endpoint, ":27482"),
					strings.Join(tags, ", "),
					"Updated",
				})
			}
		}
		if tbl.IsEmpty() && force {
			createdEntry := config.Entry{
				Endpoint: endpoint,
				Token:    tokenValue,
				Types:    tags,
			}
			savedServers = append(savedServers, createdEntry)
			tbl.AddRow(Columns{
				strings.TrimSuffix(createdEntry.Endpoint, ":27482"),
				strings.Join(tags, ", "),
				"Created",
			})
		}

		if !tbl.IsEmpty() {
			viper.Set("servers", savedServers)
			viper.WriteConfig()
			output.AddTable(tbl)
			output.Print()
		} else {
			output.ExitWithMessage(1, "\ncould not find any matching nodes\n")
		}
	},
}

func init() {
	nodeCmd.AddCommand(nodeIngestCmd)
	nodeIngestCmd.Flags().SortFlags = true
	nodeIngestCmd.Flags().BoolP("quiet", "q", false, "output only errors")
	nodeIngestCmd.Flags().StringP("endpoint", "e", "127.0.0.1:27482", "filter nodes using this endpoint")
	nodeIngestCmd.Flags().StringSliceP("type", "t", []string{}, "types to attach to found nodes")
	nodeIngestCmd.Flags().BoolP("force", "f", false, "create new entry if no matching nodes were found")
	nodeIngestCmd.MarkFlagRequired("type")
}
