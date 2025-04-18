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

func LoadDbSingleScenarioToSqlite(scenario utils.Scenario, tableName string, scenarioFileModificationTime string) {
	// TODO: Need a way to get the correct file path no matter the OS.
	// This will rerun the connection to the database if the file does not exist.
	fileName := fmt.Sprintf("./%v", os.Getenv("DB_SQLITE_FILENAME"))
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		log.Println("info - SQLite: Database file does not exist, recreating")
		LoadDbConnectToSqlite()
	}

	// Query the table
	// query := fmt.Sprintf("SELECT `file_last_modified` FROM %s", tableName)
	// var fileLastModified string
	// err := DB.Raw(query).Scan(&fileLastModified).Error
	// if err != nil {
	// 	log.Fatalf("error - SQLite: Failed to query table %s: %v", tableName, err)
	// }

	// getFileModTime := DB.Table(tableName).Select("file_last_modified").Where(utils.Scenario{Name: scenario.Name})
	// fmt.Print(getFileModTime)

	// Check if the scenario exists and file_last_modified is different
	var existingScenario utils.Scenario
	err := DB.Table(tableName).Where("name = ?", scenario.Name).First(&existingScenario).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Fatalf("error - SQLite: Failed to query table %s: %v", tableName, err)
	}

	fmt.Println(existingScenario)

	// if existingScenario.FileLastModified != scenarioFileModificationTime {
	// 	// Save = Upsert
	// 	DB.Table(tableName).Where(utils.Scenario{Name: scenario.Name}).Assign(utils.Scenario{
	// 		Type:               scenario.Type,
	// 		CredentialLocation: scenario.CredentialLocation,
	// 		FromEmail:          scenario.FromEmail,
	// 		ToEmail:            scenario.ToEmail,
	// 		Description:        scenario.Description,
	// 		AttachmentFilePath: scenario.AttachmentFilePath,
	// 		FileLastModified:   scenario.FileLastModified,
	// 	}).FirstOrCreate(&utils.Scenario{
	// 		Type:               scenario.Type,
	// 		CredentialLocation: scenario.CredentialLocation,
	// 		FromEmail:          scenario.FromEmail,
	// 		ToEmail:            scenario.ToEmail,
	// 		Description:        scenario.Description,
	// 		AttachmentFilePath: scenario.AttachmentFilePath,
	// 		FileLastModified:   scenario.FileLastModified,
	// 	})
	// }
}

func LoadDbMultipleScenariosToSqlite(tableName string) {
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
		log.Printf("-- Scenario %v loaded to database\n", i+1)
		LoadDbSingleScenarioToSqlite(scenario, tableName, ScenarioFileModificationTime)
	}
	log.Println("-- Loading scenarios, complete")
}
