package internal

import (
	"fmt"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/database"
	"github.com/spf13/cobra"
)

var databaseUserCreateCmd = &cobra.Command{
	Use:     "create USER@HOST",
	Short:   "Create database user",
	Args:    cobra.ExactArgs(1),
	Example: text.Examples([]string{}),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		validPlugins := []string{"mysql_native_password", "caching_sha2_password"}
		plugin, _ := cmd.Flags().GetString("plugin")
		for _, entry := range validPlugins {
			if entry == plugin {
				return nil
			}
		}
		return fmt.Errorf("invalid plugin, valid options include: %v", validPlugins)
	},
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("type")
		plugin, _ := cmd.Flags().GetString("plugin")
		name, from := SplitUserAndHost(args[0])

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

		request := database.UserCreateRequest{Name: name, From: from, Plugin: plugin}
		response := database.UserCreate(contexts[0], request)
		generic := response.GetGeneric()

		if response.IsOkay() {
			tbl := NewTable(Columns{"User", "Password"})
			tbl.Quiet = quiet
			tbl.AddRow(Columns{name + "@" + from, response.Payload})
			tbl.PrintTable()
			output.PrintDivider()
		} else {
			output.PrintResponse(generic)
			output.PrintDivider()
		}

		if quiet && response.IsOkay() {
			fmt.Println(response.Payload)
		}

		output.ExitCodeFromResponse(generic)
	},
}

func init() {
	OnlyRunOnNonAWS(databaseUserCreateCmd)
	databaseUserCmd.AddCommand(databaseUserCreateCmd)
	databaseUserCreateCmd.Flags().SortFlags = true
	databaseUserCreateCmd.Flags().BoolP("quiet", "q", false, "only display password")
	databaseUserCreateCmd.Flags().StringArrayP("type", "t", []string{"localhost"}, "filter servers using type selectors")
	databaseUserCreateCmd.Flags().StringP("plugin", "p", "mysql_native_password", "specify auth plugin")
}
