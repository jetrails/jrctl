package internal

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/server"
)

var serverListCmd = &cobra.Command {
	Use:   "list",
	Short: "List configured servers",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"List configured servers.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl server list",
		"jrctl server list -t www",
	}),
	Run: func ( cmd * cobra.Command, args [] string ) {
		selector, _ := cmd.Flags ().GetString ("type")
		filter := [] string {}
		emptyMsg := "No configured servers found."
		if selector != "" {
			filter = [] string { selector }
			emptyMsg = fmt.Sprintf ( "No configured servers found with type %q.", selector )
		}
		rows := [] [] string { [] string { "Server", "Type(s)", "Status", "Service(s)" } }
		runner := func ( index, total int, context server.Context ) {
			response := server.ListServices ( context )
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
		server.FilterForEach ( filter, runner )
		utils.TablePrint ( emptyMsg, rows, 1 )
	},
}

func init () {
	serverCmd.AddCommand ( serverListCmd )
	serverListCmd.Flags ().SortFlags = true
	serverListCmd.Flags ().StringP ( "type", "t", "", "specify server type selector" )
}
