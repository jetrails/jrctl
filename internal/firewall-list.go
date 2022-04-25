package internal

import (
	"fmt"
	"strings"

	"github.com/jetrails/jrctl/pkg/array"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/firewall"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

func squashOrAppendEntry(rows [][]string, entry []string, varyIndex int) [][]string {
	for rowIndex, row := range rows {
		allButVaryIsSame := true
		for colIndex, col := range row {
			if entry[colIndex] != col && colIndex != varyIndex {
				allButVaryIsSame = false
				break
			}
		}
		if allButVaryIsSame {
			if !array.ContainsString(strings.Split(rows[rowIndex][varyIndex], ", "), entry[varyIndex]) {
				rows[rowIndex][varyIndex] = rows[rowIndex][varyIndex] + ", " + entry[varyIndex]
			}
			return rows
		}
	}
	return append(rows, entry)
}

var firewallListCmd = &cobra.Command{
	Use:   "list",
	Short: "List firewall entries across configured servers",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"List firewall entries across configured servers.",
			"Specifing a server type will only display results for servers of that type.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl firewall list",
		"jrctl firewall list -t admin",
		"jrctl firewall list -t db",
		"jrctl firewall list -t www",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		selector, _ := cmd.Flags().GetString("type")
		responseRows := [][]string{{"Hostname", "Server", "Type(s)", "Response"}}
		entryTables := [][][]string{}
		filter := []string{}
		emptyMsg := "No configured servers found."
		if selector != "" {
			filter = []string{selector}
			emptyMsg = fmt.Sprintf("No configured %q server(s) found.", selector)
		}
		runner := func(index, total int, context server.Context) {
			entryRows := [][]string{{"Hostname", "Type(s)", "Action", "IPV4/CIDR", "Port(s)", "Protocol(s)", "Comment"}}
			response := firewall.List(context)
			responseRow := []string{
				response.Metadata["Hostname"],
				strings.TrimSuffix(context.Endpoint, ":27482"),
				strings.Join(context.Types, ", "),
				response.Messages[0],
			}
			responseRows = append(responseRows, responseRow)
			for _, entry := range response.Payload {
				commentEnd := strings.Index(entry.Comment, " -- ")
				if commentEnd == -1 {
					commentEnd = len(entry.Comment)
				}
				entryRow := []string{
					response.Metadata["Hostname"],
					strings.Join(context.Types, ", "),
					entry.Action,
					entry.Source,
					strings.Trim(strings.Join(strings.Fields(fmt.Sprint(entry.Ports)), ", "), "[]"),
					strings.Join(entry.Protocols, ", "),
					fmt.Sprintf("%q", strings.ReplaceAll(entry.Comment[:commentEnd], "_", " ")),
				}
				entryRows = squashOrAppendEntry(entryRows, entryRow, 4)
			}
			entryTables = append(entryTables, entryRows)
		}
		server.FilterForEach(filter, runner)
		if selector != "" && len(responseRows) > 1 {
			fmt.Printf("\nDisplaying results for %q server(s):\n", selector)
		}
		fmt.Println()
		text.TablePrint(emptyMsg, responseRows, 0)
		fmt.Println()
		if len(responseRows) > 1 {
			for _, table := range entryTables {
				text.TablePrint("No firewall entries found.", table, 0)
				fmt.Println()
			}
		}
	},
}

func init() {
	firewallCmd.AddCommand(firewallListCmd)
	firewallListCmd.Flags().SortFlags = true
	firewallListCmd.Flags().StringP("type", "t", "", "specify server type selector")
}
