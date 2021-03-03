package internal

import (
	"os"
	"fmt"
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
				utils.PrintErrors ( 1, "Client Side" )
				utils.PrintMessages ( [] string { error.Error () } )
				os.Exit ( 1 )
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
			utils.PrintErrors ( error.Code, error.Type )
			utils.PrintMessages ( [] string { error.Message } )
			os.Exit ( 1 )
		}
		url := fmt.Sprintf (
			"https://%s/secret/%s",
			viper.GetString ("secret_endpoint"),
			response.Identifier,
		)
		fmt.Println ()
		fmt.Printf ( "Identifier:  %s\n", response.Identifier )
		if response.Password != "" {
			fmt.Printf ( "Password:    %s\n", response.Password )
		}
		fmt.Printf ( "TTL:         %d seconds\n", response.TTL )
		fmt.Printf ( "\n%s\n\n", url )
		if copy {
			clipboard.WriteAll ( url )
		}
	},
}

func init () {
	secretCmd.AddCommand ( secretCreateCmd )
	secretCreateCmd.Flags ().SortFlags = false
	secretCreateCmd.Flags ().IntP ( "ttl", "t", 1 * 24 * 60 * 60, "specify custom ttl in seconds" )
	secretCreateCmd.Flags ().BoolP ( "auto-generate", "a", false, "automatically generate password" )
	secretCreateCmd.Flags ().StringP ( "password", "p", "", "protect secret with a password" )
	secretCreateCmd.Flags ().StringP ( "file", "f", "", "use file contents as secret data" )
	secretCreateCmd.Flags ().BoolP ( "clipboard", "c", false, "copy secret url to clipboard" )
}
