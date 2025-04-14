package cron

import (
	"log"
	"mail-telemetry/db"

	"github.com/robfig/cron/v3"
)

func InitCron() {
	c := cron.New()
	// At 0, 15, 30, and 45 minute mark, check for changes to config and scenarios file then load changes to db.
	log.Println("Cron: Scheduling scenarios job at '0,15,30,45 * * * *'")
	c.AddFunc("0,15,30,45 * * * *", func() {
		db.LoadDbMultipleScenariosToSqlite("scenarios")
	})
	c.Start()
}
