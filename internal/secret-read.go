package internal

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/jetrails/jrctl/pkg/env"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/secret"
	"github.com/spf13/cobra"
)

var secretReadCmd = &cobra.Command{
	Use:   "read IDENTIFIER",
	Short: "Display contents of secret",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Display contents of secret.",
			"Passing the secret identifier will allow us to retrieve the contents of the secret and print it to STDOUT.",
			"Optionally, you can copy the contents to your clipboard by passing the --clipboard flag!",
			"If the secret's URL is passed, the identifier is extracted automatically.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl secret read 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq",
		"jrctl secret read 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq -c",
		"jrctl secret read 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq -c -p secretpass",
	}),
	Args: func(cmd *cobra.Command, args []string) error {
		check := cobra.MinimumNArgs(1)
		if error := check(cmd, args); error != nil {
			return error
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		identifier := args[0]
		identifier = strings.TrimPrefix(identifier, fmt.Sprintf("https://%s/secret/", env.GetString("secret_endpoint", "secret.jetrails.cloud")))
		identifier = strings.Trim(identifier, "/")
		copy, _ := cmd.Flags().GetBool("clipboard")
		password, _ := cmd.Flags().GetString("password")
		context := secret.PublicApiContext{
			Endpoint: env.GetString("public_api_endpoint", "api-public.jetrails.cloud"),
			Debug:    env.GetBool("debug", false),
			Insecure: env.GetBool("insecure", false),
		}
		request := secret.SecretReadRequest{
			Identifier: identifier,
			Password:   password,
		}
		response, error := secret.SecretRead(context, request)
		if error != nil && error.Code != 200 {
			fmt.Printf("\n%s\n\n", error.Message)
			return
		} else {
			fmt.Printf("\n%s\n\n", response.Data)
		}
		if copy {
			clipboard.WriteAll(strings.TrimSpace(response.Data))
		}
	},
}

func init() {
	secretCmd.AddCommand(secretReadCmd)
	secretReadCmd.Flags().SortFlags = true
	secretReadCmd.Flags().StringP("password", "p", "", "password to access secret")
	secretReadCmd.Flags().BoolP("clipboard", "c", false, "copy contents to clipboard")
}
