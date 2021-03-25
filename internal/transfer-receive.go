package internal

import (
	"fmt"
	"strings"
	"io/ioutil"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/env"
	"github.com/jetrails/jrctl/sdk/transfer"
	"github.com/jetrails/jrctl/sdk/utils"
)

var transferReceiveCmd = &cobra.Command {
	Use:   "receive",
	Short: "Download file from secure server",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"Download file from secure server.",
			"If no output path is specified, then the file is stored in the current directory and will be named after the file identifier.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo",
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -o private.png",
	}),
	Args: cobra.ExactArgs ( 1 ),
	Run: func ( cmd * cobra.Command, args [] string ) {
		var password string
		var identifier string
		fmt.Sscanf ( args [ 0 ], "%36s-%32s", &identifier, &password )
		output, _ := cmd.Flags ().GetString ("output")
		if strings.TrimSpace ( output ) == "" {
			output = "./" + identifier
		}
		context := transfer.PublicApiContext {
			Endpoint: env.GetString ( "public_api_endpoint", "api-public.jetrails.cloud" ),
			Debug: env.GetBool ( "debug", false ),
			Insecure: env.GetBool ( "insecure", false ),
		}
		request := transfer.ReceiveRequest {
			Identifier: identifier,
			Password: password,
		}
		response, error := transfer.Receive ( context, request )
		if error != nil && error.Code != 200 {
			fmt.Printf ( "\n%s\n\n", error.Message )
			return
		}
		if error := ioutil.WriteFile ( output, response.Bytes, 0644 ); error == nil {
			fmt.Printf ( "\nDownloaded file to %q\n\n", output )
		} else {
			fmt.Printf ( "\nFailed to write data to %q\n\n", output )
		}
	},
}

func init () {
	transferCmd.AddCommand ( transferReceiveCmd )
	transferReceiveCmd.Flags ().SortFlags = true
	transferReceiveCmd.Flags ().StringP ( "output", "o", "", "specify output file path" )
}