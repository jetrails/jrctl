package internal

import (
	"fmt"

	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/firewall"
	"github.com/spf13/cobra"
)

var firewallAllowCmd = &cobra.Command{
	Use:   "allow",
	Short: "Permanently allows a source IP address to a specific port",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Allows a specified IP address to bypass the local system firewall by creating an 'allow' entry into the permanent firewall config.",
			"Grants unprivileged users ability to manipulate the firewall in a safe and controlled manner and keeps an audit log.",
			"Able to control a single node as well as cluster of nodes.",
		}),
	}),
	Example: text.Examples([]string{
		"# Stand-Alone Server",
		"jrctl firewall allow -a 1.1.1.1 -p 80 -p 443",
		"",
		"# Multi-Server Cluster",
		"jrctl firewall allow -t db -a 1.1.1.1 -p 3306",
		"jrctl firewall allow -t admin -a 1.1.1.1 -p 22 -c 'Office'",
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
		tags, _ := cmd.Flags().GetStringArray("tag")
		ports, _ := cmd.Flags().GetIntSlice("port")
		comment, _ := cmd.Flags().GetString("comment")
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
				request := firewall.AllowRequest{
					Address:  address,
					Ports:    ports,
					Protocol: protocol,
					Comment:  text.SanitizeString("[^a-zA-Z0-9]+", comment),
				}
				response := firewall.Allow(context, request)
				output.AddServer(
					context,
					response.GetGeneric(),
					response.GetFirstMessage(),
				)
			}
		}

		output.Print()
	},
}

func init() {
	OnlyRunOnNonAWS(firewallAllowCmd)
	firewallCmd.AddCommand(firewallAllowCmd)
	firewallAllowCmd.Flags().SortFlags = true
	firewallAllowCmd.Flags().BoolP("quiet", "q", false, "display no input")
	firewallAllowCmd.Flags().StringArrayP("tag", "t", []string{"default"}, "filter nodes using tags")
	firewallAllowCmd.Flags().StringP("address", "a", "", "ip address")
	firewallAllowCmd.Flags().StringP("file", "f", "", "use text file with line separated ips")
	firewallAllowCmd.Flags().IntSliceP("port", "p", []int{}, "port to allow, can be specified multiple times")
	firewallAllowCmd.Flags().String("protocol", "tcp", "specify 'tcp' or 'udp'")
	firewallAllowCmd.Flags().StringP("comment", "c", "NA", "add a comment to the firewall entry")
	firewallAllowCmd.MarkFlagRequired("port")
}
