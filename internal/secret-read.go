package internal

import "fmt"
import "strings"
import "github.com/jetrails/jrctl/sdk/utils"
import "github.com/jetrails/jrctl/sdk/command/secret"
import "github.com/atotto/clipboard"
import "github.com/spf13/cobra"

var readCmd = &cobra.Command {
	Use: "read <identifier>",
	Short: "Display contents of secret",
	Args: func ( cmd *cobra.Command, args [] string ) error {
		check := cobra.MinimumNArgs ( 1 )
		if error := check ( cmd, args ); error != nil {
			return error
		}
		return nil
	},
	Run: func ( cmd *cobra.Command, args [] string ) {
		identifier := args [ 0 ]
		copyToClipBoard, _ := cmd.Flags ().GetBool ("clipboard")
		password, _ := cmd.Flags ().GetString ("password")
		password = utils.PromptPassword ( "Enter Secret Password: ", password )
		var request = secret.SecretReadRequest {
			Identifier: identifier,
			Password: password,
		}
		response, error := secret.SecretRead ( request )
		utils.HandleErrorResponse ( error )
		fmt.Printf ( "\n%s\n\n", response.Data )
		if copyToClipBoard {
			clipboard.WriteAll ( strings.TrimSpace ( response.Data ) )
		}
	},
}

func init () {
	secretCmd.AddCommand ( readCmd )
	readCmd.Flags ().StringP ( "password", "p", "", "secret password (prompted if not supplied)" )
	readCmd.Flags ().BoolP ( "clipboard", "c", false, "copy contents to clipboard" )
}
