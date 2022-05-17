package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/jetrails/jrctl/pkg/env"
	"github.com/jetrails/jrctl/pkg/input"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/transfer"
	"github.com/spf13/cobra"
)

var transferReceiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Download file from secure server",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Download file from secure server.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo",
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -f",
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -o ./private/",
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -n custom.name",
		"jrctl transfer receive 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo -o ./private/ -n custom.name",
		"echo 7c6acde6-639c-47fe-8ebb-a4ac877ef72b-XPlEYzcsgnNbxwcFqKiWUoJil6MlFXGo | jrctl transfer receive",
	}),
	Aliases: []string{"download"},
	Args: func(cmd *cobra.Command, args []string) error {
		if stat, _ := os.Stdin.Stat(); stat.Mode()&os.ModeCharDevice == 0 && len(args) == 0 {
			return nil
		}
		if error := cobra.ExactArgs(1)(cmd, args); error != nil {
			return error
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		var password string
		var identifier string
		var argument string
		if len(args) == 0 {
			if bytes, error := ioutil.ReadAll(os.Stdin); error == nil {
				argument = strings.TrimSpace(string(bytes))
			}
		} else {
			argument = args[0]
		}
		fmt.Sscanf(argument, "%36s-%32s", &identifier, &password)
		context := transfer.PublicApiContext{
			Endpoint: env.GetString("public_api_endpoint", "api-public.jetrails.com"),
			Debug:    env.GetBool("debug", false),
			Insecure: env.GetBool("insecure", false),
		}
		request := transfer.ReceiveRequest{
			Identifier: identifier,
			Password:   password,
		}
		response, error := transfer.Receive(context, request)
		if error != nil && error.Code != 200 {
			if !quiet {
				fmt.Printf("\n%s\n\n", error.Message)
			}
			os.Exit(1)
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
		if !quiet {
			fmt.Println()
		}
		if _, error := os.Stat(filepath); error == nil {
			if !force && (quiet || !input.PromptYesNo("File already exists, overwrite")) {
				if !quiet {
					fmt.Printf("Skipping, did not write file to disk\n\n")
				}
				os.Exit(1)
			}
		}
		if error := ioutil.WriteFile(filepath, response.Bytes, 0644); error == nil && !quiet {
			fmt.Printf("Downloaded file to %q\n\n", filepath)
		} else if !quiet {
			fmt.Printf("Failed to write data to %q\n\n", filepath)
			os.Exit(1)
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
