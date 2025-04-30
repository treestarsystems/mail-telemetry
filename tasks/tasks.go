package tasks

import (
	"mail-telemetry/db"
	"mail-telemetry/utils"
)

func InitTasks() {
	db.LoadDbMultipleScenariosToSqlite("scenarios")

	// scenarios, _ := db.RetrieveScenarioFromSqliteAll("scenarios")
	scenarios, _ := db.RetrieveScenarioFromSqliteByColumnName("scenarios", "type", "O365")

	for _, scenario := range scenarios {
		utils.PrintStructAsPrettyJSON(scenario)
		// _, err := email.GenerateScenarioDetails(&scenario)
		// if err != nil {
		// 	log.Print(err)
		// }
	}
}
