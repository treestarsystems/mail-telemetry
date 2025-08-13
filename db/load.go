package db

import (
	"errors"
	"fmt"
	"log"
	"mail-telemetry/utils"
	"os"
	"time"

	"gorm.io/gorm"
)

func LoadDbVerifySqliteExists() {
	// This will rerun the connection to the database if the file does not exist.
	fileName := fmt.Sprintf("./%v", os.Getenv("DB_SQLITE_FILENAME"))
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		log.Println("info - SQLite: Database file does not exist, recreating")
		LoadDbConnectToSqlite()
	}
}

func LoadDbSingleScenarioToSqlite(scenario utils.Scenario, scenarioFileModificationTime string) {
	// TODO: Need a way to get the correct file path no matter the OS.
	// Check if the scenario exists and file_last_modified is different
	var existingScenario utils.Scenario
	err := DB.Table("scenarios").Where("name = ?", scenario.Name).First(&existingScenario).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("info - SQLite: Failed to query table %s: %v", "scenarios", err)
	}

	if existingScenario.FileLastModified != scenarioFileModificationTime {
		// Save = Upsert: scenario table entries
		DB.Table("scenarios").Where(utils.Scenario{Name: scenario.Name}).Assign(utils.Scenario{
			Type:                    scenario.Type,
			EnableTestVirtruEncrypt: scenario.EnableTestVirtruEncrypt,
			EnableTestDLP:           scenario.EnableTestDLP,
			FromEmails:              scenario.FromEmails,
			ToEmails:                scenario.ToEmails,
			Description:             scenario.Description,
			AttachmentFilePath:      scenario.AttachmentFilePath,
			Hosts:                   scenario.Hosts,
			Ports:                   scenario.Ports,
			Endpoints:               scenario.Endpoints,
			ClientId:                scenario.ClientId,
			SmtpUsername:            scenario.SmtpUsername,
			SmtpPassword:            scenario.SmtpPassword,
			FileLastModified:        scenario.FileLastModified,
		}).FirstOrCreate(&utils.Scenario{
			Type:                    scenario.Type,
			EnableTestVirtruEncrypt: scenario.EnableTestVirtruEncrypt,
			EnableTestDLP:           scenario.EnableTestDLP,
			FromEmails:              scenario.FromEmails,
			ToEmails:                scenario.ToEmails,
			Description:             scenario.Description,
			AttachmentFilePath:      scenario.AttachmentFilePath,
			Hosts:                   scenario.Hosts,
			Ports:                   scenario.Ports,
			Endpoints:               scenario.Endpoints,
			ClientId:                scenario.ClientId,
			SmtpUsername:            scenario.SmtpUsername,
			SmtpPassword:            scenario.SmtpPassword,
			FileLastModified:        scenario.FileLastModified,
		})
	}
}

func LoadDbSingleCredentialToSqlite(scenario utils.Scenario, scenarioFileModificationTime string) {
	// Save = Upsert: credentials table entries
	DB.Table("credentials").Where(utils.ScenarioAuth{CredentialName: scenario.Name}).Assign(utils.ScenarioAuth{
		ClientId:                            scenario.ClientId,
		ClientSecret:                        scenario.ClientSecret,
		TenantId:                            scenario.TenantId,
		GraphApiToken:                       "",
		TokenExpireAtTimeStampMilliseconds:  0,
		TokenUpdatedAtTimeStampMilliseconds: 0,
		SmtpUsername:                        scenario.SmtpUsername,
		SmtpPassword:                        scenario.SmtpPassword,
	}).FirstOrCreate(&utils.ScenarioAuth{
		ClientId:                            scenario.ClientId,
		ClientSecret:                        scenario.ClientSecret,
		TenantId:                            scenario.TenantId,
		CredentialName:                      scenario.Name,
		GraphApiToken:                       "",
		TokenExpireAtTimeStampMilliseconds:  0,
		TokenUpdatedAtTimeStampMilliseconds: 0,
		SmtpUsername:                        scenario.SmtpUsername,
		SmtpPassword:                        scenario.SmtpPassword,
	})
	// }
}

// func LoadDbMultipleScenariosToSqlite(tableName string) {
func LoadDbMultipleScenariosToSqlite() {
	LoadDbVerifySqliteExists()
	// Get current scenarios.csv file last modification time.
	fileInfo, err := os.Stat(utils.ScenariosFilePath)
	if err != nil {
		fmt.Print(err)
	}
	ScenarioFileModificationTime := fileInfo.ModTime().Format(time.RFC3339)

	log.Println("- Loading scenarios")
	scenarios, err := utils.ParseScenariosCSV(utils.ScenariosFilePath)
	if err != nil {
		log.Print(err)
		return
	}
	if len(scenarios) == 0 {
		log.Println("-- There are no scenarios to process")
		return
	}
	for i, scenario := range scenarios {
		log.Printf("-- Scenario %v: Loading to database\n", i+1)
		LoadDbSingleScenarioToSqlite(scenario, ScenarioFileModificationTime)
		LoadDbSingleCredentialToSqlite(scenario, ScenarioFileModificationTime)
	}
	log.Println("-- Loading scenarios, complete")
}
