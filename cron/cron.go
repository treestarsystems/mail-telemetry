package cron

import (
	"mail-telemetry/tasks"

	"github.com/robfig/cron/v3"
)

func InitCron() {
	c := cron.New()
	c.AddFunc("* * * * *", func() {
		tasks.TelemetryScenario_of365("Test")
	})
	c.Start()
}
