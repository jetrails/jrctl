package internal

import (
	"fmt"
	"strings"

	"github.com/jetrails/jrctl/pkg/input"
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/confirm"
	"github.com/spf13/cobra"
)

func PrintConfirmMessage(tags []string, identifier, ttl string) {
	selectors := ""
	for _, tag := range tags {
		selectors += " -t " + tag
	}
	fmt.Println("")
	fmt.Println("WARNING: This is a destructive command that cannot be undone. If you would")
	fmt.Println("like to continue, you will need to send a confirmation to the server to")
	fmt.Printf("execute this destructive command (confirmation code will exipire in %s).\n", ttl)
	fmt.Println("")
	fmt.Println("To confirm, run the following command:")
	fmt.Printf("jrctl confirm%s %s\n", selectors, identifier)
	fmt.Println("")
}

var confirmCmd = &cobra.Command{
	Use:   "confirm",
	Short: "Execute queued jobs that require confirmation",
	Example: text.Examples([]string{
		"jrctl confirm e01b1e47-c12d-453f-9905-d25fcc6c3eed",
		"echo e01b1e47-c12d-453f-9905-d25fcc6c3eed | jrctl confirm",
	}),
	Hidden: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 && input.HasDataInPipe() {
			return nil
		}
		return cobra.ExactArgs(1)(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("tag")
		identifier := input.GetFirstArgumentOrPipe(args)

		output := NewOutput(quiet, tags)
		contexts := config.GetContexts(tags)

		output.PrintTags()
		output.PrintDivider()

		if len(contexts) < 1 {
			output.ExitWithMessage(1, ErrMsgNoServers+"\n")
		}

		if len(contexts) > 1 {
			output.ExitWithMessage(5, ErrMsgRequiresOneServer+"\n")
		}

		request := confirm.ConfirmRequest{Identifier: identifier}
		response := confirm.Confirm(contexts[0], request)
		generic := response.GetGeneric()

		output.PrintResponse(generic)
		output.PrintDivider()

		columns := response.Metadata["columns"]
		values := response.Metadata["values"]
		quietValues := response.Metadata["quiet"]
		if columns != "" && values != "" {
			tbl := NewTable(strings.Split(columns, ","))
			tbl.Quiet = quiet
			if quietValues != "" {
				tbl.QuietCollection = strings.Split(quietValues, ",")
			}
			for _, row := range strings.Split(values, ";") {
				tbl.AddRow(strings.Split(row, ","))
			}
			tbl.PrintTable()
			output.PrintDivider()
		}

		output.ExitCodeFromResponse(generic)
	},
}

func init() {
	RootCmd.AddCommand(confirmCmd)
	confirmCmd.Flags().SortFlags = true
	confirmCmd.Flags().BoolP("quiet", "q", false, "display minimal output, if any")
	confirmCmd.Flags().StringArrayP("tag", "t", []string{"default"}, "filter nodes using tags")
}
