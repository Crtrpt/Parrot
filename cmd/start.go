/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/oklog/run"
	"github.com/parrot/internal/api"
	"github.com/parrot/internal/core"
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
		err := core.Start(ctx)
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
}
