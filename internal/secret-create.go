package internal

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/jetrails/jrctl/sdk/secret"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/atotto/clipboard"
)

var secretCreateCmd = &cobra.Command {
	Use:   "create",
	Short: "Create a new one-time secret",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"Create a new one-time secret.",
			"A secret's content can be populated by passing a filepath, or it can be manually specified through STDIN.",
			"Optionally, the secret's url can be copied to your clipboard by passing the --clipboard flag!",
		}),
		utils.Paragraph ( [] string {
			"The following environmental variables can be used: JR_PUBLIC_API_ENDPOINT, JR_SECRET_ENDPOINT.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl secret create",
		"jrctl secret create -c -a",
		"jrctl secret create -c -t 60",
		"jrctl secret create -c -p secretpass",
		"jrctl secret create -c -f ~/.ssh/id_rsa.pub",
	}),
	Run: func ( cmd * cobra.Command, args [] string ) {
		var content string = ""
		filepath, _ := cmd.Flags ().GetString ("file")
		copy, _ := cmd.Flags ().GetBool ("clipboard")
		ttl, _ := cmd.Flags ().GetInt ("ttl")
		generate, _ := cmd.Flags ().GetBool ("auto-generate")
		password, _ := cmd.Flags ().GetString ("password")
		if filepath != "" {
			fileContents, error := utils.ReadFile ( filepath )
			if error != nil {
				fmt.Printf ( "\nCould not read contents of file %q.\n\n", filepath )
				return
			}
			content = fileContents
		}
		if content == "" {
			content = utils.PromptContent ("Secret")
		}
		context := secret.PublicApiContext {
			Endpoint: viper.GetString ("public_api_endpoint"),
			Debug: viper.GetBool ("debug"),
		}
		request := secret.SecretCreateRequest {
			Data: content,
			Password: password,
			TTL: ttl,
			AutoGenerate: generate,
		}
		response, error := secret.SecretCreate ( context, request )
		if error.Code != 200 && error.Code != 0 {
			fmt.Printf ( "\n%s\n\n", error.Message )
			return
		}
		url := fmt.Sprintf (
			"https://%s/secret/%s",
			viper.GetString ("secret_endpoint"),
			response.Identifier,
		)
		displayPassword := "None"
		if response.Password != "" {
			displayPassword = response.Password
		}
		if copy {
			clipboard.WriteAll ( url )
		}
		rows := [] [] string {
			[] string { "TTL", "Password", "Secret URL" },
			[] string { strconv.Itoa ( ttl ) + "s", displayPassword, url },
		}
		utils.TablePrint ( "Could not create secret.", rows, 1 )
	},
}

func init () {
	secretCmd.AddCommand ( secretCreateCmd )
	secretCreateCmd.Flags ().SortFlags = true
	secretCreateCmd.Flags ().IntP ( "ttl", "t", 1 * 24 * 60 * 60, "specify custom ttl in seconds" )
	secretCreateCmd.Flags ().BoolP ( "auto-generate", "a", false, "automatically generate password" )
	secretCreateCmd.Flags ().StringP ( "password", "p", "", "protect secret with a password" )
	secretCreateCmd.Flags ().StringP ( "file", "f", "", "use file contents as secret data" )
	secretCreateCmd.Flags ().BoolP ( "clipboard", "c", false, "copy secret url to clipboard" )
}
