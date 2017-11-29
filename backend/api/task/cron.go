package task

import (
	"github.com/robfig/cron"
	"github.com/siddontang/go/log"
)

const statisticsSpec = "0 30 0 * * *" //everyday at 0:30

func RunCrons() {
	c := cron.New()
	c.AddFunc(statisticsSpec, func() {
		GetDaily()
		GetTotal()
	})
	c.Start()
	for _, v := range c.Entries() {
		log.Info("Job next run time:")
		log.Info(v.Next)
		log.Info("Job Schedule:")
		log.Info(v.Schedule)
	}
}
