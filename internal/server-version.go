package internal

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/server"
)

var serverVersionCmd = &cobra.Command {
	Use:   "version",
	Short: "Display version of configured servers",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"Display version of configured servers.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl server version",
	}),
	Run: func ( cmd * cobra.Command, args [] string ) {
		selector, _ := cmd.Flags ().GetString ("type")
		filter := [] string {}
		emptyMsg := "No configured servers found."
		if selector != "" {
			filter = [] string { selector }
			emptyMsg = fmt.Sprintf ( "No configured servers found with type %q.", selector )
		}
		rows := [] [] string { [] string { "Server", "Version" } }
		runner := func ( index, total int, context server.Context ) {
			response := server.Version ( context )
			var row [] string
			if response.Code != 200 {
				row = [] string {
					strings.TrimSuffix ( context.Endpoint, ":27482" ),
					response.Messages [ 0 ],
				}
			} else {
				row = [] string {
					strings.TrimSuffix ( context.Endpoint, ":27482" ),
					response.Payload,
				}
			}
			rows = append ( rows, row )
		}
		server.FilterForEach ( filter, runner )
		if selector != "" && len ( rows ) > 1 {
			fmt.Printf ( "\nDisplaying results with server type %q:\n", selector )
		}
		utils.TablePrint ( emptyMsg, rows, 1 )
	},
}

func init () {
	serverCmd.AddCommand ( serverVersionCmd )
	serverVersionCmd.Flags ().SortFlags = true
	serverVersionCmd.Flags ().StringP ( "type", "t", "", "specify server type selector" )
}
