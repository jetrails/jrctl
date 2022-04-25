package internal

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/jetrails/jrctl/pkg/array"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/report"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
)

var reportAuditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Month-to-date security audit to ensure access is properly limited",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Month-to-date security audit to ensure access is properly limited.",
		}),
	}),
	Example: text.Examples([]string{
		"jrctl report audit",
		"jrctl report audit -t www",
		"jrctl report audit -o json",
	}),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		output, _ := cmd.Flags().GetString("output")
		validOutput := []string{"table", "json"}
		if !array.ContainsString(validOutput, output) {
			return fmt.Errorf("output must be one of %v", validOutput)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		selector, _ := cmd.Flags().GetString("type")
		output, _ := cmd.Flags().GetString("output")
		rawResponses := []report.AuditResponse{}
		responseRows := [][]string{{"Hostname", "Server", "Type(s)"}}
		whitelistRows := [][]string{}
		activityRows := [][]string{}
		accessRows := [][]string{}
		filter := []string{}
		emptyMsg := "No configured servers found."
		if selector != "" {
			filter = []string{selector}
			emptyMsg = fmt.Sprintf("No configured %q server(s) found.", selector)
		}
		runner := func(index, total int, context server.Context) {
			response := report.Audit(context)
			rawResponses = append(rawResponses, response)
			if response.Status == "OK" {
				responseRow := []string{
					response.Metadata["hostname"].(string),
					strings.TrimSuffix(context.Endpoint, ":27482"),
					strings.Join(context.Types, ", "),
				}
				for _, entry := range response.Payload.Activity {
					entryRow := []string{
						response.Metadata["hostname"].(string),
						fmt.Sprintf("%s %02s %s", entry.Month, entry.Day, entry.Time),
						entry.Method,
						entry.Source,
						entry.User,
					}
					activityRows = append(activityRows, entryRow)
				}
				for _, entry := range response.Payload.Whitelisted {
					commentEnd := strings.Index(entry.Comment, " -- ")
					if commentEnd == -1 {
						commentEnd = len(entry.Comment)
					}
					entryRow := []string{
						response.Metadata["hostname"].(string),
						entry.Action,
						strings.Trim(strings.Join(strings.Fields(fmt.Sprint(entry.Ports)), ", "), "[]"),
						strings.Join(entry.Protocols, ", "),
						entry.Source,
						fmt.Sprintf("%q", strings.ReplaceAll(entry.Comment[:commentEnd], "_", " ")),
					}
					whitelistRows = squashOrAppendEntry(whitelistRows, entryRow, 2)
				}
				for _, user := range array.UniqueStrings(append(response.Payload.PassAccess, response.Payload.KeyAccess...)) {
					checkMap := map[bool]string{true: "✔", false: " "}
					entry := []string{
						response.Metadata["hostname"].(string),
						user,
						checkMap[array.ContainsString(response.Payload.PassAccess, user)],
						checkMap[array.ContainsString(response.Payload.KeyAccess, user)],
					}
					accessRows = append(accessRows, entry)
				}
				responseRows = append(responseRows, responseRow)
			}
		}
		server.FilterForEach(filter, runner)
		if output == "json" {
			encoded, _ := json.MarshalIndent(rawResponses, "", "\t")
			fmt.Println(string(encoded))
			return
		}
		if selector != "" && len(responseRows) > 1 {
			fmt.Printf("\nDisplaying results for %q server(s):\n", selector)
		}
		fmt.Println()
		text.HeaderPrint([]string{"Reported System(s)"})
		fmt.Println()
		text.TablePrint(emptyMsg, responseRows, 0)
		fmt.Println()
		if len(responseRows) > 1 {
			sort.SliceStable(activityRows, func(i, j int) bool {
				return activityRows[i][1] < activityRows[j][1]
			})
			whitelistRows = append([][]string{{"Hostname", "Action", "Port(s)", "Protocol(s)", "IPV4/CIDR", "Comment"}}, whitelistRows...)
			activityRows = append([][]string{{"Hostname", "Timestamp", "Method", "IP Address", "User"}}, activityRows...)
			accessRows = append([][]string{{"Hostname", "User", "Password", "SSH Key"}}, accessRows...)
			text.HeaderPrint([]string{
				"SSH Access Activity:",
				"Incoming connections from internal network IPs are not shown",
				"These networks include BACKUP_NETWORK, AUTOMATION_NETWORK, DIRECT_NETWORK, MANAGEMENT_NETWORK.",
			})
			fmt.Println()
			text.TablePrint("No access log entries found.", activityRows, 0)
			fmt.Println()
			text.HeaderPrint([]string{
				"Current Firewall Entries:",
				"Internal network IPs are obfuscated.",
				"These networks include BACKUP_NETWORK, AUTOMATION_NETWORK, DIRECT_NETWORK, MANAGEMENT_NETWORK.",
			})
			fmt.Println()
			text.TablePrint("No firewall entries found.", whitelistRows, 0)
			fmt.Println()
			text.HeaderPrint([]string{"Current SSH User List"})
			fmt.Println()
			text.TablePrint("No user access entries found.", accessRows, 0)
			fmt.Println()
		}
	},
}

func init() {
	reportCmd.AddCommand(reportAuditCmd)
	reportAuditCmd.Flags().SortFlags = true
	reportAuditCmd.Flags().StringP("type", "t", "", "specify server type selector")
	reportAuditCmd.Flags().StringP("output", "o", "table", "specify 'table' or 'json'")
}
