package internal

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/jetrails/jrctl/pkg/env"
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/api"
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
	Aliases: []string{"upload"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		copy, _ := cmd.Flags().GetBool("clipboard")
		filepath := args[0]

		tbl := NewTable(Columns{"TTL", "Identifier"})
		tbl.Quiet = quiet

		if _, err := ioutil.ReadFile(filepath); err != nil {
			tbl.ExitWithMessage(3, "\ncould not read contents of file %q.\n", filepath)
		}

		context := transfer.PublicApiContext{
			Endpoint: env.GetString("public_api_endpoint", "api-public.jetrails.com"),
			Debug:    env.GetBool("debug", false),
			Insecure: env.GetBool("insecure", false),
		}
		request := transfer.SendRequest{FilePath: filepath}
		response, err := transfer.Send(context, request)

		if err != nil && err.Code != 200 {
			generic := api.GenericResponse{
				Code:     err.Code,
				Status:   err.Type,
				Messages: []string{err.Message},
			}
			tbl.PrintDivider()
			tbl.PrintResponse(&generic)
			tbl.PrintDivider()
			tbl.ExitCodeFromResponse(&generic)
		}

		identifier := fmt.Sprintf("%s-%s", response.Identifier, response.Password)

		tbl.AddQuietEntry(identifier)
		tbl.AddRow(Columns{strconv.Itoa(response.TTL) + "s", identifier})
		tbl.PrintDivider()
		tbl.PrintTable()
		tbl.PrintDivider()

		if copy {
			clipboard.WriteAll(identifier)
		}
	},
}

func init() {
	transferCmd.AddCommand(transferSendCmd)
	transferSendCmd.Flags().SortFlags = true
	transferSendCmd.Flags().BoolP("quiet", "q", false, "only display the identifier")
	transferSendCmd.Flags().BoolP("clipboard", "c", false, "copy file identifier to clipboard")
}
