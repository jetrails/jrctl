package internal

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/firewall"
	"github.com/jetrails/jrctl/sdk/daemon"
)

var firewallAllowCmd = &cobra.Command {
	Use:   "allow",
	Short: "Add entry to firewall",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"Add entry to firewall.",
			"Ask the daemon(s) to create an allow entry to their host system's firewall.",
			"The tag flag is useful for cluster deployments and allows executing command on daemons that are tagged a certain way.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl firewall allow -a 1.1.1.1 -p 80 -p 443",
		"jrctl firewall allow -t admin -a 1.1.1.1 -p 22 -b me",
		"jrctl firewall allow -t mysql -a 1.1.1.1 -p 3306 -b me -c 'Office'",
	}),
	Run: func ( cmd * cobra.Command, args [] string ) {
		address, _ := cmd.Flags ().GetString ("address")
		ports, _ := cmd.Flags ().GetIntSlice ("port")
		blame, _ := cmd.Flags ().GetString ("blame")
		comment, _ := cmd.Flags ().GetString ("comment")
		tag, _ := cmd.Flags ().GetString ("tag")
		rows := [] [] string { [] string { "Daemon", "Status", "Response" } }
		runner := func ( index, total int, context daemon.Context ) {
			data := firewall.AllowRequest {
				Address: address,
				Ports: ports,
				Blame: utils.SafeString ( blame ),
				Comment: utils.SafeString ( comment ),
			}
			response := firewall.Add ( context, data )
			row := [] string {
				context.Endpoint,
				response.Status,
				response.Messages [ 0 ],
			}
			rows = append ( rows, row )
		}
		daemon.FilterForEach ( [] string { tag }, runner )
		utils.TablePrint ( fmt.Sprintf ( "No configured daemons found with tag %q.", tag ), rows, 1 )
	},
}

func init () {
	firewallCmd.AddCommand ( firewallAllowCmd )
	firewallAllowCmd.Flags ().SortFlags = false
	firewallAllowCmd.Flags ().StringP ( "tag", "t", "localhost", "specify deamon tag selector, useful for cluster deployments" )
	firewallAllowCmd.Flags ().StringP ( "address", "a", "", "ip address to firewall" )
	firewallAllowCmd.Flags ().IntSliceP ( "port", "p", [] int {}, "port(s) to firewall" )
	firewallAllowCmd.Flags ().StringP ( "comment", "c", "none", "add optional comment" )
	firewallAllowCmd.Flags ().StringP ( "blame", "b", utils.GetUser (), "specify blame entry" )
	firewallAllowCmd.MarkFlagRequired ("address")
	firewallAllowCmd.MarkFlagRequired ("port")
}
