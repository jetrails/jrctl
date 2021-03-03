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
			"Ask daemon(s) for a list of firewall entries.",
			"Specifing a tag selector will only query daemons with that tag.",
			"Not specifing any tag will show query all configured daemons.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl firewall list",
		"jrctl firewall list -t admin",
		"jrctl firewall list -t mysql",
	}),
	RunE: func ( cmd * cobra.Command, args [] string ) error {
		tag, _ := cmd.Flags ().GetString ("tag")
		if error := daemon.IsValidTagError ( tag ); tag != "" && error != nil {
			return error
		}
		cmd.Run ( cmd, args )
		return nil
	},
	Run: func ( cmd * cobra.Command, args [] string ) {
		tag, _ := cmd.Flags ().GetString ("tag")
		responseRows := [] [] string { [] string { "Daemon", "Status", "Response" } }
		entryRows := [] [] string { [] string { "Daemon", "IPV4/CIDR", "Port(s)" } }
		filter := [] string { tag }
		if tag == "" {
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
	firewallListCmd.Flags ().StringP ( "tag", "t", "", "specify daemon tag selector" )
}
