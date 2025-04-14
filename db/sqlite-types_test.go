package db

import (
	"mail-telemetry/utils"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestLoadDbInsertGormScenarioTable(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../test/test.db"), &gorm.Config{})
	defer os.Remove("../test/test.db")
	if err != nil {
		errorString := utils.FormatTestFailureString("Failed to create database", err, "error to be nil")
		t.Error(errorString)
	}

	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()

	// Migrate the schema
	err = db.AutoMigrate(&LoadDbInsertGormScenario{})
	if err != nil {
		errorString := utils.FormatTestFailureString("Failed to create table based on type", err, "error to be nil")
		t.Error(errorString)
	}

	// Create a record
	scenario := LoadDbInsertGormScenario{
		ID: 1,
	}
	err = db.Create(&scenario).Error
	if err != nil {
		errorString := utils.FormatTestFailureString("Failed to create record based on type", err, "error to be nil")
		t.Error(errorString)
	}

	// Retrieve the record
	var retrievedScenario LoadDbInsertGormScenario
	err = db.First(&retrievedScenario, 1).Error
	if err != nil {
		errorString := utils.FormatTestFailureString("Failed to create retrieve record based on type", err, "error to be nil")
		t.Error(errorString)
	}
	if retrievedScenario.ID != uint(1) {
		errorString := utils.FormatTestFailureString("Failed to validate retrieved record based on type", err, "error to be nil")
		t.Error(errorString)
	}
}

func TestLoadDbInsertGormCredentialTable(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../test/test.db"), &gorm.Config{})
	defer os.Remove("../test/test.db")
	if err != nil {
		errorString := utils.FormatTestFailureString("Failed to create database", err, "error to be nil")
		t.Error(errorString)
	}

	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()

	// Migrate the schema
	err = db.AutoMigrate(&LoadDbInsertGormCredential{})
	if err != nil {
		errorString := utils.FormatTestFailureString("Failed to create table based on type", err, "error to be nil")
		t.Error(errorString)
	}

	// Create a record
	scenario := LoadDbInsertGormCredential{
		ID: 1,
	}
	err = db.Create(&scenario).Error
	if err != nil {
		errorString := utils.FormatTestFailureString("Failed to create record based on type", err, "error to be nil")
		t.Error(errorString)
	}

	// Retrieve the record
	var retrievedScenario LoadDbInsertGormCredential
	err = db.First(&retrievedScenario, 1).Error
	if err != nil {
		errorString := utils.FormatTestFailureString("Failed to create retrieve record based on type", err, "error to be nil")
		t.Error(errorString)
	}
	if retrievedScenario.ID != uint(1) {
		errorString := utils.FormatTestFailureString("Failed to validate retrieved record based on type", err, "error to be nil")
		t.Error(errorString)
	}
}
