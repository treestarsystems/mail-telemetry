package tasks

import (
	"fmt"
	"mail-telemetry/db"
	"mail-telemetry/email"
	"mail-telemetry/utils"
)

func InitTasks() {
	db.LoadDbMultipleScenariosToSqlite("scenarios")

	scenarios, _ := db.RetrieveScenarioFromSqliteAll("scenarios")
	// scenarios, _ := db.RetrieveScenarioFromSqliteByColumnName("scenarios", "type", "O365")
	scenarioCount := 0
	for _, scenario := range scenarios {
		// utils.PrintStructAsPrettyJSON(scenario)
		// email.GenerateScenarioInstance(&scenario)
		scenarioInstances := email.GenerateScenarioInstance(&scenario)
		for _, instance := range scenarioInstances {
			utils.PrintStructAsPrettyJSON(instance)
			scenarioCount++
		}
	}
	fmt.Println(scenarioCount)
}
