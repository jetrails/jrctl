package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jetrails/jrctl/pkg/env"
	"github.com/jetrails/jrctl/pkg/text"
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
		if stat, _ := os.Stdin.Stat(); stat.Mode()&os.ModeCharDevice == 0 && len(args) == 0 {
			return nil
		}
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		identifier := ""
		if len(args) == 0 {
			if bytes, err := ioutil.ReadAll(os.Stdin); err == nil {
				identifier = strings.TrimSpace(string(bytes))
			}
		} else {
			identifier = args[0]
		}
		identifier = strings.TrimPrefix(identifier, fmt.Sprintf("https://%s/secret/", env.GetString("secret_endpoint", "secret.jetrails.com")))
		identifier = strings.Trim(identifier, "/")
		quiet, _ := cmd.Flags().GetBool("quiet")
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
			if !quiet {
				fmt.Printf("\n%s\n\n", err.Message)
			}
			os.Exit(1)
		}
		if !quiet {
			fmt.Printf("\nSuccessfully deleted secret: '%s'\n\n", response.Identifier)
		}
	},
}

func init() {
	secretCmd.AddCommand(secretDeleteCmd)
	secretDeleteCmd.Flags().SortFlags = true
	secretDeleteCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
}
