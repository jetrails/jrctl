package internal

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/whitelist"
)

var whitelistListCmd = &cobra.Command {
	Use:   "list",
	Short: "List whitelist whitelist entries",
	Example: utils.Examples ([] string {
		"jrctl whitelist list",
	}),
	Run: func ( cmd * cobra.Command, args [] string ) {
		context := whitelist.DaemonContext {
			Endpoint: viper.GetString ("daemon_endpoint"),
			Auth: viper.GetString ("daemon_token"),
			Debug: viper.GetBool ("debug"),
		}
		response := whitelist.List ( context )
		utils.PrintErrors ( response.Code, response.Status )
		utils.PrintMessages ( response.Messages )
		if len ( response.Payload ) > 0 {
			fmt.Println ()
			for i, entry := range response.Payload {
				fmt.Printf (
					"%-3s %s:%d\n",
					fmt.Sprintf ( "%d.", i + 1 ),
					entry.Address,
					entry.Port,
				)
			}
		}
	},
}

func init () {
	whitelistCmd.AddCommand ( whitelistListCmd )
	whitelistListCmd.Flags ().StringP ( "endpoint", "e", "localhost:27482", "specify endpoint hostname" )
	whitelistListCmd.Flags ().StringP ( "token", "t", "", "specify auth token" )
	viper.BindPFlag ( "daemon_endpoint", whitelistListCmd.Flags ().Lookup ("endpoint") )
	viper.BindPFlag ( "daemon_token", whitelistListCmd.Flags ().Lookup ("token") )
}
