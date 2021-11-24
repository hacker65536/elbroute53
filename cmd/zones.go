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
	"fmt"
	"os"
	"regexp"
	"text/tabwriter"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

// zonesCmd represents the zones command
var zonesCmd = &cobra.Command{
	Use:   "zones",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		reg := regexp.MustCompile(`/hostedzone/(.*)`)
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Fatal("unable to load SDK config")
		}

		svc := route53.NewFromConfig(cfg)

		params := &route53.ListHostedZonesInput{
			MaxItems: aws.Int32(20),
		}
		resp, err := svc.ListHostedZones(context.TODO(), params)
		if err != nil {
			//log.Fatalf("failed to list tables, %v", err)
			log.WithFields(log.Fields{
				"err": err,
			}).Fatal("failed to list hosted zones")
		}

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 0, 1, ' ', 0)
		for _, zone := range resp.HostedZones {
			fmt.Fprintf(w, "%s\t%s\t%d\t%s\n", aws.ToString(zone.Name),
				func() string {
					if !zone.Config.PrivateZone {
						return "public"
					}
					return "private"
				}(),
				aws.ToInt64(zone.ResourceRecordSetCount),
				reg.ReplaceAll(([]byte(aws.ToString(zone.Id))), []byte("$1")))
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(zonesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// zonesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// zonesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
