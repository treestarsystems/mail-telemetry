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

	// TODO: write all scenario instances to it's own table with a MD5 hash of the struct.
	// GenerateMD5HashOfStruct generates an MD5 hash of a given struct.
	// func GenerateMD5HashOfStruct(data interface{}) (string, error) {
	// 	// Convert the struct to JSON
	// 	jsonData, err := json.Marshal(data)
	// 	if err != nil {
	// 		return "", fmt.Errorf("error - Failed to marshal struct to JSON: %v", err)
	// 	}

	//		// Generate MD5 hash
	//		hash := md5.Sum(jsonData)
	//		return hex.EncodeToString(hash[:]), nil
	//	}
}
