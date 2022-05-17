package internal

import (
	"fmt"
	"strings"

	"github.com/jetrails/jrctl/pkg/array"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

var serviceRestartCmd = &cobra.Command{
	Use:   "restart SERVICE...",
	Args:  cobra.MinimumNArgs(1),
	Short: "Restart specified service(s) running on configured server(s)",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Restart specified service(s) running on configured server(s).",
			"In order to successfully restart a service, the server first validates the respected service's config file.",
			"Once deemed valid, the service is restarted.",
			"Services can be repeated and execution will happen in the order that is given.",
			"Specifing a server type will only display results for servers of that type.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl service restart nginx",
		"jrctl service restart nginx varnish",
		"jrctl service restart nginx varnish php-fpm",
		"jrctl service restart nginx varnish php-fpm-7.2 nginx",
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		validServices := server.CollectServices()
		for _, arg := range args {
			if !array.ContainsString(validServices, arg) {
				return fmt.Errorf(
					"%q is not found, available services include: %v",
					arg, "\""+strings.Join(validServices, "\", \"")+"\"",
				)
			}
		}
		cmd.Run(cmd, args)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		rows := [][]string{{"Server", "Service", "Response"}}
		selectors, _ := cmd.Flags().GetStringSlice("type")
		for _, arg := range args {
			runner := func(index, total int, context server.Context) {
				data := server.RestartRequest{Service: arg}
				response := server.Restart(context, data)
				row := []string{
					strings.TrimSuffix(context.Endpoint, ":27482"),
					arg,
					response.Messages[0],
				}
				rows = append(rows, row)
			}
			server.FiltersWithServiceForEach(selectors, arg, runner)
		}
		if !quiet {
			if len(rows) > 1 {
				fmt.Printf("\nExecuted only on %q server(s):\n", text.QuotedList(selectors))
			}
			text.TablePrint(fmt.Sprintf("Specified service(s) not running on %q server(s).", text.QuotedList(selectors)), rows, 1)
		}
	},
}

func init() {
	serviceCmd.AddCommand(serviceRestartCmd)
	serviceRestartCmd.Flags().SortFlags = true
	serviceRestartCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
	serviceRestartCmd.Flags().StringSliceP("type", "t", []string{"localhost"}, "specify server type(s), useful for cluster")
}
