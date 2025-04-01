package main

import (
	"log"
	"mail-telemetry/api"
	"mail-telemetry/db"
	"mail-telemetry/utils"

	"github.com/joho/godotenv"
)

func main() {
	// Initialize and check for command line flags
	utils.InitCommandLineFlags()

	// Load environment variables
	err := godotenv.Load(utils.EnvFilePath)
	if err != nil {
		log.Fatalf("error - Error loading .env file: %s", err)
	}

	// Parse scenarios.csv
	utils.ParseScenariosCSV("./scenarios.csv")

	// Connect to the databases
	db.LoadDbConnectToSqlite()

	// Initial run of tasks on startup as a non-blocking goroutine
	// go tasks.InitTasks()

	// Initialize cron jobs
	// cron.InitCron()

	// Start webserver
	api.StartServer()
}
