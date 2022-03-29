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
		port, _ := cmd.Flags().GetInt("port")
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
		rows := [][]string{{"Server", "Response"}}
		runner := func(index, total int, context server.Context) {
			for _, address := range addresses {
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
	firewallUnDenyCmd.Flags().StringP("file", "f", "", "use text file with line separated ips")
	firewallUnDenyCmd.Flags().IntP("port", "p", 0, "port to undeny")
	firewallUnDenyCmd.Flags().String("protocol", "tcp", "specify 'tcp' or 'udp', default is 'tcp'")
	firewallUnDenyCmd.MarkFlagRequired("port")
}
