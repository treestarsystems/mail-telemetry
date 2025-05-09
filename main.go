package main

import (
	"log"
	"mail-telemetry/api"
	"mail-telemetry/cron"
	"mail-telemetry/db"
	"mail-telemetry/tasks"
	"mail-telemetry/utils"
	"os"

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

	// Set App Name, Hostname, and Preferred Local IP that this app uses.
	utils.AppName = os.Getenv("APP_NAME")
	utils.SystemHostName = utils.GetHostName(utils.AppName)
	utils.SystemLocalIpAddress = utils.GetPreferredLocalOutboundIP()

	// Connect to the databases
	db.LoadDbConnectToSqlite()

	// Initial run of tasks on startup
	tasks.InitTasks()

	// Initialize cron jobs
	cron.InitCron()

	// Start webserver
	api.StartServer()
}
