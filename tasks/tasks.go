package tasks

import (
	"log"
	"mail-telemetry/db"
	"mail-telemetry/email"
)

func InitTasks() {
	email.InitializeEnvValuesOF365()
	db.LoadDbMultipleScenariosToSqlite()

	scenarios, err := db.RetrieveScenarioFromSqliteAll("scenarios")
	if err != nil {
		log.Printf("error - Tasks: failed to retrieve scenarios for queue generation: %v", err)
		return
	}

	queueCount := 0
	for _, scenario := range scenarios {
		scenarioInstances := email.GenerateScenarioInstance(&scenario)
		if err := db.LoadDbScenarioQueueToSqlite(scenarioInstances); err != nil {
			log.Printf("error - Tasks: failed to write scenario queue entries for scenario '%s': %v", scenario.Name, err)
			continue
		}
		queueCount += len(scenarioInstances)
	}

	log.Printf("info - Tasks: queued %d scenario instance(s) into scenarioQueue", queueCount)

	// ----- Test code
	// // scenarios, _ := db.RetrieveScenarioFromSqliteByColumnName("scenarios", "type", "O365")
	// scenarios, _ := db.RetrieveScenarioFromSqliteAll("scenarios")
	// scenarioCount := 0
	// for _, scenario := range scenarios {
	// 	// utils.PrintStructAsPrettyJSON(scenario)
	// 	// email.GenerateScenarioInstance(&scenario)
	// 	scenarioInstances := email.GenerateScenarioInstance(&scenario)
	// 	for _, instance := range scenarioInstances {
	// 		utils.PrintStructAsPrettyJSON(instance)
	// 		scenarioCount++
	// 	}
	// }
	// fmt.Println(scenarioCount)

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

	// fmt.Println(email.GenerateScenarioSubjectString(utils.RandomAplhaNumericString(20)))

	// testEmails := []string{"asdf@adsfsd.com"}
	// fmt.Println(email.ConvertEmailToPlusAddress(testEmails, utils.AppName, utils.RandomAplhaNumericString(20)))

	// fmt.Println(utils.AppName)
	// fmt.Println(utils.SystemHostName)
	// fmt.Println(utils.SystemLocalIpAddress)
}
