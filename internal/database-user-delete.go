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

var databaseUserDeleteCmd = &cobra.Command{
	Use:   "delete USER@HOST",
	Short: "",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			".",
		}),
	}),
	Example: text.Examples([]string{}),
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		selectors, _ := cmd.Flags().GetStringSlice("type")
		name, from := splitUserAndHost(args[0])
		runner := func(index, total int, context server.Context) {
			if total > 1 {
				fmt.Println("\nError: multiple servers match, must narrow down to one server using type selectors\n")
				os.Exit(1)
			}
			hash := context.Hash()
			request := database.UserDeleteRequest{
				Name: name,
				From: from,
			}
			response := database.UserDelete(context, request)
			if response.Code == 200 {
				if !quiet {
					fmt.Println("")
					fmt.Println("WARNING: This is a destructive command that cannot be undone. If you would")
					fmt.Println("like to continue, you will need to send a confirmation to the server to")
					fmt.Printf("execute this destructive command (%s).", strings.Join(response.Messages, ", "))
					fmt.Println("\n")
					fmt.Printf("Run the following command: jrctl confirm %s-%s\n", hash, response.Payload.Identifier)
					fmt.Println("")
				} else {
					fmt.Printf("%s-%s\n", hash, response.Payload.Identifier)
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
	databaseUserCmd.AddCommand(databaseUserDeleteCmd)
	databaseUserDeleteCmd.Flags().SortFlags = true
	databaseUserDeleteCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
	databaseUserDeleteCmd.Flags().StringSliceP("type", "t", []string{"localhost"}, "specify server type, useful for cluster")
}
