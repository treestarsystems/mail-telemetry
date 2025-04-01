package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func ParseScenariosCSV(csvFilePath string) []Scenario {
	var scenarios []Scenario
	// open file
	f, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Organize csv data into struct
	for i, line := range data {
		if i > 0 { // omit header line		line := Scenario{
			scenarios = append(scenarios, Scenario{
				Name:        line[0],
				Type:        line[1],
				FromEmail:   line[2],
				ToEmail:     line[3],
				Description: line[4],
			})
		}
	}
	return scenarios
}
