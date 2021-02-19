package internal

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/firewall"
)

var firewallListCmd = &cobra.Command {
	Use:   "list",
	Short: "List firewall firewall entries",
	Example: utils.Examples ([] string {
		"jrctl firewall list",
	}),
	Run: func ( cmd * cobra.Command, args [] string ) {
		context := firewall.DaemonContext {
			Endpoint: viper.GetString ("daemon_endpoint"),
			Auth: viper.GetString ("daemon_token"),
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
	viper.BindPFlag ( "daemon_endpoint", firewallListCmd.Flags ().Lookup ("endpoint") )
	viper.BindPFlag ( "daemon_token", firewallListCmd.Flags ().Lookup ("token") )
}
