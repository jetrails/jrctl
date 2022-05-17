package internal

import (
	"fmt"
	"os"

	"github.com/jetrails/jrctl/pkg/color"
	"github.com/jetrails/jrctl/pkg/env"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/version"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check to see if there is an update available",
	Example: text.Examples([]string{
		"jrctl server update",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		updateToDate, latest := version.CheckVersion(env.GetBool("debug", false))
		if latest != "" {
			if updateToDate {
				if !quiet {
					fmt.Printf("\nSoftware is up-to-date.\n\n")
				}
				os.Exit(0)
			}
			if !quiet {
				fmt.Printf(
					"\nSoftware is out-of-date.\nPlease update to the latest version: %s.\n\n",
					color.GreenString(fmt.Sprintf(version.TagUrlTemplate, latest)),
				)
			}
			os.Exit(1)
		}
		if !quiet {
			fmt.Printf(
				"\nFailed to query latest version.\nPlease check manually: %s\n\n",
				version.ReleasesUrl,
			)
		}
		os.Exit(2)
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
	updateCmd.Flags().SortFlags = true
	updateCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
}
