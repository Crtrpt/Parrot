package job

import (
	"context"

	"github.com/sirupsen/logrus"
)

func init() {
	j := &CheckJob{
		Name: "check",
	}
	Jobs[j.GetName()] = j
}

type CheckJob struct {
	Job
	Name string
}

func (j CheckJob) GetName() string {
	return j.Name
}

func (j CheckJob) Run(ctx context.Context, payload any) error {
	logrus.Info("TODO job")
	return nil
}
