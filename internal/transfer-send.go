package internal

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/env"
	"github.com/jetrails/jrctl/sdk/transfer"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/atotto/clipboard"
)

var transferSendCmd = &cobra.Command {
	Use:   "send",
	Short: "Upload file to secure server",
	Long: utils.Combine ( [] string {
		utils.Paragraph ( [] string {
			"Upload file to secure server.",
			"File is encrypted and stored for an hour.",
		}),
	}),
	Example: utils.Examples ([] string {
		"jrctl transfer send private.png",
	}),
	Args: cobra.ExactArgs ( 1 ),
	Run: func ( cmd * cobra.Command, args [] string ) {
		filepath := args [ 0 ]
		copy, _ := cmd.Flags ().GetBool ("clipboard")
		if _, error := utils.ReadFile ( filepath ); error != nil {
			fmt.Printf ( "\nCould not read contents of file %q.\n\n", filepath )
			return
		}
		context := transfer.PublicApiContext {
			Endpoint: env.GetString ( "public_api_endpoint", "api-public.jetrails.cloud" ),
			Debug: env.GetBool ( "debug", false ),
			Insecure: env.GetBool ( "insecure", false ),
		}
		request := transfer.SendRequest { FilePath: filepath }
		response, error := transfer.Send ( context, request )
		if error != nil && error.Code != 200 {
			fmt.Printf ( "\n%s\n\n", error.Message )
			return
		}
		identifier := fmt.Sprintf ( "%s-%s", response.Identifier, response.Password )
		if copy {
			clipboard.WriteAll ( identifier )
		}
		rows := [] [] string {
			[] string { "TTL", "Identifier" },
			[] string { strconv.Itoa ( response.TTL ) + "s", identifier },
		}
		utils.TablePrint ( "Could not send file.", rows, 1 )
	},
}

func init () {
	transferCmd.AddCommand ( transferSendCmd )
	transferSendCmd.Flags ().SortFlags = true
	transferSendCmd.Flags ().BoolP ( "clipboard", "c", false, "copy file identifier to clipboard" )
}