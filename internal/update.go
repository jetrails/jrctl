package internal

import (
	"fmt"

	"github.com/jetrails/jrctl/pkg/color"
	"github.com/jetrails/jrctl/pkg/env"
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"upgrade"},
	Short:   "Check for available updates",
	Example: text.Examples([]string{
		"jrctl update",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")

		output := NewOutput(quiet, []string{})
		updateToDate, latest := version.CheckVersion(env.GetBool("debug", false))

		if latest == "" {
			output.ExitWithMessage(2, fmt.Sprintf(
				"\nFailed to query latest version\nPlease check manually: %s\n",
				version.ReleasesUrl,
			))
		}
		if updateToDate {
			output.ExitWithMessage(0, "\nSoftware is up-to-date\n")
		}
		output.ExitWithMessage(1, fmt.Sprintf(
			"\nSoftware is out-of-date\nPlease update to the latest version: %s.\n",
			color.GreenString(fmt.Sprintf(version.TagUrlTemplate, latest)),
		))
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
	updateCmd.Flags().SortFlags = true
	updateCmd.Flags().BoolP("quiet", "q", false, "display no output")
}
