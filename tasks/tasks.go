package tasks

import (
	"mail-telemetry/db"
)

func InitTasks() {
	db.LoadDbMultipleScenariosToSqlite("scenarios")
}
