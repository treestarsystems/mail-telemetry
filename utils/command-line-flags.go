package utils

import (
	"flag"
	"fmt"
	"os"
)

var EnvFilePath string
var ScenariosFilePath string

func InitCommandLineFlags() {
	// Define the flags
	h := flag.String("h", "", "Show help")
	e := flag.String("e", "./.env", "Path to file containing Environment variables for this application")
	s := flag.String("s", "./scenarios.csv", "Path to scenarios.csv file")

	// Parse the flags
	flag.Parse()

	// Show help and exit if the help flag is provided
	if *h != "" {
		flag.Usage()
		os.Exit(0)
	}

	// // Set the default value of the 'e' flag to "./.env"
	if *e == "" {
		fmt.Println("No env file path provided. Using deault value(s).")
		EnvFilePath = "./.env"
	} else {
		EnvFilePath = *e
	}

	// // Set the default value of the 'd' flag to "./job-funnel.sqlite.db"
	if *s == "" {
		fmt.Println("No scenarios.csv file path provided. Using deault value(s).")
		ScenariosFilePath = "./scenarios.csv"
	} else {
		ScenariosFilePath = *s
	}
}
