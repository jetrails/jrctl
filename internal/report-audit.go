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
						entry.Source,
						fmt.Sprintf("%q", entry.Comment[:commentEnd]),
					}
					whitelistRows = append(whitelistRows, entryRow)
				}
				for _, user := range array.UniqueStrings(append(response.Payload.PasswordAccess, response.Payload.SSHAccess...)) {
					checkMap := map[bool]string{true: "✔", false: " "}
					entry := []string{
						response.Metadata["hostname"].(string),
						user,
						checkMap[array.ContainsString(response.Payload.PasswordAccess, user)],
						checkMap[array.ContainsString(response.Payload.SSHAccess, user)],
					}
					accessRows = append(accessRows, entry)
				}
				responseRows = append(responseRows, responseRow)
			}
		}
		server.FilterForEach(filter, runner)
		if output == "json" {
			encoded, _ := json.MarshalIndent(rawResponses, "", "\t")
			// fmt.Printf ( "%#v", rawResponses )
			fmt.Println(string(encoded))
			return
		}
		if selector != "" && len(responseRows) > 1 {
			fmt.Printf("\nDisplaying results for %q server(s):\n", selector)
		}
		fmt.Println()
		text.HeaderPrint("Matching Servers")
		fmt.Println()
		text.TablePrint(emptyMsg, responseRows, 0)
		fmt.Println()
		if len(responseRows) > 1 {
			sort.SliceStable(activityRows, func(i, j int) bool {
				return activityRows[i][1] < activityRows[j][1]
			})
			whitelistRows = append([][]string{{"Hostname", "Action", "Port", "IP Address", "Comment"}}, whitelistRows...)
			activityRows = append([][]string{{"Hostname", "Timestamp", "Method", "IP Address", "User"}}, activityRows...)
			accessRows = append([][]string{{"Hostname", "User", "Password", "SSH"}}, accessRows...)
			text.HeaderPrint("Access Log")
			fmt.Println()
			text.TablePrint("No audit reports to display.", activityRows, 0)
			fmt.Println()
			text.HeaderPrint("Firewall Entries")
			fmt.Println()
			text.TablePrint("No audit reports to display.", whitelistRows, 0)
			fmt.Println()
			text.HeaderPrint("User Access")
			fmt.Println()
			text.TablePrint("No audit reports to display.", accessRows, 0)
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
