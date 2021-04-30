package internal

import (
	"fmt"
	"strings"

	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/firewall"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

var firewallAllowCmd = &cobra.Command{
	Use:   "allow",
	Short: "Permanently allows a source IP address to a specific port",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Allows a specified IP address to bypass the local system firewall by creating an 'allow' entry into the permanent firewall config.",
			"Grants unprivileged users ability to manipulate the firewall in a safe and controlled manner and keeps an audit log.",
			"Able to control a single (localhost) server as well as cluster of servers.",
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
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		address, _ := cmd.Flags().GetString("address")
		ports, _ := cmd.Flags().GetIntSlice("port")
		comment, _ := cmd.Flags().GetString("comment")
		protocol, _ := cmd.Flags().GetString("protocol")
		selector, _ := cmd.Flags().GetString("type")
		if comment == "" {
			comment = "None"
		}
		rows := [][]string{[]string{"Server", "Response"}}
		runner := func(index, total int, context server.Context) {
			data := firewall.AllowRequest{
				Address:  address,
				Ports:    ports,
				Protocol: protocol,
				Comment:  text.SanitizeString("[^a-zA-Z0-9]+", comment),
			}
			response := firewall.Allow(context, data)
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
	firewallCmd.AddCommand(firewallAllowCmd)
	firewallAllowCmd.Flags().SortFlags = true
	firewallAllowCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
	firewallAllowCmd.Flags().StringP("type", "t", "localhost", "specify server type, useful for cluster")
	firewallAllowCmd.Flags().StringP("address", "a", "", "ip address")
	firewallAllowCmd.Flags().IntSliceP("port", "p", []int{}, "port to allow, can be specified multiple times")
	firewallAllowCmd.Flags().String("protocol", "tcp", "specify 'tcp' or 'udp'")
	firewallAllowCmd.Flags().StringP("comment", "c", "", "add a comment to the firewall entry (optional)")
	firewallAllowCmd.MarkFlagRequired("address")
	firewallAllowCmd.MarkFlagRequired("port")
}
