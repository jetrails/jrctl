package internal

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/jetrails/jrctl/sdk/secret"
)

var deleteCmd = &cobra.Command {
	Use:   "delete <identifier>",
	Short: "Delete secret without viewing contents",
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
	secretCmd.AddCommand ( deleteCmd )
}
