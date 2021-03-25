package internal

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
)

var completionCmd = &cobra.Command {
	Use: "completion [bash|zsh|fish]",
	Short: "Generate completion script",
	Example: utils.Examples ([] string {
		"# bash - Linux",
		"jrctl completion bash > /etc/bash_completion.d/jrctl",
		"# bash - MacOS",
		"jrctl completion bash > /usr/local/etc/bash_completion.d/jrctl",
		"# zsh",
		"jrctl completion zsh > \"${fpath[1]}/_jrctl\"",
		"# fish",
		"jrctl completion fish > ~/.config/fish/completions/jrctl.fish",
	}),
	Hidden: true,
	DisableFlagsInUseLine: true,
	ValidArgs: [] string { "bash", "zsh", "fish" },
	Args: cobra.ExactValidArgs ( 1 ),
	Run: func ( cmd * cobra.Command, args [] string ) {
		switch args [ 0 ] {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				cmd.Root().GenFishCompletion(os.Stdout, true)
		}
	},
}

func init () {
	rootCmd.AddCommand ( completionCmd )
	completionCmd.Flags ().SortFlags = true
}
