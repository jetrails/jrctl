package internal

import (
	"fmt"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/firewall"
	"github.com/spf13/cobra"
)

var firewallUnDenyCmd = &cobra.Command{
	Use:   "undeny",
	Short: "Deletes deny entry given a source IP address and a port number",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Removes a 'deny' entry.",
			"Able to control a single (localhost) server as well as cluster of servers.",
		}),
	}),
	Example: text.Examples([]string{
		"# Stand-Alone Server",
		"jrctl firewall undeny -a 1.1.1.1 -p 22",
		"",
		"# Multi-Server Cluster",
		"jrctl firewall undeny -t db -a 1.1.1.1 -p 3306",
		"jrctl firewall undeny -t admin -a 1.1.1.1 -p 22",
	}),
	Args: func(cmd *cobra.Command, args []string) error {
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
		tags, _ := cmd.Flags().GetStringArray("type")
		port, _ := cmd.Flags().GetInt("port")
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
				request := firewall.UnDenyRequest{
					Address:  address,
					Port:     port,
					Protocol: protocol,
				}
				response := firewall.UnDeny(context, request)
				output.AddServer(
					context,
					response.GetGeneric(),
					response.Messages[0],
				)
			}
		}

		output.Print()
	},
}

func init() {
	firewallCmd.AddCommand(firewallUnDenyCmd)
	firewallUnDenyCmd.Flags().SortFlags = true
	firewallUnDenyCmd.Flags().BoolP("quiet", "q", false, "display no input")
	firewallUnDenyCmd.Flags().StringArrayP("type", "t", []string{"localhost"}, "filter servers using type selectors")
	firewallUnDenyCmd.Flags().StringP("address", "a", "", "ip address")
	firewallUnDenyCmd.Flags().StringP("file", "f", "", "use text file with line separated ips")
	firewallUnDenyCmd.Flags().IntP("port", "p", 0, "port to undeny")
	firewallUnDenyCmd.Flags().String("protocol", "tcp", "specify 'tcp' or 'udp', default is 'tcp'")
	firewallUnDenyCmd.MarkFlagRequired("port")
}
