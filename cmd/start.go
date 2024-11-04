/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/oklog/run"
	"github.com/parrot/internal/api"
	"github.com/parrot/internal/host"
	"github.com/parrot/internal/job"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long:  `启动服务`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		err := host.Start(ctx)
		if err != nil {
			panic(err)
		}

		var g run.Group

		g.Add(func() error {
			return api.Start(ctx)
		}, func(err error) {
			api.Stop(ctx)
		})

		g.Add(func() error {
			return job.Start(ctx)
		}, func(err error) {
			job.Stop(ctx)
		})

		g.Run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("service", "api", "api,crontab")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
