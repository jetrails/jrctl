package internal

import "fmt"
import "github.com/spf13/cobra"
import "github.com/spf13/viper"
import "github.com/jetrails/jrctl/sdk/command/secret"
import "github.com/jetrails/jrctl/sdk/utils"
import "github.com/atotto/clipboard"

var secretCreateCmd = &cobra.Command {
	Use:   "create",
	Short: "Create a new one-time secret",
	Run: func ( cmd *cobra.Command, args [] string ) {
		var postfix = viper.GetString ("endpoint_postfix")
		var content = ""
		filePath, _ := cmd.Flags ().GetString ("file")
		copyToClipBoard, _ := cmd.Flags ().GetBool ("clipboard")
		ttl, _ := cmd.Flags ().GetInt ("ttl")
		password, _ := cmd.Flags ().GetString ("password")
		autoGenerate, _ := cmd.Flags ().GetBool ("auto-generate")
		if filePath != "" {
			content = utils.ReadFile ( filePath )
		}
		if content == "" {
			content = utils.PromptContent ("Secret")
		}
		var request = secret.SecretCreateRequest {
			Data: content,
			Password: password,
			TTL: ttl,
			AutoGenerate: autoGenerate,
		}
		response, error := secret.SecretCreate ( request )
		utils.HandleErrorResponse ( error )
		var url = fmt.Sprintf ( "https://secret%s.jetrails.cloud/secret/%s", postfix, response.Identifier )
		fmt.Println ()
		fmt.Printf ( "Identifier:  %s\n", response.Identifier )
		fmt.Printf ( "Password:    %s\n", response.Password )
		fmt.Printf ( "TTL:         %d seconds\n", response.TTL )
		fmt.Printf ( "\n%s\n\n", url )
		if copyToClipBoard {
			clipboard.WriteAll ( url )
		}
	},
}

func init () {
	secretCmd.AddCommand ( secretCreateCmd )
	secretCreateCmd.Flags ().IntP ( "ttl", "t", 1 * 24 * 60 * 60, "specify custom ttl in seconds" )
	secretCreateCmd.Flags ().BoolP ( "auto-generate", "a", false, "automatically generate password" )
	secretCreateCmd.Flags ().StringP ( "password", "p", "", "specify custom password for secret, required" )
	secretCreateCmd.Flags ().StringP ( "file", "f", "", "use file contents as secret data" )
	secretCreateCmd.Flags ().BoolP ( "clipboard", "c", false, "copy secret url to clipboard" )
	secretCreateCmd.MarkFlagRequired ("password")
}
