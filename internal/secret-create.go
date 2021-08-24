package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/jetrails/jrctl/pkg/env"
	"github.com/jetrails/jrctl/pkg/input"
	"github.com/jetrails/jrctl/pkg/text"
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
		var content string = ""
		quiet, _ := cmd.Flags().GetBool("quiet")
		filepath, _ := cmd.Flags().GetString("file")
		copy, _ := cmd.Flags().GetBool("clipboard")
		ttl, _ := cmd.Flags().GetInt("ttl")
		generate, _ := cmd.Flags().GetBool("auto-generate")
		password, _ := cmd.Flags().GetString("password")
		if filepath != "" {
			fileContents, error := ioutil.ReadFile(filepath)
			if error != nil {
				if !quiet {
					fmt.Printf("\nCould not read contents of file %q.\n\n", filepath)
				}
				os.Exit(1)
			}
			content = string(fileContents)
		}
		if stat, _ := os.Stdin.Stat(); stat.Mode()&os.ModeCharDevice == 0 {
			if bytes, error := ioutil.ReadAll(os.Stdin); error == nil {
				content = string(bytes)
			}
		}
		if content == "" {
			content = input.PromptContent("Secret")
		}
		context := secret.PublicApiContext{
			Endpoint: env.GetString("public_api_endpoint", "api-public.jetrails.cloud"),
			Debug:    env.GetBool("debug", false),
			Insecure: env.GetBool("insecure", false),
		}
		request := secret.SecretCreateRequest{
			Data:         content,
			Password:     password,
			TTL:          ttl,
			AutoGenerate: generate,
		}
		response, error := secret.SecretCreate(context, request)
		if error != nil && error.Code != 200 {
			if !quiet {
				fmt.Printf("\n%s\n\n", error.Message)
			}
			os.Exit(1)
		}
		url := fmt.Sprintf(
			"https://%s/secret/%s",
			env.GetString("secret_endpoint", "secret.jetrails.cloud"),
			response.Identifier,
		)
		displayPassword := "None"
		if response.Password != "" {
			displayPassword = response.Password
		}
		if copy {
			clipboard.WriteAll(url)
		}
		rows := [][]string{
			{"TTL", "Password", "Secret URL"},
			{strconv.Itoa(ttl) + "s", displayPassword, url},
		}
		if !quiet {
			text.TablePrint("Could not create secret.", rows, 1)
		} else {
			fmt.Println(url)
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
