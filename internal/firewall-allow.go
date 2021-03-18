package internal

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/firewall"
	"github.com/jetrails/jrctl/sdk/server"
)

var firewallAllowCmd = &cobra.Command {
	Use:   "allow",
	Short: "Permanently allows a source IP address to a specific port",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"Allows a specified IP address to bypass the local system firewall by creating an 'allow' entry into the permanent firewall config.",
			"Grants unprivileged users ability to manipulate the firewall in a safe and controlled manner and keeps an audit log.",
			"Able to control a single (localhost) system as well as clusters.",
		}),
	}),
	Example: utils.Examples ([] string {
		"// This is an example",
		"jrctl firewall allow -a 1.1.1.1 -p 80 -p 443",
		"",
		"// This is an example",
		"jrctl firewall allow -t db -a 1.1.1.1 -p 3306",
		"jrctl firewall allow -t admin -a 1.1.1.1 -p 22 -c 'Office'",
	}),
	Run: func ( cmd * cobra.Command, args [] string ) {
		address, _ := cmd.Flags ().GetString ("address")
		ports, _ := cmd.Flags ().GetIntSlice ("port")
		comment, _ := cmd.Flags ().GetString ("comment")
		selector, _ := cmd.Flags ().GetString ("type")
		rows := [] [] string { [] string { "Server", "Status", "Response" } }
		runner := func ( index, total int, context server.Context ) {
			data := firewall.AllowRequest {
				Address: address,
				Ports: ports,
				Blame: utils.SafeString ( utils.GetUser () ),
				Comment: utils.SafeString ( comment ),
			}
			response := firewall.Add ( context, data )
			row := [] string {
				strings.TrimSuffix ( context.Endpoint, ":27482" ),
				response.Status,
				response.Messages [ 0 ],
			}
			rows = append ( rows, row )
		}
		server.FilterForEach ( [] string { selector }, runner )
		utils.TablePrint ( fmt.Sprintf ( "No configured servers found with type %q.", selector ), rows, 1 )
	},
}

func init () {
	firewallCmd.AddCommand ( firewallAllowCmd )
	firewallAllowCmd.Flags ().SortFlags = true
	firewallAllowCmd.Flags ().StringP ( "type", "t", "localhost", "specify deamon type selector, useful for cluster deployments" )
	firewallAllowCmd.Flags ().StringP ( "address", "a", "", "ip address" )
	firewallAllowCmd.Flags ().IntSliceP ( "port", "p", [] int {}, "port to allow, can be specified multiple times" )
	firewallAllowCmd.Flags ().StringP ( "comment", "c", "none", "add a comment to the firewall entry (optional)" )
	firewallAllowCmd.MarkFlagRequired ("address")
	firewallAllowCmd.MarkFlagRequired ("port")
}
