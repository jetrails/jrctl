package internal

import (
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
)

var firewallCmd = &cobra.Command{
	Use:   "firewall",
	Short: "Interact with server firewall",
}

func ResolveAddressInput(file, address string, addressChanged bool) []string {
	var addresses []string
	if addressChanged {
		addresses = []string{address}
	} else {
		fileContents, _ := ioutil.ReadFile(file)
		lines := strings.Split(string(fileContents), "\n")
		for _, line := range lines {
			line = strings.Trim(line, " \t\r")
			if line != "" {
				addresses = append(addresses, line)
			}
		}
	}
	return addresses
}

func init() {
	OnlyRunOnNonAWS(firewallCmd)
	RootCmd.AddCommand(firewallCmd)
	firewallCmd.Flags().SortFlags = true
}
