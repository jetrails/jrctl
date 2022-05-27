package internal

import (
	"fmt"
	"strings"

	"github.com/jetrails/jrctl/pkg/env"
	"github.com/jetrails/jrctl/pkg/input"
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/api"
	"github.com/jetrails/jrctl/sdk/secret"
	"github.com/spf13/cobra"
)

var secretDeleteCmd = &cobra.Command{
	Use:   "delete IDENTIFIER",
	Short: "Delete secret without viewing contents",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Delete secret without viewing contents.",
			"Passing the secret identifier will make a request to destroy the secret without displaying the secret's contents.",
			"If the secret's URL is passed, the identifier is extracted automatically.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl secret delete 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq",
		"echo 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq | jrctl secret delete",
	}),
	Args: func(cmd *cobra.Command, args []string) error {
		if input.HasDataInPipe() && len(args) == 0 {
			return nil
		}
		return cobra.ExactArgs(1)(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")

		tbl := NewTable(Columns{})
		tbl.Quiet = quiet

		identifier := ""
		if len(args) == 0 {
			identifier = input.GetPipeData()
		} else {
			identifier = args[0]
		}

		prefix := fmt.Sprintf("https://%s/secret/", env.GetString("secret_endpoint", "secret.jetrails.com"))
		identifier = strings.TrimPrefix(identifier, prefix)
		identifier = strings.Trim(identifier, "/")

		context := secret.PublicApiContext{
			Endpoint: env.GetString("public_api_endpoint", "api-public.jetrails.com"),
			Debug:    env.GetBool("debug", false),
			Insecure: env.GetBool("insecure", false),
		}
		request := secret.SecretDeleteRequest{
			Identifier: identifier,
		}
		response, err := secret.SecretDelete(context, request)

		if err != nil && err.Code != 200 {
			generic := api.GenericResponse{
				Code:     err.Code,
				Status:   err.Type,
				Messages: []string{err.Message},
			}
			tbl.PrintDivider()
			tbl.PrintResponse(&generic)
			tbl.PrintDivider()
			tbl.ExitCodeFromResponse(&generic)
		}

		tbl.ExitWithMessage(0, "\nSuccessfully deleted secret: '%s'\n", response.Identifier)
	},
}

func init() {
	secretCmd.AddCommand(secretDeleteCmd)
	secretDeleteCmd.Flags().SortFlags = true
	secretDeleteCmd.Flags().BoolP("quiet", "q", false, "display no output")
}
