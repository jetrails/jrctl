package internal

import (
	"strings"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/daemon"
)

var daemonListCmd = &cobra.Command {
	Use:   "list",
	Short: "List configured daemons",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"List configured daemons.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl daemon list",
	}),
	Run: func ( cmd * cobra.Command, args [] string ) {
		rows := [] [] string { [] string { "Daemon", "Tag(s)" } }
		filter := [] string {}
		runner := func ( index, total int, context daemon.Context ) {
			row := [] string { context.Endpoint, strings.Join ( context.Tags, ", " ) }
			rows = append ( rows, row )
		}
		daemon.FilterForEach ( filter, runner )
		utils.TablePrint ( "No configured daemons found.", rows, 1 )
	},
}

func init () {
	daemonCmd.AddCommand ( daemonListCmd )
	daemonListCmd.Flags ().SortFlags = false
}
