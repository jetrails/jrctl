package internal

import (
	"strings"

	"github.com/spf13/cobra"
)

func SplitUserAndHost(input string) (string, string) {
	from := "localhost"
	parts := strings.Split(input, "@")
	user := parts[0]
	if len(parts) > 1 {
		from = parts[1]
	}
	return user, from
}

var databaseCmd = &cobra.Command{
	Use:     "database",
	Aliases: []string{"db"},
	Short:   "Manage databases in deployment",
}

func init() {
	RootCmd.AddCommand(databaseCmd)
	databaseCmd.Flags().SortFlags = true
}
