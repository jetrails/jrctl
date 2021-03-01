package internal

import (
	"fmt"
	"github.com/spf13/viper"
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
		}),
		utils.Paragraph ( [] string {
			"The following environmental variables can be passed in place of the 'endpoint' and 'token' flags: JR_DAEMON_ENDPOINT, JR_DAEMON_TOKEN.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl firewall list",
	}),
	PreRun: func ( cmd * cobra.Command, args [] string ) {
		viper.BindPFlag ( "daemon_endpoint", cmd.Flags ().Lookup ("endpoint") )
		viper.BindPFlag ( "daemon_token", cmd.Flags ().Lookup ("token") )
		if viper.GetString ("daemon_token") == "" {
			viper.Set ( "daemon_token", utils.LoadDaemonAuth () )
		}
	},
	Run: func ( cmd * cobra.Command, args [] string ) {
		context := daemon.Context {
			Endpoint: viper.GetString ("daemon_endpoint"),
			Token: viper.GetString ("daemon_token"),
			Debug: viper.GetBool ("debug"),
		}
		response := firewall.List ( context )
		utils.PrintErrors ( response.Code, response.Status )
		utils.PrintMessages ( response.Messages )
		if len ( response.Payload ) > 0 {
			fmt.Println ()
			for i, entry := range response.Payload {
				fmt.Printf (
					"%-3s %s -> %v\n",
					fmt.Sprintf ( "%d.", i + 1 ),
					entry.Address,
					entry.Port,
				)
			}
		}
	},
}

func init () {
	firewallCmd.AddCommand ( firewallListCmd )
	firewallListCmd.Flags ().StringP ( "endpoint", "e", "localhost:27482", "specify endpoint hostname" )
	firewallListCmd.Flags ().StringP ( "token", "t", "", "specify auth token" )
}
