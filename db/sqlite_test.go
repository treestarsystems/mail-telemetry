package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDbConnectToSqlite(t *testing.T) {
	// Set environment variables for testing
	tempDbFile := "test_sqlite.db"
	os.Setenv("DB_SQLITE_FILENAME", tempDbFile)
	os.Setenv("DB_TABLE_NAMES", "scenarios,credentials")
	defer os.Remove(tempDbFile) // Clean up the test database file

	// Call the function to test
	LoadDbConnectToSqlite()

	// Assert that the database connection is established
	assert.NotNil(t, DB, "Database connection should not be nil")

	// Assert that the tables are created
	tables := []string{"scenarios", "credentials"}
	for _, table := range tables {
		exists := DB.Migrator().HasTable(table)
		assert.True(t, exists, "Table %s should exist", table)
	}
}