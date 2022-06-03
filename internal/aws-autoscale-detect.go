package internal

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	. "github.com/jetrails/jrctl/pkg/output"
	"github.com/jetrails/jrctl/pkg/text"
	"github.com/spf13/cobra"
)

var (
	ErrAwsImdsCredsMissing         = errors.New("failed to extract credentials from IMDS service")
	ErrAwsAutoScalingGroupNotFound = errors.New("failed to find autoscaling group with that name")
	ErrAwsInstanceDetails          = errors.New("failed to query instance details")
)

func GetAWSConfig() (aws.Config, error) {
	return config.LoadDefaultConfig(
		context.TODO(),
		config.WithEC2IMDSRegion(),
		config.WithEC2RoleCredentialOptions(func(opts *ec2rolecreds.Options) {
			opts.Client = imds.New(imds.Options{})
		}),
	)
}

func GetInstanceIdentityDocument(cfg aws.Config) *imds.InstanceIdentityDocument {
	client := imds.NewFromConfig(cfg)
	document, err := client.GetInstanceIdentityDocument(context.TODO(), &imds.GetInstanceIdentityDocumentInput{})
	if err == nil {
		return &document.InstanceIdentityDocument
	}
	return nil
}

func GetAutoScalingGroupInstances(cfg aws.Config, asgName string) ([]types.Instance, error) {
	entries := []types.Instance{}
	autoscalingClient := autoscaling.NewFromConfig(cfg)
	ec2Client := ec2.NewFromConfig(cfg)

	asgs, err := autoscalingClient.DescribeAutoScalingGroups(
		context.TODO(),
		&autoscaling.DescribeAutoScalingGroupsInput{AutoScalingGroupNames: []string{asgName}},
	)

	if err != nil {
		return entries, ErrAwsAutoScalingGroupNotFound
	}

	for _, asg := range asgs.AutoScalingGroups {
		instanceIds := []string{}
		for _, instance := range asg.Instances {
			instanceIds = append(instanceIds, aws.ToString(instance.InstanceId))
		}
		infos, err := ec2Client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
			InstanceIds: instanceIds,
			Filters: []types.Filter{
				{
					Name:   aws.String("instance-state-name"),
					Values: []string{"running"},
				},
			},
		})
		if err != nil {
			return []types.Instance{}, ErrAwsInstanceDetails
		}
		for _, reservation := range infos.Reservations {
			for _, instance := range reservation.Instances {
				entries = append(entries, instance)
			}
		}
	}

	return entries, nil
}

var awsAutoscaleDetectCmd = &cobra.Command{
	Use:   "autoscale-detect AUTOSCALING_GROUP_NAME",
	Short: "Query aws for instances in autoscaling group",
	Example: text.Examples([]string{
		"jrctl aws autoscale-detect example-asg",
		"jrctl aws autoscale-detect example-asg -q",
	}),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		asgName := args[0]

		output := NewOutput(quiet, []string{})
		output.DisplayTags = false
		output.DisplayServers = false
		output.FailOnNoResults = true
		output.ExitCodeNoResults = 2
		output.ErrMsgNoResults = Lines{"could not find autoscaling group with that name"}

		tbl := output.CreateTable(Columns{
			"Instance ID",
			"Image ID",
			"State",
			"Private IP",
			"Public IP",
			"Launch Time",
		})

		cfg, err := GetAWSConfig()
		if err != nil {
			output.ExitWithMessage(3, "\n"+ErrAwsImdsCredsMissing.Error()+"\n")
		}

		if instances, err := GetAutoScalingGroupInstances(cfg, asgName); err == nil {
			for _, instance := range instances {
				tbl.AddQuietEntry(aws.ToString(instance.PrivateIpAddress))
				tbl.AddRow(Columns{
					aws.ToString(instance.InstanceId),
					aws.ToString(instance.ImageId),
					"Running",
					aws.ToString(instance.PrivateIpAddress),
					aws.ToString(instance.PublicIpAddress),
					instance.LaunchTime.String(),
				})
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

		output.Print()
	},
}

func init() {
	OnlyRunOnAWS(awsAutoscaleDetectCmd)
	awsCmd.AddCommand(awsAutoscaleDetectCmd)
	awsAutoscaleDetectCmd.Flags().SortFlags = true
	awsAutoscaleDetectCmd.Flags().BoolP("quiet", "q", false, "display only private ip address")
}
