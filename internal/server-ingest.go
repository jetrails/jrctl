package internal

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jetrails/jrctl/pkg/array"
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
		"echo -n TOKEN | jrctl server ingest -t jump -e 10.10.10.7",
		"echo -n TOKEN | jrctl server ingest -t web -e 10.10.10.6 -f",
	}),
	Args: func(cmd *cobra.Command, args []string) error {
		types, _ := cmd.Flags().GetStringSlice("type")
		if len(types) == 0 {
			return errors.New(fmt.Sprintf("%q expects at least one value", "type"))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		force, _ := cmd.Flags().GetBool("force")
		endpoint, _ := cmd.Flags().GetString("endpoint")
		var tokenValue string = ""
		if stat, _ := os.Stdin.Stat(); stat.Mode()&os.ModeCharDevice == 0 {
			if bytes, error := ioutil.ReadAll(os.Stdin); error == nil {
				tokenValue = strings.TrimSpace(string(bytes))
			} else {
				if !quiet {
					fmt.Printf("\nFailed to read piped data.\n\n")
				}
				os.Exit(1)
			}
		} else {
			if !quiet {
				fmt.Printf("\nMust pipe token to STDIN, see help for example.\n\n")
			}
			os.Exit(1)
		}
		types, _ := cmd.Flags().GetStringSlice("type")
		emptyMsg := fmt.Sprintf("No configured server with type(s) \"%s\" found.", strings.Join(types, "\", \""))
		rows := [][]string{[]string{"Server", "Type(s)", "Action"}}
		savedServers := []server.Entry{}
		if error := viper.UnmarshalKey("servers", &savedServers); error != nil {
			if !quiet {
				fmt.Printf("\nFailed to parse current config file.\n\n")
			}
			os.Exit(1)
		}
		for i, savedServer := range savedServers {
			if array.HasValidStringValues(types, savedServer.Types) {
				savedServers[i].Token = tokenValue
				row := []string{
					strings.TrimSuffix(savedServer.Endpoint, ":27482"),
					strings.Join(savedServer.Types, ", "),
					"Updated",
				}
				rows = append(rows, row)
			}
		}
		if force {
			createdEntry := server.Entry{
				Endpoint: endpoint,
				Token:    tokenValue,
				Types:    types,
			}
			savedServers = append(savedServers, createdEntry)
			row := []string{
				strings.TrimSuffix(createdEntry.Endpoint, ":27482"),
				strings.Join(createdEntry.Types, ", "),
				"Created",
			}
			rows = append(rows, row)
		}
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
								if !quiet {
									fmt.Printf("\nIngested token for server(s) with type(s) \"%s\".\n", strings.Join(types, "\", \""))
								}
							} else {
								if !quiet {
									fmt.Printf("\nFailed to write config to file.\n\n")
								}
								os.Exit(1)
							}
						} else {
							if !quiet {
								fmt.Printf("\nFailed to marshal to yaml.\n\n")
							}
							os.Exit(1)
						}
					}
				} else {
					if !quiet {
						fmt.Printf("\nFailed to parse config file.\n\n")
					}
					os.Exit(1)
				}
			} else {
				if !quiet {
					fmt.Printf("\nFailed to open config file.\n\n")
				}
				os.Exit(1)
			}
		}
		if !quiet {
			text.TablePrint(emptyMsg, rows, 1)
		}
	},
}

func init() {
	serverCmd.AddCommand(serverIngestCmd)
	serverIngestCmd.Flags().SortFlags = true
	serverIngestCmd.Flags().BoolP("quiet", "q", false, "output as little information as possible")
	serverIngestCmd.Flags().BoolP("force", "f", false, "create new entry even if matching entries exist")
	serverIngestCmd.Flags().StringP("endpoint", "e", "127.0.0.1:27482", "server endpoint used for new entries only")
	serverIngestCmd.Flags().StringSliceP("type", "t", []string{}, "type selector(s), all must be present to match entry")
	serverIngestCmd.MarkFlagRequired("type")
}
