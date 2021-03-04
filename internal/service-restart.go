package internal

import (
	"fmt"
	"errors"
	"strings"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/service"
	"github.com/jetrails/jrctl/sdk/daemon"
)

var serviceRestartCmd = &cobra.Command {
	Use: "restart SERVICE...",
	Args: cobra.MinimumNArgs ( 1 ),
	ValidArgs: [] string { "apache", "nginx", "mysql", "varnish" },
	Short: "Restart apache, nginx, mysql, or varnish service",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"Restart apache, nginx, mysql, or varnish service.",
			"Ask the daemon(s) to restart a given service.",
			"In order to successfully restart it, the daemon first validates the respected service's configuration.",
			"Services can be repeated and execution will happen in the order that is given.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl service restart nginx",
		"jrctl service restart nginx varnish",
		"jrctl service restart nginx varnish nginx",
	}),
	RunE: func ( cmd * cobra.Command, args [] string ) error {
		if error := cobra.OnlyValidArgs ( cmd, args ); error != nil {
			return errors.New ( fmt.Sprintf (
				"%s\nvalid arguments include: %v",
				error.Error (),
				strings.Join ( cmd.ValidArgs, ", " ),
			))
		}
		cmd.Run ( cmd, args )
		return nil
	},
	Run: func ( cmd * cobra.Command, args [] string ) {
		rows := [] [] string { [] string { "Daemon", "Status", "Service", "Response" } }
		for _, arg := range args {
			filter := [] string { arg }
			runner := func ( index, total int, context daemon.Context ) {
				data := service.RestartRequest { Service: arg }
				response := service.Restart ( context, data )
				row := [] string {
					context.Endpoint,
					response.Status,
					arg,
					response.Messages [ 0 ],
				}
				rows = append ( rows, row )
			}
			daemon.FilterForEach ( filter, runner )
		}
		utils.TablePrint ("No configured daemons running passed services.", rows, 1 )
	},
}

func init () {
	serviceCmd.AddCommand ( serviceRestartCmd )
	serviceRestartCmd.Flags ().SortFlags = false
}
