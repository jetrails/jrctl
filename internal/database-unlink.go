package internal

import (
	"fmt"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/database"
	"github.com/spf13/cobra"
)

var databaseUnlinkCmd = &cobra.Command{
	Use:     "unlink USER@FROM DB_NAME",
	Short:   "Remove database user to specific database",
	Args:    cobra.ExactArgs(2),
	Example: text.Examples([]string{}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("type")
		user, from := SplitUserAndHost(args[0])
		dbName := args[1]

		output := NewOutput(quiet, tags)
		contexts := config.GetContexts(tags)

		output.PrintTags()

		if len(contexts) < 1 {
			output.ExitWithMessage(1, ErrMsgNoServers+"\n")
		}

		if len(contexts) > 1 {
			output.ExitWithMessage(5, ErrMsgRequiresOneServer+"\n")
		}

		request := database.UnlinkRequest{
			Database: dbName,
			Name:     user,
			From:     from,
		}
		response := database.Unlink(contexts[0], request)
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
	databaseCmd.AddCommand(databaseUnlinkCmd)
	databaseUnlinkCmd.Flags().SortFlags = true
	databaseUnlinkCmd.Flags().BoolP("quiet", "q", false, "only display confirmation code")
	databaseUnlinkCmd.Flags().StringArrayP("type", "t", []string{"localhost"}, "filter servers using type selectors")
}
