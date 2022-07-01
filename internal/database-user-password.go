package internal

import (
	"fmt"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/database"
	"github.com/spf13/cobra"
)

var databaseUserPasswordCmd = &cobra.Command{
	Use:   "password USER@HOST",
	Short: "Roll database user's password",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			".",
		}),
	}),
	Example: text.Examples([]string{}),
	Args:    cobra.ExactArgs(1),
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

		if len(contexts) < 1 {
			output.ExitWithMessage(1, ErrMsgNoServers+"\n")
		}

		if len(contexts) > 1 {
			output.ExitWithMessage(5, ErrMsgRequiresOneServer+"\n")
		}

		request := database.UserPasswordRequest{Name: name, From: from, Plugin: plugin}
		response := database.UserPassword(contexts[0], request)
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
	OnlyRunOnNonAWS(databaseUserPasswordCmd)
	databaseUserCmd.AddCommand(databaseUserPasswordCmd)
	databaseUserPasswordCmd.Flags().SortFlags = true
	databaseUserPasswordCmd.Flags().BoolP("quiet", "q", false, "only display confirmation code")
	databaseUserPasswordCmd.Flags().StringArrayP("type", "t", []string{"localhost"}, "filter servers using type selectors")
	databaseUserPasswordCmd.Flags().StringP("plugin", "p", "mysql_native_password", "specify auth plugin")
}
