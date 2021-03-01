package internal

import (
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/service"
	"github.com/jetrails/jrctl/sdk/daemon"
)

var serviceRestartCmd = &cobra.Command {
	Use: "restart SERVICE",
	Args: cobra.ExactValidArgs ( 1 ),
	ValidArgs: [] string { "apache", "nginx", "mysql", "varnish" },
	Short: "Restart apache, nginx, mysql, or varnish service",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"Restart apache, nginx, mysql, or varnish service.",
			"Ask the daemon to restart a given service.",
			"In order to successfully restart it, the daemon first validates the respected service's configuration.",
		}),
		utils.Paragraph ( [] string {
			"The following environmental variables can be passed in place of the 'endpoint' and 'token' flags: JR_DAEMON_ENDPOINT, JR_DAEMON_TOKEN.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl service restart apache",
		"jrctl service restart nginx",
		"jrctl service restart mysql",
		"jrctl service restart varnish",
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
		data := service.RestartRequest {
			Service: args [ 0 ],
		}
		response := service.Restart ( context, data )
		utils.PrintErrors ( response.Code, response.Status )
		utils.PrintMessages ( response.Messages )
	},
}

func init () {
	serviceCmd.AddCommand ( serviceRestartCmd )
	serviceRestartCmd.Flags ().StringP ( "endpoint", "e", "localhost:27482", "specify endpoint hostname" )
	serviceRestartCmd.Flags ().StringP ( "token", "t", "", "specify auth token" )
}
