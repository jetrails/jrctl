package internal

import (
	"fmt"
	"strings"

	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/database"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

var databaseListCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			".",
		}),
	}),
	Example: text.Examples([]string{}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		selector, _ := cmd.Flags().GetString("type")
		rows := [][]string{{"Server", "Database", "User(s)"}}
		runner := func(index, total int, context server.Context) {
			response := database.ListDatabases(context)
			if response.Code == 200 {
				for _, entry := range response.Payload {
					if len(entry.Users) == 0 {
						rows = append(rows, []string{
							strings.TrimSuffix(context.Endpoint, ":27482"),
							entry.Name,
							"-",
						})
					} else {
						usersWithHost := []string{}
						for _, user := range entry.Users {
							userWithHost := fmt.Sprintf("%s@%s", user.Name, strings.ReplaceAll(user.From, "%", "anywhere"))
							usersWithHost = append(usersWithHost, userWithHost)
						}
						rows = append(rows, []string{
							strings.TrimSuffix(context.Endpoint, ":27482"),
							entry.Name,
							strings.Join(usersWithHost, ", "),
						})
					}
				}
			} else {
				fmt.Printf("\n%s [%d]\n\n", response.Status, response.Code)
			}
		}
		server.FilterForEach([]string{selector}, runner)
		if quiet {
			for _, row := range rows[1:] {
				fmt.Println (row[1])
			}
		} else {
			text.TablePrint("No entries found", rows, 1)
		}
	},
}

func init() {
	databaseCmd.AddCommand(databaseListCmd)
	databaseListCmd.Flags().SortFlags = true
	databaseListCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
	databaseListCmd.Flags().StringP("type", "t", "localhost", "specify server type, useful for cluster")
}
