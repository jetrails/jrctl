package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/jetrails/jrctl/pkg/env"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/transfer"
	"github.com/spf13/cobra"
)

var transferSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Upload file to secure server",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Upload file to secure server.",
			"File is encrypted and stored for an hour.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl transfer send private.png",
	}),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		filepath := args[0]
		copy, _ := cmd.Flags().GetBool("clipboard")
		if _, error := ioutil.ReadFile(filepath); error != nil {
			fmt.Printf("\nCould not read contents of file %q.\n\n", filepath)
			return
		}
		context := transfer.PublicApiContext{
			Endpoint: env.GetString("public_api_endpoint", "api-public.jetrails.cloud"),
			Debug:    env.GetBool("debug", false),
			Insecure: env.GetBool("insecure", false),
		}
		request := transfer.SendRequest{FilePath: filepath}
		response, error := transfer.Send(context, request)
		if error != nil && error.Code != 200 {
			if !quiet {
				fmt.Printf("\n%s\n\n", error.Message)
			}
			os.Exit(1)
		}
		identifier := fmt.Sprintf("%s-%s", response.Identifier, response.Password)
		if copy {
			clipboard.WriteAll(identifier)
		}
		rows := [][]string{
			[]string{"TTL", "Identifier"},
			[]string{strconv.Itoa(response.TTL) + "s", identifier},
		}
		if !quiet {
			text.TablePrint("Could not send file.", rows, 1)
		} else {
			fmt.Println(identifier)
		}
	},
}

func init() {
	transferCmd.AddCommand(transferSendCmd)
	transferSendCmd.Flags().SortFlags = true
	transferSendCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
	transferSendCmd.Flags().BoolP("clipboard", "c", false, "copy file identifier to clipboard")
}
