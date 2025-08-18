package db

import (
	"errors"
	"fmt"
	"log"
	"mail-telemetry/utils"
)

func RetrieveScenarioFromSqliteAll(tableName string) ([]utils.ScenarioTable, error) {
	var scenarios []utils.ScenarioTable
	if tableName == "" {
		return scenarios, errors.New("error - RetrieveScenarioFromSqliteAll: tableName can not be empty")
	}
	// Retrieve all records from the table
	err := DB.Table(tableName).Find(&scenarios).Error
	if err != nil {
		log.Printf("error - SQLite: Failed to retrieve records from table %s: %v", tableName, err)
		return scenarios, err
	}

	return scenarios, nil
}

func RetrieveScenarioFromSqliteByColumnName(tableName, columnName, scenarioName string) ([]utils.ScenarioTable, error) {
	var scenarios []utils.ScenarioTable

	// Validate parameters
	if tableName == "" {
		return scenarios, errors.New("error - RetrieveScenarioFromSqliteByName: tableName cannot be empty")
	}
	if columnName == "" {
		return scenarios, errors.New("error - RetrieveScenarioFromSqliteByName: columnName cannot be empty")
	}
	if scenarioName == "" {
		return scenarios, errors.New("error - RetrieveScenarioFromSqliteByName: scenarioName cannot be empty")
	}

	whereClauseColumnName := fmt.Sprintf("%s = ?", columnName)

	// Retrieve matching records from the table
	err := DB.Table(tableName).Where(whereClauseColumnName, scenarioName).Find(&scenarios).Error
	if err != nil {
		log.Printf("error - SQLite: Failed to retrieve record(s) from table %s: %v", tableName, err)
		return scenarios, err
	}

	return scenarios, nil
}

func RetrieveCredentialFromSqliteAll(tableName string) ([]utils.ScenarioAuth, error) {
	var scenarioCredential []utils.ScenarioAuth
	if tableName == "" {
		return scenarioCredential, errors.New("error - RetrieveCredentialFromSqliteAll: tableName can not be empty")
	}
	// Retrieve all records from the table
	err := DB.Table(tableName).Find(&scenarioCredential).Error
	if err != nil {
		log.Printf("error - SQLite: Failed to retrieve records from table %s: %v", tableName, err)
		return scenarioCredential, err
	}

	return scenarioCredential, nil
}

func RetrieveCredentialFromSqliteByColumnName(tableName, columnName, scenarioName string) ([]utils.ScenarioAuth, error) {
	var scenarioCredentials []utils.ScenarioAuth

	// Validate parameters
	if tableName == "" {
		return scenarioCredentials, errors.New("error - RetrieveCredentialFromSqliteByColumnName: tableName cannot be empty")
	}
	if columnName == "" {
		return scenarioCredentials, errors.New("error - RetrieveCredentialFromSqliteByColumnName: columnName cannot be empty")
	}
	if scenarioName == "" {
		return scenarioCredentials, errors.New("error - RetrieveCredentialFromSqliteByColumnName: scenarioName cannot be empty")
	}

	whereClauseColumnName := fmt.Sprintf("%s = ?", columnName)

	// Retrieve matching records from the table
	err := DB.Table(tableName).Where(whereClauseColumnName, scenarioName).Find(&scenarioCredentials).Error
	if err != nil {
		log.Printf("error - SQLite: Failed to retrieve record(s) from table %s: %v", tableName, err)
		return scenarioCredentials, err
	}

	return scenarioCredentials, nil
}
