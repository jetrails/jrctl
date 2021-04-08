package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var serverIngestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "Ingest server token and save it to config",
	Long: text.Combine([]string{
		text.Paragraph([]string{
			"Ingest server token and save it to config.",
		}),
	}),
	Example: text.Examples([]string{
		"echo -n TOKEN | jrctl server ingest -t localhost",
	}),
	Run: func(cmd *cobra.Command, args []string) {
		var tokenValue string = ""
		if stat, _ := os.Stdin.Stat(); stat.Mode()&os.ModeCharDevice == 0 {
			if bytes, error := ioutil.ReadAll(os.Stdin); error == nil {
				tokenValue = strings.TrimSpace(string(bytes))
			} else {
				fmt.Printf("\nFailed to read piped data.\n\n")
				return
			}
		} else {
			fmt.Printf("\nMust pipe token to STDIN, see help for example.\n\n")
			return
		}
		selector, _ := cmd.Flags().GetString("type")
		filter := []string{selector}
		emptyMsg := fmt.Sprintf("No configured %q server(s) found.", selector)
		rows := [][]string{[]string{"Server", "Type(s)", "Action"}}
		savedServers := []server.Entry{}
		if error := viper.UnmarshalKey("servers", &savedServers); error != nil {
			fmt.Printf("\nFailed to parse current config file.\n\n")
			return
		}
		runner := func(index, total int, context server.Context) {
			for i, savedServer := range savedServers {
				sort.Strings(context.Types)
				sort.Strings(savedServer.Types)
				if true &&
					savedServer.Endpoint == context.Endpoint &&
					savedServer.Token == context.Token &&
					strings.Join(savedServer.Types, "|") == strings.Join(context.Types, "|") {
					savedServers[i].Token = tokenValue
					row := []string{
						strings.TrimSuffix(context.Endpoint, ":27482"),
						strings.Join(context.Types, ", "),
						"Ingested",
					}
					rows = append(rows, row)
				}
			}
		}
		server.FilterForEach(filter, runner)
		if len(rows) > 1 {
			if data, error := ioutil.ReadFile(viper.ConfigFileUsed()); error == nil {
				var c interface{}
				if error = yaml.Unmarshal([]byte(data), &c); error == nil {
					if _, ok := c.(map[string]interface{})["servers"].([]interface{}); ok {
						c.(map[string]interface{})["servers"] = savedServers
						if d, error := yaml.Marshal(&c); error == nil {
							if file, error := os.OpenFile(viper.ConfigFileUsed(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600); error == nil {
								defer file.Close()
								file.Write(d)
								fmt.Printf("\nIngested token for %q server(s):\n", selector)
							} else {
								fmt.Printf("\nFailed to write config to file.\n\n")
								return
							}
						} else {
							fmt.Printf("\nFailed to marshal to yaml.\n\n")
							return
						}
					}
				} else {
					fmt.Printf("\nFailed to parse config file.\n\n")
					return
				}
			} else {
				fmt.Printf("\nFailed to open config file.\n\n")
				return
			}
		}
		text.TablePrint(emptyMsg, rows, 1)
	},
}

func init() {
	serverCmd.AddCommand(serverIngestCmd)
	serverIngestCmd.Flags().SortFlags = true
	serverIngestCmd.Flags().StringP("type", "t", "", "specify server type selector")
	serverIngestCmd.MarkFlagRequired("type")
}
