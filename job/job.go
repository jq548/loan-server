package job

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"loan-server/config"
)

type Job struct {
	MyConfig *config.Config
}

func (job *Job) StartJob(spec string, fun func()) {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc(spec, fun)
	if err != nil {
		zap.L().Error("Start Job failed")
	} else {
		c.Start()
	}
}
