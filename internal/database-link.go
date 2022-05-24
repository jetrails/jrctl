package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/database"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

var databaseLinkCmd = &cobra.Command{
	Use:   "link USER@FROM DB_NAME",
	Short: "",
	Args:  cobra.ExactArgs(2),
	Long: text.Combine([]string{
		text.Paragraph([]string{
			".",
		}),
	}),
	Example: text.Examples([]string{}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		selectors, _ := cmd.Flags().GetStringSlice("type")
		user, from := splitUserAndHost(args[0])
		dbName := args[1]
		runner := func(index, total int, context server.Context) {
			if total > 1 {
				fmt.Println("\nError: multiple servers match, must narrow down to one server using type selectors\n")
				os.Exit(1)
			}
			request := database.LinkRequest{
				Database: dbName,
				Name:     user,
				From:     from,
			}
			response := database.Link(context, request)
			if response.Code == 200 {
				if !quiet {
					fmt.Printf("\n%s\n\n", strings.Join(response.Messages, ", "))
				}
				os.Exit(0)
			} else {
				if !quiet {
					fmt.Printf("\n%s [%d]: %s\n\n", response.Status, response.Code, strings.Join(response.Messages, ", "))
				}
				os.Exit(1)
			}
		}
		if total := server.FilterForEach(selectors, runner); total != 1 {
			fmt.Println("\nError: no matching servers found, please change type selector\n")
			os.Exit(1)
		}
	},
}

func init() {
	databaseCmd.AddCommand(databaseLinkCmd)
	databaseLinkCmd.Flags().SortFlags = true
	databaseLinkCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
	databaseLinkCmd.Flags().StringSliceP("type", "t", []string{"localhost"}, "specify server type, useful for cluster")
}
