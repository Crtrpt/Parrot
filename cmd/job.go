/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/parrot/internal/host"
	"github.com/parrot/internal/job"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// jobCmd represents the job command
var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		j, err := cmd.Flags().GetString("job")
		if err != nil {
			panic(err)
		}
		ctx := context.Background()

		if i, ok := job.Jobs[j]; ok {
			err := host.Start(ctx)
			if err != nil {
				panic(err)
			}
			err = i.Run(ctx, args)
			if err != nil {
				panic(err)
			}
		}
		for _, v := range job.Jobs {
			logrus.Infof("job:%s", v.GetName())
		}
	},
}

func init() {
	rootCmd.AddCommand(jobCmd)

	jobCmd.PersistentFlags().StringP("job", "j", "", "A help for foo")
	jobCmd.PersistentFlags().StringP("payload", "p", "", "payload argv")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
