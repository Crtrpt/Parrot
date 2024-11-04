package job

import (
	"context"

	"github.com/parrot/internal/host"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

var Cron *cron.Cron

func Start(ctx context.Context) (err error) {
	c := cron.New()
	for _, item := range host.Cfg.Crontab {
		item := item
		logrus.Info("add func", item.Spec)
		c.AddFunc(item.Spec, func() { Jobs[item.Name].Run(context.Background(), item.Argv) })
	}
	c.Start()
	return nil
}

func Stop(ctx context.Context) {
	logrus.Info("stop crontab")
}
