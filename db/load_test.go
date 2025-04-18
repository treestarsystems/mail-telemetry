package db

import (
	"log"
	"testing"

	"mail-telemetry/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	// Create a temporary SQLite database for testing
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	// Migrate the Scenario schema
	err = db.AutoMigrate(&utils.Scenario{})
	if err != nil {
		log.Fatalf("Failed to migrate test database: %v", err)
	}

	return db
}

func TestLoadDbSingleScenarioToSqlite(t *testing.T) {
	// Setup test database
	DB = setupTestDB()
	fileLastModifiedTimeString := "2025-04-18T10:08:58-04:00"
	// Create a test scenario
	scenario := utils.Scenario{
		Name:               "Test Scenario",
		Type:               "OF365",
		CredentialLocation: "database",
		FromEmail:          "from@example.com",
		ToEmail:            "to@example.com",
		Description:        "This is a test scenario",
		AttachmentFilePath: "",
		FileLastModified:   fileLastModifiedTimeString,
	}

	// Call the function
	LoadDbSingleScenarioToSqlite(scenario, "scenarios", fileLastModifiedTimeString)

	// Verify the scenario was inserted into the database
	var result utils.Scenario
	err := DB.Table("scenarios").Where("name = ?", scenario.Name).First(&result).Error
	if err != nil {
		errorString := utils.FormatTestFailureString("Failed to insert scenario", err, "error to be nil")
		t.Error(errorString)
	}

	// Check if the inserted scenario matches the original
	if result.Name != scenario.Name || result.CredentialLocation != scenario.CredentialLocation ||
		result.FromEmail != scenario.FromEmail || result.ToEmail != scenario.ToEmail ||
		result.Description != scenario.Description || result.AttachmentFilePath != scenario.AttachmentFilePath ||
		result.FileLastModified != scenario.FileLastModified {
		errorString := utils.FormatTestFailureString("Failed to validate inserted scenario", err, "error to be nil")
		t.Error(errorString)
	}
}
