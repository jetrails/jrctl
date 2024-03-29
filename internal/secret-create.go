package internal

import (
	"fmt"
	"io/ioutil"

	"github.com/atotto/clipboard"
	"github.com/jetrails/jrctl/pkg/env"
	"github.com/jetrails/jrctl/pkg/input"
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/api"
	"github.com/jetrails/jrctl/sdk/secret"
	"github.com/spf13/cobra"
)

var secretCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new one-time secret",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Create a new one-time secret.",
			"A secret's content can be populated by passing a filepath, or it can be manually specified through STDIN.",
			"Optionally, the secret's url can be copied to your clipboard by passing the --clipboard flag!",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl secret create",
		"jrctl secret create -c -a",
		"jrctl secret create -c -t 60",
		"jrctl secret create -c -p secretpass",
		"jrctl secret create -c -f ~/.ssh/id_rsa.pub",
		"echo 'Hello World' | jrctl secret create",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		filepath, _ := cmd.Flags().GetString("file")
		copy, _ := cmd.Flags().GetBool("clipboard")
		ttl, _ := cmd.Flags().GetInt("ttl")
		generate, _ := cmd.Flags().GetBool("auto-generate")
		password, _ := cmd.Flags().GetString("password")

		tbl := NewTable(Columns{
			"TTL",
			"Password",
			"Secret URL",
		})

		var content string = ""
		if filepath != "" {
			if bytes, err := ioutil.ReadFile(filepath); err != nil {
				tbl.ExitWithMessage(1, "\ncould not read contents of file %q\n\n", filepath)
			} else {
				content = string(bytes)
			}
		}
		if input.HasDataInPipe() {
			content = input.GetPipeData()
		}
		if content == "" {
			content = input.PromptContent("Secret")
		}

		context := secret.PublicApiContext{
			Endpoint: env.GetString("public_api_endpoint", "api-public.jetrails.com"),
			Debug:    env.GetBool("debug", false),
			Insecure: env.GetBool("insecure", false),
		}
		request := secret.SecretCreateRequest{
			Data:         content,
			Password:     password,
			TTL:          ttl,
			AutoGenerate: generate,
		}
		response, err := secret.SecretCreate(context, request)

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
		url := fmt.Sprintf(
			"https://%s/secret/%s",
			env.GetString("secret_endpoint", "secret.jetrails.com"),
			response.Identifier,
		)

		tbl.AddQuietEntry(url)
		tbl.AddRow(Columns{fmt.Sprintf("%ds", ttl), response.Password, url})
		tbl.Quiet = quiet
		tbl.PrintDivider()
		tbl.PrintTable()
		tbl.PrintDivider()

		if copy {
			clipboard.WriteAll(url)
		}
	},
}

func init() {
	secretCmd.AddCommand(secretCreateCmd)
	secretCreateCmd.Flags().SortFlags = true
	secretCreateCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
	secretCreateCmd.Flags().IntP("ttl", "t", 1*24*60*60, "specify custom ttl in seconds")
	secretCreateCmd.Flags().BoolP("auto-generate", "a", false, "automatically generate password")
	secretCreateCmd.Flags().StringP("password", "p", "", "protect secret with a password")
	secretCreateCmd.Flags().StringP("file", "f", "", "use file contents as secret data")
	secretCreateCmd.Flags().BoolP("clipboard", "c", false, "copy secret url to clipboard")
}
