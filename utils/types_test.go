package utils

import (
	"encoding/json"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCredentialJSONMarshalling(t *testing.T) {
	// Create a Credential instance
	credential := Credential{
		Name:         "Test Credential",
		Username:     "testuser",
		Password:     "testpassword",
		ClientId:     "test-client-id",
		ClientSecret: "test-client-secret",
	}

	// Marshal the Credential to JSON
	jsonData, err := json.Marshal(credential)
	if err != nil {
		errorString := FormatTestFailureString("Failed to marshal Credential", err, "byte array")
		t.Error(errorString)
	}

	// Pick up where we left off here.

	// Unmarshal the JSON back to a Credential
	var unmarshalledCredential Credential
	err = json.Unmarshal(jsonData, &unmarshalledCredential)
	if err != nil {
		errorString := FormatTestFailureString("Failed to unmarshal Credential", err, "error to be nil")
		t.Error(errorString)
	}

	// Verify the unmarshalled data matches the original
	if credential != unmarshalledCredential {
		errorString := FormatTestFailureString("Match unmarshalled data with original", unmarshalledCredential, credential)
		t.Error(errorString)
	}
}

func TestScenarioJSONMarshalling(t *testing.T) {
	// Create a Scenario instance
	scenario := Scenario{
		Name:               "Test Scenario",
		Type:               "OF365",
		CredentialLocation: "database",
		FromEmail:          "from@example.com",
		ToEmail:            "to@example.com",
		Description:        "This is a test scenario",
	}

	// Marshal the Scenario to JSON
	jsonData, err := json.Marshal(scenario)
	if err != nil {
		errorString := FormatTestFailureString("Failed to marshal Scenario", err, "byte array")
		t.Error(errorString)
	}

	// Unmarshal the JSON back to a Scenario
	var unmarshalledScenario Scenario
	err = json.Unmarshal(jsonData, &unmarshalledScenario)
	if err != nil {
		errorString := FormatTestFailureString("Failed to unmarshal Scenario", err, "error to be nil")
		t.Error(errorString)
	}

	// Verify the unmarshalled data matches the original
	if scenario != unmarshalledScenario {
		errorString := FormatTestFailureString("Match unmarshalled data with original", unmarshalledScenario, scenario)
		t.Error(errorString)

	}
}

func TestLoadDbInsertGormScenarioTable(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../test/test.db"), &gorm.Config{})
	defer os.Remove("../test/test.db")
	if err != nil {
		errorString := FormatTestFailureString("Failed to create database", err, "error to be nil")
		t.Error(errorString)
	}

	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()

	// Migrate the schema
	err = db.AutoMigrate(&LoadDbInsertGormScenario{})
	if err != nil {
		errorString := FormatTestFailureString("Failed to create table based on type", err, "error to be nil")
		t.Error(errorString)
	}

	// Create a record
	scenario := LoadDbInsertGormScenario{
		ID: 1,
	}
	err = db.Create(&scenario).Error
	if err != nil {
		errorString := FormatTestFailureString("Failed to create record based on type", err, "error to be nil")
		t.Error(errorString)
	}

	// Retrieve the record
	var retrievedScenario LoadDbInsertGormScenario
	err = db.First(&retrievedScenario, 1).Error
	if err != nil {
		errorString := FormatTestFailureString("Failed to create retrieve record based on type", err, "error to be nil")
		t.Error(errorString)
	}
	if retrievedScenario.ID != uint(1) {
		errorString := FormatTestFailureString("Failed to validate retrieved record based on type", err, "error to be nil")
		t.Error(errorString)
	}
}

func TestLoadDbInsertGormCredentialTable(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../test/test.db"), &gorm.Config{})
	defer os.Remove("../test/test.db")
	if err != nil {
		errorString := FormatTestFailureString("Failed to create database", err, "error to be nil")
		t.Error(errorString)
	}

	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()

	// Migrate the schema
	err = db.AutoMigrate(&LoadDbInsertGormCredential{})
	if err != nil {
		errorString := FormatTestFailureString("Failed to create table based on type", err, "error to be nil")
		t.Error(errorString)
	}

	// Create a record
	scenario := LoadDbInsertGormCredential{
		ID: 1,
	}
	err = db.Create(&scenario).Error
	if err != nil {
		errorString := FormatTestFailureString("Failed to create record based on type", err, "error to be nil")
		t.Error(errorString)
	}

	// Retrieve the record
	var retrievedScenario LoadDbInsertGormCredential
	err = db.First(&retrievedScenario, 1).Error
	if err != nil {
		errorString := FormatTestFailureString("Failed to create retrieve record based on type", err, "error to be nil")
		t.Error(errorString)
	}
	if retrievedScenario.ID != uint(1) {
		errorString := FormatTestFailureString("Failed to validate retrieved record based on type", err, "error to be nil")
		t.Error(errorString)
	}
}
