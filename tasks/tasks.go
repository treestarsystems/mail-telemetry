package tasks

import (
	"mail-telemetry/db"
)

func InitTasks() {
	db.LoadDbMultipleScenariosToSqlite("scenarios")

	// scenarios, _ := db.RetrieveScenarioFromSqliteAll("scenarios")
	// scenarios, _ := db.RetrieveScenarioFromSqliteByColumnName("scenarios", "type", "O365")

	// for _, scenario := range scenarios {
	// 	_, err := email.GenerateScenarioDetails(scenario)
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// }
}
