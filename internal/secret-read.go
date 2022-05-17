package internal

import (
	"fmt"
	"io/ioutil"
	"os"
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
		"echo 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq | jrctl secret read",
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
		copy, _ := cmd.Flags().GetBool("clipboard")
		password, _ := cmd.Flags().GetString("password")
		context := secret.PublicApiContext{
			Endpoint: env.GetString("public_api_endpoint", "api-public.jetrails.com"),
			Debug:    env.GetBool("debug", false),
			Insecure: env.GetBool("insecure", false),
		}
		request := secret.SecretReadRequest{
			Identifier: identifier,
			Password:   password,
		}
		response, err := secret.SecretRead(context, request)
		if err != nil && err.Code != 200 {
			if !quiet {
				fmt.Printf("\n%s\n\n", err.Message)
			}
			os.Exit(1)
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
