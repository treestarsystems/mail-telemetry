package tasks

import (
	"mail-telemetry/db"
)

func InitTasks() {
	db.LoadDbMultipleScenariosToSqlite("scenarios")

	// scenarios, _ := db.RetrieveScenarioFromSqliteAll("scenarios")
	// scenarios, _ := db.RetrieveScenarioFromSqliteByColumnName("scenarios", "type", "O365")
}
