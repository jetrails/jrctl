package internal

import (
	"fmt"
	"strings"
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
		tag, _ := cmd.Flags ().GetString ("tag")
		filter := [] string {}
		emptyMsg := "No configured daemons found."
		if tag != "" {
			filter = [] string { tag }
			emptyMsg = fmt.Sprintf ( "No configured daemons found with tag %q.", tag )
		}
		rows := [] [] string { [] string { "Daemon", "Status", "Version" } }
		runner := func ( index, total int, context daemon.Context ) {
			response := daemon.Version ( context )
			var row [] string
			if response.Code != 200 {
				row = [] string {
					strings.TrimSuffix ( context.Endpoint, ":27482" ),
					response.Status,
					response.Messages [ 0 ],
				}
			} else {
				row = [] string {
					strings.TrimSuffix ( context.Endpoint, ":27482" ),
					response.Status,
					response.Payload,
				}
			}
			rows = append ( rows, row )
		}
		daemon.FilterForEach ( filter, runner )
		utils.TablePrint ( emptyMsg, rows, 1 )
	},
}

func init () {
	daemonCmd.AddCommand ( daemonVersionCmd )
	daemonVersionCmd.Flags ().SortFlags = true
	daemonVersionCmd.Flags ().StringP ( "tag", "t", "", "specify daemon tag selector" )
}
