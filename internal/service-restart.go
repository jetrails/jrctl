package internal

import (
	"strings"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/service"
	"github.com/spf13/cobra"
)

var serviceRestartCmd = &cobra.Command{
	Use:   "restart SERVICE...",
	Args:  cobra.MinimumNArgs(1),
	Short: "Restart specified services in deployment",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Restart specified services in deployment.",
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
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("type")

		output := NewOutput(quiet, tags)
		output.DisplayServers = false
		output.FailOnNoServers = true
		output.FailOnNoResults = true
		output.ExitCodeNoServers = 1
		output.ExitCodeNoResults = 2

		tbl := NewTable(Columns{
			"Server",
			"Service",
			"Response",
		})

		for _, arg := range args {
			for _, context := range config.GetContexts(tags) {
				listResponse := service.ListServices(context)
				if _, hasService := listResponse.Payload[arg]; hasService {
					request := service.RestartRequest{Service: arg}
					response := service.Restart(context, request)
					output.AddUniqueServer(
						context,
						listResponse.GetGeneric(),
						listResponse.Messages[0],
					)
					tbl.AddRow(Columns{
						strings.TrimSuffix(context.Endpoint, ":27482"),
						arg,
						response.Messages[0],
					})
				}
			}
		}

		output.AddTable(tbl)
		output.Print()
	},
}

func init() {
	serviceCmd.AddCommand(serviceRestartCmd)
	serviceRestartCmd.Flags().SortFlags = true
	serviceRestartCmd.Flags().BoolP("quiet", "q", false, "display no output")
	serviceRestartCmd.Flags().StringArrayP("type", "t", []string{"localhost"}, "filter servers using type selectors")
}
