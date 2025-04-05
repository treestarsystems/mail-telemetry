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
		log.Printf("info - SQLite: Database file does not exist, recreating...\n")
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
	scenarios := utils.ParseScenariosCSV("./scenarios.csv")
	for _, scescenario := range scenarios {
		LoadDbSingleScenarioToSqlite(scescenario, tableName)
	}
}
