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

	//"github.com/fatih/color"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/spf13/cobra"
)

// dbsCmd represents the dbs command
var dbsCmd = &cobra.Command{
	Use:   "dbs",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ListDBs()
	},
}

func init() {
	rootCmd.AddCommand(dbsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ListDBs() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := rds.NewFromConfig(cfg)

	params := &rds.DescribeDBInstancesInput{}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 1, ' ', 0)
	//yellow := color.New(color.FgYellow).SprintFunc()

	paginator := rds.NewDescribeDBInstancesPaginator(svc, params)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			// handle error
		}
		colors := []Color{DefaultText, DefaultText, DefaultText}
		for _, dbs := range output.DBInstances {
			PrintRow(w, PaintRow(colors, []string{
				func() string {
					return aws.ToString(dbs.DBInstanceIdentifier)
				}(),
				func() string {
					if aws.ToString(dbs.DBInstanceStatus) == "available" {
						return timeConvert(dbs.InstanceCreateTime).Format(timeFormat)
					} else {
						return aws.ToString(dbs.DBInstanceStatus)
					}

				}(),
				func() string {
					if aws.ToString(dbs.DBInstanceStatus) == "available" {
						return aws.ToString(dbs.Endpoint.Address)
					} else {
						return "nodata"
					}
				}(),
			}),
			)
		}
	}
	w.Flush()
}
