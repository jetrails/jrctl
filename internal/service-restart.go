package internal

import (
	"errors"
	"regexp"
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/service"
	"github.com/jetrails/jrctl/sdk/daemon"
)

var serviceRestartCmd = &cobra.Command {
	Use: "restart SERVICE...",
	Args: cobra.MinimumNArgs ( 1 ),
	Short: "Restart apache, nginx, mysql, or varnish service",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"Restart apache, nginx, mysql, varnish, or php-fpm-* service.",
			"Valid entries for php-fpm services would be prefixed with 'php-fpm-' and followed by a version number.",
			"Ask the daemon(s) to restart a given service.",
			"In order to successfully restart it, the daemon first validates the respected service's configuration.",
			"Services can be repeated and execution will happen in the order that is given.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl service restart nginx",
		"jrctl service restart nginx varnish",
		"jrctl service restart nginx varnish php-fpm-7.2",
		"jrctl service restart nginx varnish php-fpm-7.2 nginx",
	}),
	RunE: func ( cmd * cobra.Command, args [] string ) error {
		pattern := regexp.MustCompile (`^(?:apache|nginx|mysql|varnish|php-fpm-\d+(?:\.\d+)*)$`)
		for _, arg := range args {
			if !pattern.MatchString ( arg ) {
				valid := [] string { "apache", "nginx", "mysql", "varnish", "php-fpm-*" }
				return errors.New ( fmt.Sprintf (
					"invalid service %q\nvalid services include: %v\nwhere \"*\" is replaced with a valid version string",
					arg, valid,
				))
			}
		}
		cmd.Run ( cmd, args )
		return nil
	},
	Run: func ( cmd * cobra.Command, args [] string ) {
		rows := [] [] string { [] string { "Daemon", "Status", "Service", "Response" } }
		tag, _ := cmd.Flags ().GetString ("tag")
		for _, arg := range args {
			filter := [] string { tag, arg }
			runner := func ( index, total int, context daemon.Context ) {
				data := service.RestartRequest { Service: arg, Version: "" }
				if strings.HasPrefix ( arg, "php-fpm-" ) {
					data.Service = "php-fpm"
					data.Version = strings.Join ( strings.Split ( arg, "-" ) [2:], "-" )
				}
				response := service.Restart ( context, data )
				row := [] string {
					strings.TrimSuffix ( context.Endpoint, ":27482" ),
					response.Status,
					arg,
					response.Messages [ 0 ],
				}
				rows = append ( rows, row )
			}
			daemon.FilterForEach ( filter, runner )
		}
		utils.TablePrint ( fmt.Sprintf ( "Specified services not running on server type %q.", tag ), rows, 1 )
	},
}

func init () {
	serviceCmd.AddCommand ( serviceRestartCmd )
	serviceRestartCmd.Flags ().SortFlags = true
	serviceRestartCmd.Flags ().StringP ( "tag", "t", "localhost", "specify deamon tag selector, useful for cluster deployments" )
}
