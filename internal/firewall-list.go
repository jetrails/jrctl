package internal

import (
	"fmt"
	"strings"
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
			"Ask daemon(s) for a list of firewall entries.",
			"Specifing a type selector will only query daemons with that type.",
			"Not specifing any type will show query all configured daemons.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl firewall list",
		"jrctl firewall list -t admin",
		"jrctl firewall list -t db",
	}),
	Run: func ( cmd * cobra.Command, args [] string ) {
		selector, _ := cmd.Flags ().GetString ("type")
		responseRows := [] [] string { [] string { "Daemon", "Status", "Response" } }
		entryRows := [] [] string { [] string { "Daemon", "IPV4/CIDR", "Port(s)" } }
		filter := [] string {}
		emptyMsg := "No configured daemons found."
		if selector != "" {
			filter = [] string { selector }
			emptyMsg = fmt.Sprintf ( "No configured daemons found with type %q.", selector )
		}
		runner := func ( index, total int, context daemon.Context ) {
			response := firewall.List ( context )
			responseRow := [] string {
				strings.TrimSuffix ( context.Endpoint, ":27482" ),
				response.Status,
				response.Messages [ 0 ],
			}
			responseRows = append ( responseRows, responseRow )
			for _, entry := range response.Payload {
				entryRow := [] string {
					strings.TrimSuffix ( context.Endpoint, ":27482" ),
					entry.Address,
					strings.Trim(strings.Join(strings.Fields(fmt.Sprint(entry.Port)), ", "), "[]"),
				}
				entryRows = append ( entryRows, entryRow )
			}
		}
		daemon.FilterForEach ( filter, runner )
		fmt.Println ()
		utils.TablePrint ( emptyMsg, responseRows, 0 )
		fmt.Println ()
		if len ( responseRows ) > 1 {
			utils.TablePrint ( "No firewall entries found.", entryRows, 0 )
			fmt.Println ()
		}
	},
}

func init () {
	firewallCmd.AddCommand ( firewallListCmd )
	firewallListCmd.Flags ().SortFlags = true
	firewallListCmd.Flags ().StringP ( "type", "t", "", "specify daemon type selector" )
}
