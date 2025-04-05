package cron

import (
	"mail-telemetry/utils"

	"github.com/robfig/cron/v3"
)

func InitCron() {
	c := cron.New()
	// At 0, 15, 30, and 45 minute mark, check for changes to config and scenarios file then load changes to db.
	c.AddFunc("0,15,30,45 * * * *", func() {
		utils.ParseScenariosCSV("./scenarios.csv")
	})
	c.Start()
}
