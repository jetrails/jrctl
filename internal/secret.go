package internal

import (
	"github.com/spf13/cobra"
	"github.com/jetrails/jrctl/sdk/utils"
)

var secretCmd = &cobra.Command {
	Use:   "secret",
	Short: "Interact with our one-time secret service",
	Example: utils.Examples ([] string {
		"jrctl secret read 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq",
		"jrctl secret read 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq -c",
		"jrctl secret read 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq -c -p secretpass",
		"jrctl secret delete 89ea32e9-e8a5-435d-97ce-78804be250b7-IUQhHYRq",
		"jrctl secret create",
		"jrctl secret create -c -a",
		"jrctl secret create -c -t 60",
		"jrctl secret create -c -p secretpass",
		"jrctl secret create -c -f ~/.ssh/id_rsa.pub",
	}),
}

func init () {
	rootCmd.AddCommand ( secretCmd )
}
