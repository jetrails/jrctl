package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/firewall"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

var firewallDenyCmd = &cobra.Command{
	Use:   "deny",
	Short: "Permanently denies a source IP address to a specific port",
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
		ports, _ := cmd.Flags().GetIntSlice("port")
		comment, _ := cmd.Flags().GetString("comment")
		protocol, _ := cmd.Flags().GetString("protocol")
		selector, _ := cmd.Flags().GetString("type")
		address, _ := cmd.Flags().GetString("address")
		file, _ := cmd.Flags().GetString("file")
		var addresses []string
		if cmd.Flag("address").Changed {
			addresses = []string{address}
		} else {
			fileContents, fileError := ioutil.ReadFile(file)
			if fileError != nil {
				fmt.Printf("could not read contents of file %q", file)
				os.Exit(1)
			}
			lines := strings.Split(string(fileContents), "\n")
			for _, line := range lines {
				line = strings.Trim(line, " \t\r")
				if line != "" {
					addresses = append(addresses, line)
				}
			}
			if len(addresses) == 0 {
				fmt.Printf("file %q appears to have no entries\n", file)
				os.Exit(1)
			}
		}
		if comment == "" {
			comment = "None"
		}
		rows := [][]string{{"Server", "Response"}}
		runner := func(index, total int, context server.Context) {
			for _, address := range addresses {
				data := firewall.DenyRequest{
					Address:  address,
					Ports:    ports,
					Protocol: protocol,
					Comment:  text.SanitizeString("[^a-zA-Z0-9]+", comment),
				}
				response := firewall.Deny(context, data)
				row := []string{
					strings.TrimSuffix(context.Endpoint, ":27482"),
					response.Messages[0],
				}
				rows = append(rows, row)
			}
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
	firewallCmd.AddCommand(firewallDenyCmd)
	firewallDenyCmd.Flags().SortFlags = true
	firewallDenyCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
	firewallDenyCmd.Flags().StringP("type", "t", "localhost", "specify server type, useful for cluster")
	firewallDenyCmd.Flags().StringP("address", "a", "", "ip address")
	firewallDenyCmd.Flags().StringP("file", "f", "", "use text file with line separated ips")
	firewallDenyCmd.Flags().IntSliceP("port", "p", []int{}, "port to deny, can be specified multiple times")
	firewallDenyCmd.Flags().String("protocol", "tcp", "specify 'tcp' or 'udp', default is 'tcp'")
	firewallDenyCmd.Flags().StringP("comment", "c", "", "add a comment to the firewall entry (optional)")
	firewallDenyCmd.MarkFlagRequired("port")
}
