package internal

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/whitelist"
)

var whitelistAddCmd = &cobra.Command {
	Use:   "add",
	Short: "Add entry to whitelist",
	Example: utils.Examples ([] string {
		"jrctl whitelist add -a 1.1.1.1 -p 80 -p 443",
		"jrctl whitelist add -a 1.1.1.1 -p 80,443 -b me",
		"jrctl whitelist add -a 1.1.1.1 -p 80,443 -b me -c 'Office'",
	}),
	Run: func ( cmd * cobra.Command, args [] string ) {
		address, _ := cmd.Flags ().GetString ("address")
		ports, _ := cmd.Flags ().GetIntSlice ("port")
		blame, _ := cmd.Flags ().GetString ("blame")
		comment, _ := cmd.Flags ().GetString ("comment")
		context := whitelist.DaemonContext {
			Endpoint: viper.GetString ("daemon_endpoint"),
			Auth: viper.GetString ("daemon_token"),
			Debug: viper.GetBool ("debug"),
		}
		data := whitelist.AddRequest {
			Address: address,
			Ports: ports,
			Blame: utils.SafeString ( blame ),
			Comment: utils.SafeString ( comment ),
		}
		response := whitelist.Add ( context, data )
		utils.PrintErrors ( response.Code, response.Status )
		utils.PrintMessages ( response.Messages )
	},
}

func init () {
	whitelistCmd.AddCommand ( whitelistAddCmd )
	whitelistAddCmd.Flags ().StringP ( "endpoint", "e", "localhost:27482", "specify endpoint hostname" )
	whitelistAddCmd.Flags ().StringP ( "token", "t", "", "specify auth token" )
	whitelistAddCmd.Flags ().StringP ( "address", "a", "", "IP address to whitelist" )
	whitelistAddCmd.Flags ().IntSliceP ( "port", "p", [] int {}, "port(s) to whitelist" )
	whitelistAddCmd.Flags ().StringP ( "comment", "c", "none", "add optional comment" )
	whitelistAddCmd.Flags ().StringP ( "blame", "b", utils.GetUser (), "specify blame entry" )
	whitelistAddCmd.MarkFlagRequired ("address")
	whitelistAddCmd.MarkFlagRequired ("port")
	viper.BindPFlag ( "daemon_endpoint", whitelistAddCmd.Flags ().Lookup ("endpoint") )
	viper.BindPFlag ( "daemon_token", whitelistAddCmd.Flags ().Lookup ("token") )
}
