package internal

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/daemon"
)

var serviceListCmd = &cobra.Command {
	Use:   "list",
	Short: "List available services",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"List available services.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl service list",
		"jrctl service list -t www",
	}),
	Run: func ( cmd * cobra.Command, args [] string ) {
		selector, _ := cmd.Flags ().GetString ("type")
		filter := [] string {}
		emptyMsg := "No configured daemons found."
		if selector != "" {
			filter = [] string { selector }
			emptyMsg = fmt.Sprintf ( "No configured daemons found with type %q.", selector )
		}
		rows := [] [] string { [] string { "Daemon", "Status", "Service(s)" } }
		runner := func ( index, total int, context daemon.Context ) {
			response := daemon.ListServices ( context )
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
					strings.Join ( response.Payload, ", " ),
				}
			}
			if row [ 2 ] == "" {
				row [ 2 ] = "-"
			}
			rows = append ( rows, row )
		}
		daemon.FilterForEach ( filter, runner )
		utils.TablePrint ( emptyMsg, rows, 1 )
	},
}

func init () {
	serviceCmd.AddCommand ( serviceListCmd )
	serviceListCmd.Flags ().SortFlags = true
	serviceListCmd.Flags ().StringP ( "type", "t", "", "specify service type selector" )
}
