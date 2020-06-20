package internal

import "fmt"
import "github.com/spf13/cobra"
import "github.com/jetrails/jrctl/sdk/utils"
import "github.com/jetrails/jrctl/sdk/command/secret"

var deleteCmd = &cobra.Command {
	Use:   "delete <identifier>",
	Short: "Delete secret without viewing contents",
	Args: func ( cmd *cobra.Command, args [] string ) error {
		check := cobra.MinimumNArgs ( 1 )
		if error := check ( cmd, args ); error != nil {
			return error
		}
		return nil
	},
	Run: func ( cmd *cobra.Command, args [] string ) {
		var request = secret.SecretDeleteRequest {
			Identifier: args [ 0 ],
		}
		response, error := secret.SecretDelete ( request )
		utils.HandleErrorResponse ( error )
		fmt.Printf ( "Successfully deleted secret: '%s'\n", response.Identifier )
	},
}

func init () {
	secretCmd.AddCommand ( deleteCmd )
}
