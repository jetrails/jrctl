package internal

import (
	"fmt"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/database"
	"github.com/spf13/cobra"
)

var databaseUserDeleteCmd = &cobra.Command{
	Use:     "delete USER@HOST",
	Short:   "Delete database user",
	Example: text.Examples([]string{}),
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("tag")
		name, from := SplitUserAndHost(args[0])

		output := NewOutput(quiet, tags)
		contexts := config.GetContexts(tags)

		output.PrintTags()

		if len(contexts) < 1 {
			output.ExitWithMessage(1, ErrMsgNoServers+"\n")
		}

		if len(contexts) > 1 {
			output.ExitWithMessage(5, ErrMsgRequiresOneServer+"\n")
		}

		request := database.UserDeleteRequest{Name: name, From: from}
		response := database.UserDelete(contexts[0], request)
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
	// OnlyRunOnNonAWS(databaseUserDeleteCmd)
	databaseUserCmd.AddCommand(databaseUserDeleteCmd)
	databaseUserDeleteCmd.Flags().SortFlags = true
	databaseUserDeleteCmd.Flags().BoolP("quiet", "q", false, "only display confirmation code")
	databaseUserDeleteCmd.Flags().StringArrayP("tag", "t", []string{"default"}, "filter nodes using tags")
}
