/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"os"
	"text/tabwriter"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	elb "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

// elbsCmd represents the elbs command
var elbsCmd = &cobra.Command{
	Use:   "elbs",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		ListElbs()
	},
}

func init() {
	rootCmd.AddCommand(elbsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// elbsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// elbsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func ListElbs() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 1, ' ', 0)
	colors := []Color{DefaultText, DefaultText, DefaultText, DefaultText, DefaultText}

	svc2 := elbv2.NewFromConfig(cfg)
	params2 := &elbv2.DescribeLoadBalancersInput{}
	p2 := elbv2.NewDescribeLoadBalancersPaginator(svc2, params2)

	for p2.HasMorePages() {
		output2, err := p2.NextPage(context.TODO())
		if err != nil {
			// handle error

			log.WithFields(log.Fields{
				"err": err,
			}).Fatal("Can't list elbv2")
		}
		for _, lbs2 := range output2.LoadBalancers {
			PrintRow(w, PaintRow(colors, []string{
				aws.ToString(lbs2.LoadBalancerName),
				string(lbs2.Scheme),
				aws.ToString(lbs2.VpcId),
				string(lbs2.Type),
				timeConvert(lbs2.CreatedTime).Format(timeFormat),
			}),
			)
		}
	}

	svc1 := elb.NewFromConfig(cfg)
	params1 := &elb.DescribeLoadBalancersInput{}
	p1 := elb.NewDescribeLoadBalancersPaginator(svc1, params1)

	for p1.HasMorePages() {
		output1, err := p1.NextPage(context.TODO())
		if err != nil {
			// handle error
			log.WithFields(log.Fields{
				"err": err,
			}).Fatal("Can't list clb")
		}
		for _, lbs1 := range output1.LoadBalancerDescriptions {
			PrintRow(w, PaintRow(colors, []string{
				aws.ToString(lbs1.LoadBalancerName),
				aws.ToString(lbs1.Scheme),
				aws.ToString(lbs1.VPCId),
				"classic",
				timeConvert(lbs1.CreatedTime).Format(timeFormat),
			}),
			)
		}
	}

	w.Flush()
}
