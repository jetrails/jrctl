package internal

import (
	"strconv"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/daemon"
)

var daemonVersionCmd = &cobra.Command {
	Use:   "version",
	Short: "Display version of configured daemons",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"Display version of configured daemons.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl daemon version",
	}),
	Run: func ( cmd * cobra.Command, args [] string ) {
		rows := [] [] string { [] string { "Daemon", "Status", "Version" } }
		filter := [] string {}
		runner := func ( index, total int, context daemon.Context ) {
			response := daemon.Version ( context )
			var row [] string
			if response.Code != 200 {
				row = [] string {
					context.Endpoint,
					strconv.Itoa ( response.Code ),
					response.Messages [ 0 ],
				}
			} else {
				row = [] string {
					context.Endpoint,
					strconv.Itoa ( response.Code ),
					response.Payload,
				}
			}
			rows = append ( rows, row )
		}
		daemon.FilterForEach ( filter, runner )
		utils.TablePrint ( "No configured daemons found.", rows, 1 )
	},
}

func init () {
	daemonCmd.AddCommand ( daemonVersionCmd )
	daemonVersionCmd.Flags ().SortFlags = false
}
