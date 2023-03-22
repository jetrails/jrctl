package internal

import (
	"strings"

	"github.com/jetrails/jrctl/pkg/input"
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var nodeDeleteCmd = &cobra.Command{
	Use:   "delete ENDPOINT",
	Short: "Delete node from config",
	Example: text.Examples([]string{
		"jrctl node delete 127.0.0.1:27482",
		"jrctl node delete 127.0.0.1:27482 -f",
	}),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := args[0]
		quiet, _ := cmd.Flags().GetBool("quiet")
		force, _ := cmd.Flags().GetBool("force")

		output := NewOutput(quiet, []string{})
		output.DisplayServers = false

		tbl := NewTable(Columns{
			"Endpoint",
			"Tag(s)",
			"Action",
		})

		savedNodes := []config.Entry{}
		newNodes := []config.Entry{}
		if err := viper.UnmarshalKey("nodes", &savedNodes); err != nil {
			output.ExitWithMessage(4, "\nfailed to parse current config file\n")
		}
		for _, savedNode := range savedNodes {
			if savedNode.Endpoint == endpoint {
				if !force && (quiet || !input.PromptYesNo("\nare you sure you want to delete this node")) {
					tbl.ExitWithMessage(3, "\nskipping, did not delete node\n")
				}
				tbl.AddRow(Columns{
					savedNode.Endpoint,
					strings.Join(savedNode.Tags, ", "),
					"Deleted",
				})
			} else {
				newNodes = append(newNodes, savedNode)
			}
		}

		if !tbl.IsEmpty() {
			viper.Set("nodes", newNodes)
			viper.WriteConfig()
			output.AddTable(tbl)
			output.Print()
		} else {
			output.ExitWithMessage(1, "\ncould not find any matching nodes\n")
		}
	},
}

func init() {
	nodeCmd.AddCommand(nodeDeleteCmd)
	nodeDeleteCmd.Flags().SortFlags = true
	nodeDeleteCmd.Flags().BoolP("quiet", "q", false, "output only errors")
	nodeDeleteCmd.Flags().BoolP("force", "f", false, "delete without confirmation")
}
