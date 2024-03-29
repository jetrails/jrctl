package internal

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jetrails/jrctl/pkg/array"
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/jetrails/jrctl/sdk/report"
	"github.com/spf13/cobra"
)

var reportAuditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Month-to-date security audit to ensure access is properly limited",
	Example: text.Examples([]string{
		"jrctl report audit",
		"jrctl report audit -t www",
		"jrctl report audit -o json",
	}),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("format")
		validOutput := []string{"table", "json"}
		if !array.ContainsString(validOutput, format) {
			return fmt.Errorf("format must be one of %v", validOutput)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("format")
		tags, _ := cmd.Flags().GetStringArray("tag")

		responses := []report.AuditResponse{}
		output := NewOutput(false, tags)
		output.DisplayServers = true
		output.FailOnNoServers = true
		output.FailOnNoResults = true
		output.ExitCodeNoServers = 1
		output.ExitCodeNoResults = 2

		tblWhitelist := NewTable(Columns{
			"Hostname",
			"Action",
			"Port(s)",
			"Protocol(s)",
			"IPV4/CIDR",
			"Comment",
		})
		tblActivity := NewTable(Columns{
			"Hostname",
			"Timestamp",
			"Method",
			"IP Address",
			"User",
		})
		tblAccess := NewTable(Columns{
			"Hostname",
			"User",
			"Password",
			"SSH Key",
		})
		tblDatabases := NewTable(Columns{
			"Hostname",
			"Database",
			"User(s)",
		})
		tblDatabaseUsers := NewTable(Columns{
			"Hostname",
			"User",
			"Database(s)",
		})

		for _, context := range config.GetContexts(tags) {
			response := report.Audit(context)
			responses = append(responses, response)
			output.AddServer(
				context,
				response.GetGeneric(),
				response.Status,
			)
			if response.IsOkay() {
				for _, entry := range response.Payload.Activity {
					tblActivity.AddRow(Columns{
						response.Metadata["hostname"],
						fmt.Sprintf("%s %02s %s", entry.Month, entry.Day, entry.Time),
						entry.Method,
						entry.Source,
						entry.User,
					})
				}
				for _, entry := range response.Payload.Whitelisted {
					tblWhitelist.AddRow(Columns{
						response.Metadata["hostname"],
						entry.Action,
						strings.Trim(strings.Join(strings.Fields(fmt.Sprint(entry.Ports)), ", "), "[]"),
						strings.Join(entry.Protocols, ", "),
						entry.Source,
						fmt.Sprintf("%q", extractFirewallComment(entry.Comment)),
					})
				}
				for _, user := range array.UniqueStrings(append(response.Payload.PassAccess, response.Payload.KeyAccess...)) {
					checkMap := map[bool]string{true: "✔", false: " "}
					tblAccess.AddRow(Columns{
						response.Metadata["hostname"],
						user,
						checkMap[array.ContainsString(response.Payload.PassAccess, user)],
						checkMap[array.ContainsString(response.Payload.KeyAccess, user)],
					})
				}
				for _, db := range response.Payload.Databases {
					users := ""
					del := ""
					for _, user := range db.Users {
						users += del + user.Name + "@" + user.From
						del = ", "
					}
					tblDatabases.AddRow(Columns{
						response.Metadata["hostname"],
						db.Name,
						users,
					})
				}
				for _, user := range response.Payload.DatabaseUsers {
					dbs := ""
					del := ""
					for _, db := range user.Databases {
						dbs += del + db.Name
						del = ", "
					}
					tblDatabaseUsers.AddRow(Columns{
						response.Metadata["hostname"],
						user.Name + "@" + user.From,
						dbs,
					})
				}
			}
		}

		if format == "json" {
			encoded, _ := json.MarshalIndent(responses, "", "\t")
			fmt.Println(string(encoded))
			return
		}

		output.Servers.Title = Lines{
			"Reported System(s)",
		}
		tblActivity.Title = Lines{
			"SSH Access Activity:",
			"Incoming connections from internal network IPs are not shown",
			"These networks include BACKUP_NETWORK, AUTOMATION_NETWORK, DIRECT_NETWORK, MANAGEMENT_NETWORK.",
		}
		tblWhitelist.Title = Lines{
			"Current Firewall Entries:",
			"Internal network IPs are obfuscated.",
			"These networks include BACKUP_NETWORK, AUTOMATION_NETWORK, DIRECT_NETWORK, MANAGEMENT_NETWORK.",
		}
		tblAccess.Title = Lines{
			"Current SSH User List",
		}
		tblDatabases.Title = Lines{
			"Current Databases",
		}
		tblDatabaseUsers.Title = Lines{
			"Current Database Users",
		}

		tblWhitelist.SquashOnPivot(2)
		tblActivity.Sort(1)

		output.AddTable(tblActivity)
		output.AddTable(tblWhitelist)
		output.AddTable(tblAccess)
		output.AddTable(tblDatabases)
		output.AddTable(tblDatabaseUsers)

		tblActivity.EmptyMessage = Lines{"No entries found"}
		tblWhitelist.EmptyMessage = Lines{"No entries found"}
		tblAccess.EmptyMessage = Lines{"No entries found"}
		tblDatabases.EmptyMessage = Lines{"No entries found"}
		tblDatabaseUsers.EmptyMessage = Lines{"No entries found"}
		output.ErrMsgNoResults = Lines{}

		output.Print()
	},
}

func init() {
	reportCmd.AddCommand(reportAuditCmd)
	reportAuditCmd.Flags().SortFlags = true
	reportAuditCmd.Flags().StringP("format", "f", "table", "specify 'table' or 'json'")
	reportAuditCmd.Flags().StringArrayP("tag", "t", []string{}, "filter nodes using tags")
}
