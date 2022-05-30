package internal

import (
	"strings"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/database"
	"github.com/spf13/cobra"
)

var databaseUserListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Display database users in deployment",
	Example: text.Examples([]string{}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("type")

		output := NewOutput(quiet, tags)
		output.DisplayServers = false
		output.FailOnNoServers = true
		output.FailOnNoResults = true
		output.ExitCodeNoServers = 1
		output.ExitCodeNoResults = 2

		tbl := output.CreateTable(Columns{
			"Server",
			"User",
			"Databases(s)",
		})

		for _, context := range config.GetContexts(tags) {
			response := database.ListUsers(context)
			output.AddServer(
				context,
				response.GetGeneric(),
				response.Status,
			)
			for _, entry := range response.Payload {
				tbl.AddQuietEntry(entry.Name + "@" + entry.From)
				if len(entry.Databases) == 0 {
					tbl.AddRow(Columns{
						strings.TrimSuffix(context.Endpoint, ":27482"),
						entry.Name + "@" + entry.From,
						"-",
					})
				} else {
					dbsWithHost := []string{}
					for _, db := range entry.Databases {
						dbsWithHost = append(dbsWithHost, db.Name)
					}
					tbl.AddRow(Columns{
						strings.TrimSuffix(context.Endpoint, ":27482"),
						entry.Name + "@" + entry.From,
						strings.Join(dbsWithHost, ", "),
					})
				}
			}
		}

		output.Print()
	},
}

func init() {
	databaseUserCmd.AddCommand(databaseUserListCmd)
	databaseUserListCmd.Flags().SortFlags = true
	databaseUserListCmd.Flags().BoolP("quiet", "q", false, "only display database user names")
	databaseUserListCmd.Flags().StringArrayP("type", "t", []string{"localhost"}, "specify server type, useful for cluster")
}
