package internal

import (
	"fmt"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/database"
	"github.com/spf13/cobra"
)

var databaseDeleteCmd = &cobra.Command{
	Use:     "delete DB_NAME",
	Short:   "Delete a database",
	Args:    cobra.ExactArgs(1),
	Example: text.Examples([]string{}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("type")
		dbName := args[0]

		output := NewOutput(quiet, tags)
		contexts := config.GetContexts(tags)

		output.PrintTags()

		if len(contexts) < 1 {
			output.ExitWithMessage(1, ErrMsgNoServers+"\n")
		}

		if len(contexts) > 1 {
			output.ExitWithMessage(5, ErrMsgRequiresOneServer+"\n")
		}

		request := database.DeleteRequest{Name: dbName}
		response := database.Delete(contexts[0], request)
		generic := response.GetGeneric()

		if quiet && response.IsOkay() {
			fmt.Println(response.Payload.Identifier)
		} else if response.IsOkay() {
			PrintConfirmMessage(tags, response.Payload.Identifier, response.Metadata["ttl"])
		} else {
			output.PrintDivider()
			output.PrintResponse(generic)
			output.PrintDivider()
		}

		output.ExitCodeFromResponse(generic)
	},
}

func init() {
	databaseCmd.AddCommand(databaseDeleteCmd)
	databaseDeleteCmd.Flags().SortFlags = true
	databaseDeleteCmd.Flags().BoolP("quiet", "q", false, "only display confirmation code")
	databaseDeleteCmd.Flags().StringArrayP("type", "t", []string{"localhost"}, "filter servers using type selectors")
}
