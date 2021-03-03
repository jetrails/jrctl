package internal

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/firewall"
	"github.com/jetrails/jrctl/sdk/daemon"
)

var firewallListCmd = &cobra.Command {
	Use:   "list",
	Short: "List firewall entries",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"List firewall entries.",
			"Ask the daemon for a list of firewall entries.",
			"Specifing the service will only return results with that service.",
			"Not specifing any service will show everything available.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl firewall list",
		"jrctl firewall list -s admin",
		"jrctl firewall list -s mysql",
	}),
	RunE: func ( cmd * cobra.Command, args [] string ) error {
		service, _ := cmd.Flags ().GetString ("service")
		if error := daemon.IsValidServiceError ( service ); service != "" && error != nil {
			return error
		}
		cmd.Run ( cmd, args )
		return nil
	},
	Run: func ( cmd * cobra.Command, args [] string ) {
		service, _ := cmd.Flags ().GetString ("service")
		responseRows := [] [] string { [] string { "Daemon", "Status", "Response" } }
		entryRows := [] [] string { [] string { "Daemon", "IPV4/CIDR", "Port(s)" } }
		filter := [] string { service }
		if service == "" {
			filter = [] string {}
		}
		runner := func ( index, total int, context daemon.Context ) {
			response := firewall.List ( context )
			responseRow := [] string {
				context.Endpoint,
				strconv.Itoa ( response.Code ),
				response.Messages [ 0 ],
			}
			responseRows = append ( responseRows, responseRow )
			for _, entry := range response.Payload {
				entryRow := [] string {
					context.Endpoint,
					entry.Address,
					strings.Trim(strings.Join(strings.Fields(fmt.Sprint(entry.Port)), ", "), "[]"),
				}
				entryRows = append ( entryRows, entryRow )
			}
		}
		daemon.FilterForEach ( filter, runner )
		utils.TablePrint ( "No configured daemons found.", responseRows, 1 )
		if len ( responseRows ) > 1 {
			utils.TablePrint ( "No firewall entries found.", entryRows, 1 )
		}
	},
}

func init () {
	firewallCmd.AddCommand ( firewallListCmd )
	firewallListCmd.Flags ().SortFlags = false
	firewallListCmd.Flags ().StringP ( "service", "s", "", "filter by service" )
}
