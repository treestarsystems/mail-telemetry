package tasks

import (
	"log"
	"mail-telemetry/db"
)

func InitTasks() {
	db.LoadDbMultipleScenariosToSqlite("scenarios")
}

func TelemetryScenario_of365(scenarioName string) {
	log.Printf("Task - Initiating Scenario: %s\n", scenarioName)
	log.Printf("Task - Scenario Complete: %s\n", scenarioName)
}
