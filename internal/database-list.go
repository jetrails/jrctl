package internal

import (
	"fmt"
	"strings"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/database"
	"github.com/spf13/cobra"
)

var databaseListCmd = &cobra.Command{
	Use:   "list",
	Short: "Display databases in deployment",
	Example: text.Examples([]string{
		"jrctl database list",
		"jrctl database list -q",
		"jrctl database list -t db",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("type")

		output := NewOutput(quiet, tags)
		output.DisplayServers = true
		output.FailOnNoServers = true
		output.FailOnNoResults = true
		output.ExitCodeNoServers = 1
		output.ExitCodeNoResults = 2

		tbl := output.CreateTable(Columns{
			"Server",
			"Database",
			"User(s)",
		})

		for _, context := range config.GetContexts(tags) {
			response := database.ListDatabases(context)
			output.AddServer(
				context,
				response.GetGeneric(),
				response.GetFirstMessage(),
			)
			for _, entry := range response.Payload {
				tbl.AddQuietEntry(entry.Name)
				if len(entry.Users) == 0 {
					tbl.AddRow(Columns{
						strings.TrimSuffix(context.Endpoint, ":27482"),
						entry.Name,
					})
				} else {
					formattedUsers := []string{}
					for _, user := range entry.Users {
						formattedUsers = append(formattedUsers, fmt.Sprintf("%s@%s", user.Name, user.From))
					}
					tbl.AddRow(Columns{
						strings.TrimSuffix(context.Endpoint, ":27482"),
						entry.Name,
						strings.Join(formattedUsers, ", "),
					})
				}
			}
		}

		output.Print()
	},
}

func init() {
	databaseCmd.AddCommand(databaseListCmd)
	databaseListCmd.Flags().SortFlags = true
	databaseListCmd.Flags().BoolP("quiet", "q", false, "only display database names")
	databaseListCmd.Flags().StringArrayP("type", "t", []string{"localhost"}, "filter servers using type selectors")
}
