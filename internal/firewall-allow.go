package internal

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			"Ask the daemon to create an allow entry in the system's firewall.",
		}),
		utils.Paragraph ( [] string {
			"The following environmental variables can be passed in place of the 'endpoint' and 'token' flags: JR_DAEMON_ENDPOINT, JR_DAEMON_TOKEN.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl firewall allow -a 1.1.1.1 -p 80 -p 443",
		"jrctl firewall allow -a 1.1.1.1 -p 80,443 -b me",
		"jrctl firewall allow -a 1.1.1.1 -p 80,443 -b me -c 'Office'",
	}),
	PreRun: func ( cmd * cobra.Command, args [] string ) {
		viper.BindPFlag ( "daemon_endpoint", cmd.Flags ().Lookup ("endpoint") )
		viper.BindPFlag ( "daemon_token", cmd.Flags ().Lookup ("token") )
		if viper.GetString ("daemon_token") == "" {
			viper.Set ( "daemon_token", utils.LoadDaemonAuth () )
		}
	},
	Run: func ( cmd * cobra.Command, args [] string ) {
		address, _ := cmd.Flags ().GetString ("address")
		ports, _ := cmd.Flags ().GetIntSlice ("port")
		blame, _ := cmd.Flags ().GetString ("blame")
		comment, _ := cmd.Flags ().GetString ("comment")
		context := daemon.Context {
			Endpoint: viper.GetString ("daemon_endpoint"),
			Token: viper.GetString ("daemon_token"),
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
}
