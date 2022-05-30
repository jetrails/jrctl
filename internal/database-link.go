package internal

import (
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/database"
	"github.com/spf13/cobra"
)

var databaseLinkCmd = &cobra.Command{
	Use:     "link USER@FROM DB_NAME",
	Short:   "Add database user to specific database",
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
		output.PrintDivider()

		if len(contexts) < 1 {
			output.ExitWithMessage(1, ErrMsgNoServers+"\n")
		}

		if len(contexts) > 1 {
			output.ExitWithMessage(5, ErrMsgRequiresOneServer+"\n")
		}

		request := database.LinkRequest{
			Database: dbName,
			Name:     user,
			From:     from,
		}
		response := database.Link(contexts[0], request)
		generic := response.GetGeneric()

		output.PrintResponse(generic)
		output.PrintDivider()
		output.ExitCodeFromResponse(generic)
	},
}

func init() {
	databaseCmd.AddCommand(databaseLinkCmd)
	databaseLinkCmd.Flags().SortFlags = true
	databaseLinkCmd.Flags().BoolP("quiet", "q", false, "display no output")
	databaseLinkCmd.Flags().StringArrayP("type", "t", []string{"localhost"}, "filter servers using type selectors")
}
