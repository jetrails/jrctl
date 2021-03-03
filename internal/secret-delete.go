package internal

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/secret"
)

var secretDeleteCmd = &cobra.Command {
	Use:   "delete IDENTIFIER",
	Short: "Delete secret without viewing contents",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"Delete secret without viewing contents.",
			"Passing the secret identifier will make a request to destroy the secret without displaying the secret's contents.",
		}),
		utils.Paragraph ( [] string {
			"The following environmental variables can be used: JR_PUBLIC_API_ENDPOINT.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl secret delete 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq",
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
		context := secret.PublicApiContext {
			Endpoint: viper.GetString ("public_api_endpoint"),
			Debug: viper.GetBool ("debug"),
		}
		request := secret.SecretDeleteRequest {
			Identifier: identifier,
		}
		response, error := secret.SecretDelete ( context, request )
		if error.Code != 200 && error.Code != 0 {
			utils.PrintErrors ( error.Code, error.Type )
			utils.PrintMessages ( [] string { error.Message } )
			os.Exit ( 1 )
		}
		fmt.Printf ( "Successfully deleted secret: '%s'\n", response.Identifier )
	},
}

func init () {
	secretCmd.AddCommand ( secretDeleteCmd )
	secretDeleteCmd.Flags ().SortFlags = false
}
