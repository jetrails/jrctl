package internal

import (
	"errors"
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/server"
)

var serverRestartCmd = &cobra.Command {
	Use: "restart SERVICE...",
	Args: cobra.MinimumNArgs ( 1 ),
	Short: "Restart specified service(s) running on configured server(s)",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"Restart specified service(s) running on configured server(s).",
			"In order to successfully restart a service, the server first validates the respected service's config file.",
			"Once deemed valid, the service is restarted.",
			"For a list of available running services, run 'jrctl server list'.",
			"Services can be repeated and execution will happen in the order that is given.",
			"Specifing a server type will only display results for servers of that type.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl service restart nginx",
		"jrctl service restart nginx varnish",
		"jrctl service restart nginx varnish php-fpm",
		"jrctl service restart nginx varnish php-fpm-7.2 nginx",
	}),
	RunE: func ( cmd * cobra.Command, args [] string ) error {
		validServices := server.CollectServices ()
		for _, arg := range args {
			if !utils.Includes ( arg, validServices ) {
				return errors.New ( fmt.Sprintf (
					"%q is not found, available services include: %v",
					arg, "\"" + strings.Join ( validServices, "\", \"" ) + "\"",
				))
			}
		}
		cmd.Run ( cmd, args )
		return nil
	},
	Run: func ( cmd * cobra.Command, args [] string ) {
		rows := [] [] string { [] string { "Server", "Service", "Response" } }
		selector, _ := cmd.Flags ().GetString ("type")
		for _, arg := range args {
			runner := func ( index, total int, context server.Context ) {
				data := server.RestartRequest { Service: arg, Version: "" }
				if strings.HasPrefix ( arg, "php-fpm" ) {
					data.Service = "php-fpm"
					data.Version = strings.Join ( strings.Split ( arg, "-" ) [2:], "-" )
				}
				response := server.Restart ( context, data )
				row := [] string {
					strings.TrimSuffix ( context.Endpoint, ":27482" ),
					arg,
					response.Messages [ 0 ],
				}
				rows = append ( rows, row )
			}
			server.FilterWithServiceForEach ( selector, arg, runner )
		}
		if len ( rows ) > 1 {
			fmt.Printf ( "\nExecuted only on %q server(s):\n", selector )
		}
		utils.TablePrint ( fmt.Sprintf ( "Specified service(s) not running on %q server(s).", selector ), rows, 1 )
	},
}

func init () {
	serverCmd.AddCommand ( serverRestartCmd )
	serverRestartCmd.Flags ().SortFlags = true
	serverRestartCmd.Flags ().StringP ( "type", "t", "localhost", "specify server type, useful for cluster" )
}
