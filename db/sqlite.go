package db

import (
	"log"
	"mail-telemetry/utils"
	"os"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLite variables
var DB *gorm.DB
var SQLiteDbName *string

func LoadDbConnectToSqlite() {
	tableNames := strings.Split(os.Getenv("DB_TABLE_NAMES"), ",")
	hasScenarioQueueTable := false
	for _, tableName := range tableNames {
		if strings.TrimSpace(tableName) == "scenarioQueue" {
			hasScenarioQueueTable = true
			break
		}
	}
	if !hasScenarioQueueTable {
		tableNames = append(tableNames, "scenarioQueue")
	}

	sqliteDbName := os.Getenv("DB_SQLITE_FILENAME")
	SQLiteDbName = &sqliteDbName

	db, err := gorm.Open(sqlite.Open(*SQLiteDbName), &gorm.Config{})
	if err != nil {
		log.Printf("error - SQLite: Unable to establish database connection: %s\n", err)
	}

	// Migrate the schema/Create the tables.
	for _, tableName := range tableNames {
		tableName = strings.TrimSpace(tableName)
		if tableName == "scenarios" {
			err = db.Table(tableName).AutoMigrate(&utils.LoadDbInsertGormScenario{})
			if err != nil {
				log.Printf("error - SQLite: Unable to migrate the Scenarios schema: %s\n", err)
			}
		}
		if tableName == "credentials" {
			err = db.Table(tableName).AutoMigrate(&utils.LoadDbInsertGormCredential{})
			if err != nil {
				log.Printf("error - SQLite: Unable to migrate the Credentials schema: %s\n", err)
			}
		}
		if tableName == "scenarioQueue" {
			err = db.Table(tableName).AutoMigrate(&utils.LoadDbInsertGormScenarioQueue{})
			if err != nil {
				log.Printf("error - SQLite: Unable to migrate the ScenarioQueue schema: %s\n", err)
			}
		}
	}
	DB = db
}
