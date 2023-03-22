package internal

import (
	"strings"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/service"
	"github.com/spf13/cobra"
)

var serviceEnableCmd = &cobra.Command{
	Use:   "enable SERVICE...",
	Args:  cobra.MinimumNArgs(1),
	Short: "Enable specified services in deployment",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Enable specified services in deployment.",
			"Services can be repeated and execution will happen in the order that is given.",
			"Specifing a tag will display nodes that have that tag.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl service enable nginx",
		"jrctl service enable nginx varnish",
		"jrctl service enable nginx varnish php-fpm",
		"jrctl service enable nginx varnish php-fpm-7.2 nginx",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("tag")

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
					request := service.EnableRequest{Service: arg}
					response := service.Enable(context, request)
					output.AddUniqueServer(
						context,
						listResponse.GetGeneric(),
						listResponse.GetFirstMessage(),
					)
					tbl.AddRow(Columns{
						strings.TrimSuffix(context.Endpoint, ":27482"),
						arg,
						response.GetFirstMessage(),
					})
				}
			}
		}

		output.AddTable(tbl)
		output.Print()
	},
}

func init() {
	serviceCmd.AddCommand(serviceEnableCmd)
	serviceEnableCmd.Flags().SortFlags = true
	serviceEnableCmd.Flags().BoolP("quiet", "q", false, "display no output")
	serviceEnableCmd.Flags().StringArrayP("tag", "t", []string{"default"}, "filter nodes using tags")
}
