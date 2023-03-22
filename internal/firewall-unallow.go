package internal

import (
	"fmt"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/firewall"
	"github.com/spf13/cobra"
)

var firewallUnAllowCmd = &cobra.Command{
	Use:   "unallow",
	Short: "Deletes allow entry given a source IP address and a port number",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Removes a 'allow' entry.",
			"Able to control a single node as well as cluster of nodes.",
		}),
	}),
	Example: text.Examples([]string{
		"# Stand-Alone Server",
		"jrctl firewall unallow -a 1.1.1.1 -p 22",
		"",
		"# Multi-Server Cluster",
		"jrctl firewall unallow -t db -a 1.1.1.1 -p 3306",
		"jrctl firewall unallow -t admin -a 1.1.1.1 -p 22",
		"jrctl firewall unallow -t admin -a 1.1.1.1 -p 80,443",
	}),
	Args: func(cmd *cobra.Command, args []string) error {
		ports, _ := cmd.Flags().GetIntSlice("port")
		if len( ports ) < 1 {
			return fmt.Errorf("must pass at least one port with the 'port' flag")
		}
		if !cmd.Flag("address").Changed && !cmd.Flag("file").Changed {
			return fmt.Errorf("must pass either the 'address' or 'file' flag")
		}
		if cmd.Flag("address").Changed && cmd.Flag("file").Changed {
			return fmt.Errorf("cannot pass both the 'address' and 'file' flag")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tags, _ := cmd.Flags().GetStringArray("tag")
		ports, _ := cmd.Flags().GetIntSlice("port")
		protocol, _ := cmd.Flags().GetString("protocol")
		address, _ := cmd.Flags().GetString("address")
		file, _ := cmd.Flags().GetString("file")
		addresses := ResolveAddressInput(file, address, cmd.Flag("address").Changed)

		output := NewOutput(quiet, tags)
		output.DisplayServers = true
		output.FailOnNoServers = true
		output.ExitCodeNoServers = 1

		for _, context := range config.GetContexts(tags) {
			for _, address := range addresses {
				for _, port := range ports {
					request := firewall.UnAllowRequest{
						Address:  address,
						Port:     port,
						Protocol: protocol,
					}
					response := firewall.UnAllow(context, request)
					output.AddServer(
						context,
						response.GetGeneric(),
						response.GetFirstMessage(),
					)
				}
			}
		}

		output.Print()
	},
}

func init() {
	OnlyRunOnNonAWS(firewallUnAllowCmd)
	firewallCmd.AddCommand(firewallUnAllowCmd)
	firewallUnAllowCmd.Flags().SortFlags = true
	firewallUnAllowCmd.Flags().BoolP("quiet", "q", false, "display no input")
	firewallUnAllowCmd.Flags().StringArrayP("tag", "t", []string{"default"}, "filter nodes using tags")
	firewallUnAllowCmd.Flags().StringP("address", "a", "", "ip address")
	firewallUnAllowCmd.Flags().StringP("file", "f", "", "use text file with line separated ips")
	firewallUnAllowCmd.Flags().IntSliceP("port", "p", [] int {}, "port to unallow, can be specified multiple times")
	firewallUnAllowCmd.Flags().String("protocol", "tcp", "specify 'tcp' or 'udp', default is 'tcp'")
	firewallUnAllowCmd.MarkFlagRequired("port")
}
