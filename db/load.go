package db

import (
	"errors"
	"fmt"
	"log"
	"mail-telemetry/utils"
	"os"
)

func LoadDbSingleScenarioToSqlite(scenario utils.Scenario, tableName string) {
	// TODO: Need a way to get the correct file path no matter the OS.
	// This will rerun the connection to the database if the file does not exist.
	fileName := fmt.Sprintf("./%v", os.Getenv("DB_SQLITE_FILENAME"))
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		log.Println("info - SQLite: Database file does not exist, recreating")
		LoadDbConnectToSqlite()
	}

	// Save = Upsert
	DB.Table(tableName).Where(utils.Scenario{Name: scenario.Name}).Assign(utils.Scenario{
		Type:               scenario.Type,
		CredentialLocation: scenario.CredentialLocation,
		FromEmail:          scenario.FromEmail,
		ToEmail:            scenario.ToEmail,
		Description:        scenario.Description,
	}).FirstOrCreate(&utils.Scenario{
		Type:               scenario.Type,
		CredentialLocation: scenario.CredentialLocation,
		FromEmail:          scenario.FromEmail,
		ToEmail:            scenario.ToEmail,
		Description:        scenario.Description,
	})
}

func LoadDbMultipleScenariosToSqlite(tableName string) {
	log.Println("- Loading scenarios")
	scenarios, err := utils.ParseScenariosCSV("./scenarios.csv")
	if err != nil {
		log.Print(err)
		return
	}
	if len(scenarios) == 0 {
		log.Println("-- There are no scenarios to process")
		return
	}
	for i, scescenario := range scenarios {
		log.Printf("-- Scenario %v loaded to database\n", i+1)
		LoadDbSingleScenarioToSqlite(scescenario, tableName)
	}
	log.Println("-- Loading scenarios, complete")
}
