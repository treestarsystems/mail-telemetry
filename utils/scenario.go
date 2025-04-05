package utils

import (
	"encoding/csv"
	"fmt"
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

	// Standard error string.
	prependErrorString := "Malformed scenarios file."

	// read csv values using csv.Reader
	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalf("%s %s", prependErrorString, err)
	}

	// Check for required headers in file before further processing.
	isValidScenariosFile := false
scenarioValidationLoop:
	for _, line := range data {
		// lineNumber := i + 1
		fmt.Print(line)

		// *** The library has a check for field counts per line.
		// Check each line has the correct amount of fields.
		// expectedFieldLength := 6
		// if len(line) != expectedFieldLength {
		// 	log.Fatalf("%s Each line should contain %v comma separated values. Line %v contains %v items.", prependErrorString, expectedFieldLength, lineNumber, len(line))
		// 	break scenarioValidationLoop
		// }

		// Check that all fields contain data.
		for j, field := range line {
			// Check that required fields are not empty strings.
			if field != "description" {
				if field == "" {
					log.Fatalf("%s %v - This field can not be empty", prependErrorString, field)
				}
			}

			// 	// Check email fields have a valid email address.
			// 	if field == "from" || field == "to" {
			// 		_, err := mail.ParseAddress(field)
			// 		if err != nil {
			// 			log.Fatalf("%s The ", prependErrorString)
			// 		}
			// 	}
			// 	// Check credentialsLocation has a valid value.
			// 	fmt.Print(j, field)
			// 	if field == "credentialLocation" {
			// 		validCredentialLocationStrings := []string{"file", "database", "secretstore"}
			// 		if !slices.Contains(validCredentialLocationStrings, field) {
			// 			log.Fatalf("%s Invalid credentialLocation string. Shold be either one of the following values: %v", prependErrorString, strings.Join(validCredentialLocationStrings, ""))
			// 			return scenarios
			// 		}
			// 	}
		}
		isValidScenariosFile = true
	}

	// Organize csv data into struct
	if isValidScenariosFile {
		for i, line := range data {
			if i > 0 { // omit header line		line := Scenario{
				scenarios = append(scenarios, Scenario{
					Name:               line[0],
					Type:               line[1],
					CredentialLocation: line[2],
					FromEmail:          line[3],
					ToEmail:            line[4],
					Description:        line[5],
				})
			}
		}
	}
	return scenarios
}
