/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/parrot/internal/core"
	"github.com/parrot/internal/job"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// jobCmd represents the job command
var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "A brief description of your command",
	Long:  `手动运行指定job`,
	Run: func(cmd *cobra.Command, args []string) {
		j, err := cmd.Flags().GetString("job")
		if err != nil {
			panic(err)
		}
		ctx := context.Background()

		if i, ok := job.Jobs[j]; ok {
			err := core.Start(ctx)
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
}
