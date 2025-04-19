package db

import (
	"errors"
	"fmt"
	"log"
	"mail-telemetry/utils"
)

func RetrieveScenarioFromSqliteAll(tableName string) ([]utils.Scenario, error) {
	var scenarios []utils.Scenario
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

func RetrieveScenarioFromSqliteByColumnName(tableName, columnName, scenarioName string) ([]utils.Scenario, error) {
	var scenarios []utils.Scenario

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
