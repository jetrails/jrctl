package internal

import (
	"fmt"
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
		"jrctl daemon list -t www",
	}),
	Run: func ( cmd * cobra.Command, args [] string ) {
		selector, _ := cmd.Flags ().GetString ("type")
		filter := [] string {}
		emptyMsg := "No configured daemons found."
		if selector != "" {
			filter = [] string { selector }
			emptyMsg = fmt.Sprintf ( "No configured daemons found with type %q.", selector )
		}
		rows := [] [] string { [] string { "Daemon", "Server Type(s)", "Status", "Service(s)" } }
		runner := func ( index, total int, context daemon.Context ) {
			response := daemon.ListServices ( context )
			var row [] string
			if response.Code != 200 {
				row = [] string {
					strings.TrimSuffix ( context.Endpoint, ":27482" ),
					strings.Join ( context.Types, ", " ),
					response.Status,
					response.Messages [ 0 ],
				}
			} else {
				row = [] string {
					strings.TrimSuffix ( context.Endpoint, ":27482" ),
					strings.Join ( context.Types, ", " ),
					response.Status,
					strings.Join ( response.Payload, ", " ),
				}
			}
			if row [ 1 ] == "" {
				row [ 1 ] = "-"
			}
			if row [ 3 ] == "" {
				row [ 3 ] = "-"
			}
			rows = append ( rows, row )
		}
		daemon.FilterForEach ( filter, runner )
		utils.TablePrint ( emptyMsg, rows, 1 )
	},
}

func init () {
	daemonCmd.AddCommand ( daemonListCmd )
	daemonListCmd.Flags ().SortFlags = true
	daemonListCmd.Flags ().StringP ( "type", "t", "", "specify daemon type selector" )
}
