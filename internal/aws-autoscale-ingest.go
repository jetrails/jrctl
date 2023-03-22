package internal

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/jetrails/jrctl/pkg/array"
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/jetrails/jrctl/sdk/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NormalizeEndpoint(endpoint string) string {
	if strings.Contains(endpoint, ":") {
		return endpoint
	}
	return fmt.Sprintf("%s:27482", endpoint)
}

func NormalizeEndpoints(endpoints []string) []string {
	for i, endpoint := range endpoints {
		endpoints[i] = NormalizeEndpoint(endpoint)
	}
	return endpoints
}

var awsAutoscaleIngestCmd = &cobra.Command{
	Use:   "autoscale-ingest AUTOSCALING_GROUP_NAME",
	Short: "Display databases in deployment",
	Example: text.Examples([]string{
		"jrctl aws autoscale-ingest example-asg",
		"jrctl aws autoscale-ingest example-asg -t www",
		"jrctl aws autoscale-ingest example-asg -q",
	}),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		tag, _ := cmd.Flags().GetString("tag")
		tags := []string{tag}
		asgName := args[0]

		output := NewOutput(quiet, tags)
		privateIps := []string{}

		cfg, err := GetAWSConfig()
		if err != nil {
			output.ExitWithMessage(3, "\n"+ErrAwsImdsCredsMissing.Error()+"\n")
		}
		if instances, err := GetAutoScalingGroupInstances(cfg, asgName); err == nil {
			for _, instance := range instances {
				privateIps = append(privateIps, aws.ToString(instance.PrivateIpAddress))
			}
		} else {
			switch err {
			case ErrAwsAutoScalingGroupNotFound:
				output.ExitWithMessage(4, "\n"+err.Error()+"\n")
			case ErrAwsInstanceDetails:
				output.ExitWithMessage(5, "\n"+err.Error()+"\n")
			default:
				output.ExitWithMessage(6, "\nunknown error\n")
			}
		}
		endpoints := NormalizeEndpoints(privateIps)
		if len(endpoints) < 1 {
			output.ExitWithMessage(7, "\nmust pass at least one endpoint\n")
		}

		contexts := config.GetContexts(tags)
		nodes := []config.Entry{}
		viper.UnmarshalKey("nodes", &nodes)

		if len(contexts) < 1 {
			output.PrintTags()
			output.ExitWithMessage(8, "\nno nodes found with given tags\n")
		}

		if !config.ContextsHaveSameToken(contexts) {
			output.PrintTags()
			output.ExitWithMessage(9, "\nfound differing tokens, autoscale requires same tokens\n")
		}

		tbl := output.CreateTable(Columns{
			"Endpoint",
			"Action",
			"Tag(s)",
		})

		for _, context := range contexts {
			var action string
			if array.ContainsString(endpoints, context.Endpoint) {
				action = "Skipped"
			} else {
				action = "Deleted"
				filtered := []config.Entry{}
				for _, s := range nodes {
					if s.Endpoint != context.Endpoint {
						filtered = append(filtered, s)
					}
				}
				nodes = filtered
			}
			tbl.AddRow(Columns{
				context.Endpoint,
				action,
				strings.Join(tags, ", "),
			})
		}

		for _, endpoint := range endpoints {
			if !config.ContextsHaveSomeEndpoint(contexts, []string{endpoint}) {
				tbl.AddRow(Columns{
					endpoint,
					"Created",
					strings.Join(tags, ", "),
				})
				entry := config.Entry{
					Endpoint: endpoint,
					Token:    contexts[0].Token,
					Tags:    tags,
				}
				nodes = append(nodes, entry)
			}
		}

		viper.Set("nodes", nodes)
		viper.WriteConfig()

		output.Print()
	},
}

func init() {
	OnlyRunOnAWS(awsAutoscaleIngestCmd)
	awsCmd.AddCommand(awsAutoscaleIngestCmd)
	awsAutoscaleIngestCmd.Flags().SortFlags = true
	awsAutoscaleIngestCmd.Flags().BoolP("quiet", "q", false, "display no output")
	awsAutoscaleIngestCmd.Flags().StringP("tag", "t", "www", "filter nodes using tag")
}
