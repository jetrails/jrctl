package internal

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/confirm"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

func getArgOrPipe(args []string) string {
	if len(args) == 0 {
		if bytes, err := ioutil.ReadAll(os.Stdin); err == nil {
			return strings.TrimSpace(string(bytes))
		}
	} else {
		return args[0]
	}
	return ""
}

var confirmCmd = &cobra.Command{
	Use:   "confirm",
	Short: "",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			".",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl confirm e01b1e47-c12d-453f-9905-d25fcc6c3eed",
		"echo e01b1e47-c12d-453f-9905-d25fcc6c3eed | jrctl confirm",
	}),
	Args: func(cmd *cobra.Command, args []string) error {
		if stat, _ := os.Stdin.Stat(); stat.Mode()&os.ModeCharDevice == 0 && len(args) == 0 {
			return nil
		}
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		input := getArgOrPipe(args)
		if len(input) != 45 {
			return errors.New("confirmation code must be 45 chars in length")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		input := getArgOrPipe(args)
		hash := input[0:8]
		identifier := input[9:]
		runner := func(index, total int, context server.Context) {
			if context.Hash() == hash {
				data := confirm.ConfirmRequest{
					Identifier: identifier,
				}
				response := confirm.Confirm(context, data)
				if response.Status == "OK" {
					if !quiet {
						fmt.Println(response)
					}
					os.Exit(0)
				} else {
					if !quiet {
						fmt.Println(response)
					}
					os.Exit(1)
				}
			}
		}
		server.ForEach(runner)
	},
}

func init() {
	RootCmd.AddCommand(confirmCmd)
	confirmCmd.Flags().SortFlags = true
	confirmCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
}
