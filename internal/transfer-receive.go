package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/jetrails/jrctl/pkg/env"
	"github.com/jetrails/jrctl/pkg/input"
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/api"
	"github.com/jetrails/jrctl/sdk/transfer"
	"github.com/spf13/cobra"
)

var transferReceiveCmd = &cobra.Command{
	Use:     "receive",
	Aliases: []string{"download"},
	Short:   "Download file from secure server",
	Example: text.Examples([]string{
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo",
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -f",
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -o ./private/",
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -n custom.name",
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -o ./private/ -n custom.name",
		"echo 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo | jrctl transfer receive",
	}),
	Args: func(cmd *cobra.Command, args []string) error {
		if input.HasDataInPipe() && len(args) == 0 {
			return nil
		}
		return cobra.ExactArgs(1)(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		force, _ := cmd.Flags().GetBool("force")
		outDir, _ := cmd.Flags().GetString("out-dir")
		name, _ := cmd.Flags().GetString("name")
		argument := input.GetFirstArgumentOrPipe(args)

		var password string
		var identifier string
		fmt.Sscanf(argument, "%36s-%32s", &identifier, &password)

		tbl := NewTable(Columns{})
		tbl.Quiet = quiet

		context := transfer.PublicApiContext{
			Endpoint: env.GetString("public_api_endpoint", "api-public.jetrails.com"),
			Debug:    env.GetBool("debug", false),
			Insecure: env.GetBool("insecure", false),
		}
		request := transfer.ReceiveRequest{
			Identifier: identifier,
			Password:   password,
		}
		response, err := transfer.Receive(context, request)

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

		if strings.TrimSpace(outDir) == "" {
			outDir = "."
		}
		if strings.TrimSpace(name) == "" {
			name = response.FileName
		}
		filepath := path.Join(outDir, name)
		os.MkdirAll(outDir, 0755)

		if _, err := os.Stat(filepath); err == nil {
			if !force && (quiet || !input.PromptYesNo("\nfile already exists, overwrite")) {
				tbl.ExitWithMessage(3, "\nskipping, did not write file to disk\n")
			}
		}
		if err := ioutil.WriteFile(filepath, response.Bytes, 0644); err == nil && !quiet {
			tbl.ExitWithMessage(0, "\ndownloaded file to %q\n", filepath)
		} else if !quiet {
			tbl.ExitWithMessage(4, "\nfailed to write data to %q\n", filepath)
		}
	},
}

func init() {
	transferCmd.AddCommand(transferReceiveCmd)
	transferReceiveCmd.Flags().SortFlags = true
	transferReceiveCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
	transferReceiveCmd.Flags().StringP("out-dir", "o", "", "specify download directory, default $PWD")
	transferReceiveCmd.Flags().StringP("name", "n", "", "specify file name")
	transferReceiveCmd.Flags().BoolP("force", "f", false, "force download, overwrite existing file")
}
