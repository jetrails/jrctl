package internal

import (
	"fmt"
	"strings"

	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/firewall"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

var firewallUnDenyCmd = &cobra.Command{
	Use:   "undeny",
	Short: "Deletes deny entry given a source IP address and a port number",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Denies a specified IP address to bypass the local system firewall by creating an 'deny' entry into the permanent firewall config.",
			"Grants unprivileged users ability to manipulate the firewall in a safe and controlled manner and keeps an audit log.",
			"Able to control a single (localhost) server as well as cluster of servers.",
		}),
	}),
	Example: text.Examples([]string{
		"# Stand-Alone Server",
		"jrctl firewall deny -a 1.1.1.1 -p 80 -p 443",
		"",
		"# Multi-Server Cluster",
		"jrctl firewall deny -t db -a 1.1.1.1 -p 3306",
		"jrctl firewall deny -t admin -a 1.1.1.1 -p 22 -c 'Office'",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		address, _ := cmd.Flags().GetString("address")
		port, _ := cmd.Flags().GetInt("port")
		protocol, _ := cmd.Flags().GetString("protocol")
		selector, _ := cmd.Flags().GetString("type")
		rows := [][]string{[]string{"Server", "Response"}}
		runner := func(index, total int, context server.Context) {
			data := firewall.UnDenyRequest{
				Address:  address,
				Port:     port,
				Protocol: protocol,
			}
			response := firewall.UnDeny(context, data)
			row := []string{
				strings.TrimSuffix(context.Endpoint, ":27482"),
				response.Messages[0],
			}
			rows = append(rows, row)
		}
		server.FilterForEach([]string{selector}, runner)
		if !quiet {
			if len(rows) > 1 {
				fmt.Printf("\nExecuted only on %q server(s):\n", selector)
			}
			text.TablePrint(fmt.Sprintf("No configured %q server(s) found.", selector), rows, 1)
		}
	},
}

func init() {
	firewallCmd.AddCommand(firewallUnDenyCmd)
	firewallUnDenyCmd.Flags().SortFlags = true
	firewallUnDenyCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
	firewallUnDenyCmd.Flags().StringP("type", "t", "localhost", "specify server type, useful for cluster")
	firewallUnDenyCmd.Flags().StringP("address", "a", "", "ip address")
	firewallUnDenyCmd.Flags().IntP("port", "p", 0, "port to undeny")
	firewallUnDenyCmd.Flags().String("protocol", "tcp", "specify 'tcp' or 'udp', default is 'tcp'")
	firewallUnDenyCmd.MarkFlagRequired("address")
	firewallUnDenyCmd.MarkFlagRequired("port")
}
