package internal

import (
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/database"
	"github.com/spf13/cobra"
)

var databaseCreateCmd = &cobra.Command{
	Use:     "create DB_NAME",
	Short:   "Create a database",
	Args:    cobra.ExactArgs(1),
	Example: text.Examples([]string{}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("tag")
		dbName := args[0]

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

		request := database.CreateRequest{Name: dbName}
		response := database.Create(contexts[0], request)
		generic := response.GetGeneric()

		output.PrintResponse(generic)
		output.PrintDivider()
		output.ExitCodeFromResponse(generic)
	},
}

func init() {
	// OnlyRunOnNonAWS(databaseCreateCmd)
	databaseCmd.AddCommand(databaseCreateCmd)
	databaseCreateCmd.Flags().SortFlags = true
	databaseCreateCmd.Flags().BoolP("quiet", "q", false, "display no output")
	databaseCreateCmd.Flags().StringArrayP("tag", "t", []string{"default"}, "filter nodes using tags")
}
