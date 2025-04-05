package tasks

import (
	"log"
)

func InitTasks() {
	// TelemetryScenario_of365("Test")
}

func TelemetryScenario_of365(scenarioName string) {
	log.Printf("Task - Initiating Scenario: %s\n", scenarioName)
	log.Printf("Task - Scenario Complete: %s\n", scenarioName)
}
