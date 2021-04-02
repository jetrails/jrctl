package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/jetrails/jrctl/sdk/env"
	"github.com/jetrails/jrctl/sdk/transfer"
	"github.com/jetrails/jrctl/sdk/utils"
	"github.com/spf13/cobra"
)

var transferReceiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Download file from secure server",
	Long: utils.Combine([]string{
		utils.Paragraph([]string{
			"Download file from secure server.",
		}),
	}),
	Example: utils.Examples([]string{
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo",
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -f",
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -o ./private/",
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -n custom.name",
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -o ./private/ -n custom.name",
	}),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var password string
		var identifier string
		fmt.Sscanf(args[0], "%36s-%32s", &identifier, &password)
		context := transfer.PublicApiContext{
			Endpoint: env.GetString("public_api_endpoint", "api-public.jetrails.cloud"),
			Debug:    env.GetBool("debug", false),
			Insecure: env.GetBool("insecure", false),
		}
		request := transfer.ReceiveRequest{
			Identifier: identifier,
			Password:   password,
		}
		response, error := transfer.Receive(context, request)
		if error != nil && error.Code != 200 {
			fmt.Printf("\n%s\n\n", error.Message)
			return
		}
		force, _ := cmd.Flags().GetBool("force")
		outDir, _ := cmd.Flags().GetString("out-dir")
		name, _ := cmd.Flags().GetString("name")
		if strings.TrimSpace(outDir) == "" {
			outDir = "."
		}
		if strings.TrimSpace(name) == "" {
			name = response.FileName
		}
		filepath := path.Join(outDir, name)
		os.MkdirAll(outDir, 0755)
		fmt.Println()
		if _, error := os.Stat(filepath); error == nil {
			if !force && !utils.PromptYesNo("File already exists, overwrite") {
				fmt.Printf("Skipping, did not write file to disk\n\n")
				return
			}
		}
		if error := ioutil.WriteFile(filepath, response.Bytes, 0644); error == nil {
			fmt.Printf("Downloaded file to %q\n\n", filepath)
		} else {
			fmt.Printf("Failed to write data to %q\n\n", filepath)
		}
	},
}

func init() {
	transferCmd.AddCommand(transferReceiveCmd)
	transferReceiveCmd.Flags().SortFlags = true
	transferReceiveCmd.Flags().StringP("out-dir", "o", "", "specify download directory, default $PWD")
	transferReceiveCmd.Flags().StringP("name", "n", "", "specify file name")
	transferReceiveCmd.Flags().BoolP("force", "f", false, "force download, overwrite existing file")
}
