package internal

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/firewall"
)

var firewallAllowCmd = &cobra.Command {
	Use:   "allow",
	Short: "Add entry to firewall",
	Example: utils.Examples ([] string {
		"jrctl firewall allow -a 1.1.1.1 -p 80 -p 443",
		"jrctl firewall allow -a 1.1.1.1 -p 80,443 -b me",
		"jrctl firewall allow -a 1.1.1.1 -p 80,443 -b me -c 'Office'",
	}),
	Run: func ( cmd * cobra.Command, args [] string ) {
		address, _ := cmd.Flags ().GetString ("address")
		ports, _ := cmd.Flags ().GetIntSlice ("port")
		blame, _ := cmd.Flags ().GetString ("blame")
		comment, _ := cmd.Flags ().GetString ("comment")
		context := firewall.DaemonContext {
			Endpoint: viper.GetString ("daemon_endpoint"),
			Auth: viper.GetString ("daemon_token"),
			Debug: viper.GetBool ("debug"),
		}
		data := firewall.AllowRequest {
			Address: address,
			Ports: ports,
			Blame: utils.SafeString ( blame ),
			Comment: utils.SafeString ( comment ),
		}
		response := firewall.Add ( context, data )
		utils.PrintErrors ( response.Code, response.Status )
		utils.PrintMessages ( response.Messages )
	},
}

func init () {
	firewallCmd.AddCommand ( firewallAllowCmd )
	firewallAllowCmd.Flags ().StringP ( "endpoint", "e", "localhost:27482", "specify endpoint hostname" )
	firewallAllowCmd.Flags ().StringP ( "token", "t", "", "specify auth token" )
	firewallAllowCmd.Flags ().StringP ( "address", "a", "", "IP address to firewall" )
	firewallAllowCmd.Flags ().IntSliceP ( "port", "p", [] int {}, "port(s) to firewall" )
	firewallAllowCmd.Flags ().StringP ( "comment", "c", "none", "add optional comment" )
	firewallAllowCmd.Flags ().StringP ( "blame", "b", utils.GetUser (), "specify blame entry" )
	firewallAllowCmd.MarkFlagRequired ("address")
	firewallAllowCmd.MarkFlagRequired ("port")
	viper.BindPFlag ( "daemon_endpoint", firewallAllowCmd.Flags ().Lookup ("endpoint") )
	viper.BindPFlag ( "daemon_token", firewallAllowCmd.Flags ().Lookup ("token") )
}
