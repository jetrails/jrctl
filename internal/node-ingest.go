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
		"echo -n TOKEN | jrctl node ingest -t default",
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
		tags, _ := cmd.Flags().GetStringSlice("tag")
		force, _ := cmd.Flags().GetBool("force")
		endpoint, _ := cmd.Flags().GetString("endpoint")
		tokenValue := input.GetPipeData()

		output := NewOutput(quiet, tags)
		output.DisplayServers = false

		tbl := NewTable(Columns{
			"Endpoint",
			"Tag(s)",
			"Action",
		})

		savedNodes := []config.Entry{}
		if err := viper.UnmarshalKey("nodes", &savedNodes); err != nil {
			output.ExitWithMessage(4, "\nfailed to parse current config file\n")
		}
		for i, savedNode := range savedNodes {
			if savedNode.Endpoint == endpoint {
				savedNodes[i].Token = tokenValue
				savedNodes[i].Tags = tags
				tbl.AddRow(Columns{
					strings.TrimSuffix(savedNode.Endpoint, ":27482"),
					strings.Join(tags, ", "),
					"Updated",
				})
			}
		}
		if tbl.IsEmpty() && force {
			createdEntry := config.Entry{
				Endpoint: endpoint,
				Token:    tokenValue,
				Tags:    tags,
			}
			savedNodes = append(savedNodes, createdEntry)
			tbl.AddRow(Columns{
				strings.TrimSuffix(createdEntry.Endpoint, ":27482"),
				strings.Join(tags, ", "),
				"Created",
			})
		}

		if !tbl.IsEmpty() {
			viper.Set("nodes", savedNodes)
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
	nodeIngestCmd.Flags().StringSliceP("tag", "t", []string{}, "tags to attach to found nodes")
	nodeIngestCmd.Flags().BoolP("force", "f", false, "create new entry if no matching nodes were found")
	nodeIngestCmd.MarkFlagRequired("tag")
}
