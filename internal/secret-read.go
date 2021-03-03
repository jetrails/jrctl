package internal

import (
	"fmt"
	"strings"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/secret"
	"github.com/atotto/clipboard"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

var secretReadCmd = &cobra.Command {
	Use: "read IDENTIFIER",
	Short: "Display contents of secret",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"Display contents of secret.",
			"Passing the secret identifier will allow us to retrieve the contents of the secret and print it to STDOUT.",
			"Optionally, you can copy the contents to your clipboard by passing the --clipboard flag!",
		}),
		utils.Paragraph ( [] string {
			"The following environmental variables can be used: JR_PUBLIC_API_ENDPOINT.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl secret read 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq",
		"jrctl secret read 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq -c",
		"jrctl secret read 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq -c -p secretpass",
	}),
	Args: func ( cmd * cobra.Command, args [] string ) error {
		check := cobra.MinimumNArgs ( 1 )
		if error := check ( cmd, args ); error != nil {
			return error
		}
		return nil
	},
	Run: func ( cmd * cobra.Command, args [] string ) {
		identifier := args [ 0 ]
		copy, _ := cmd.Flags ().GetBool ("clipboard")
		password, _ := cmd.Flags ().GetString ("password")
		context := secret.PublicApiContext {
			Endpoint: viper.GetString ("public_api_endpoint"),
			Debug: viper.GetBool ("debug"),
		}
		request := secret.SecretReadRequest {
			Identifier: identifier,
			Password: password,
		}
		response, error := secret.SecretRead ( context, request )
		if error.Code != 200 && error.Code != 0 {
			utils.PrintErrors ( error.Code, error.Type )
			utils.PrintMessages ( [] string { error.Message } )
		} else {
			fmt.Printf ( "\n%s\n\n", response.Data )
		}
		if copy {
			clipboard.WriteAll ( strings.TrimSpace ( response.Data ) )
		}
	},
}

func init () {
	secretCmd.AddCommand ( secretReadCmd )
	secretReadCmd.Flags ().SortFlags = false
	secretReadCmd.Flags ().StringP ( "password", "p", "", "password to access secret" )
	secretReadCmd.Flags ().BoolP ( "clipboard", "c", false, "copy contents to clipboard" )
}
