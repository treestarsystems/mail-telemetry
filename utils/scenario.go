package utils

import (
	"encoding/csv"
	"log"
	"net/mail"
	"os"
	"slices"
	"strings"
)

func ParseScenariosCSV(csvFilePath string) []Scenario {
	var scenarios []Scenario
	f, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// Standard error string.
	prependErrorString := "malformed scenarios file."

	// Read csv values using csv.Reader
	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalf("%s %s", prependErrorString, err)
	}

	// Store headers as a map of field names and it's associated index.
	headersMap := make(map[int]string)

	// Set default value to determine futher processing.
	isValidScenariosFile := false

	// This validation loop will invalidate the WHOLE file if one row is malformed/incorrect.
	for lineIndex, line := range data {
		lineNumber := lineIndex + 1
		// Generate map of headers and their associated index to use later in if-statments.
		if lineIndex == 0 {
			for fieldIndex, field := range line {
				headersMap[fieldIndex] = field
			}
		} else {
			// Check that all fields contain the correct data.
			for fieldIndex, fieldData := range line {
				currentFieldName := headersMap[fieldIndex]
				// Check that required fields are not empty strings.
				if currentFieldName != "description" {
					if fieldData == "" {
						log.Printf("%s The \"%s\" field can not be empty.", prependErrorString, currentFieldName)
						return scenarios
					}
				}
				// Check email fields have a valid email address.
				if currentFieldName == "from" || currentFieldName == "to" {
					_, err := mail.ParseAddress(fieldData)
					if err != nil {
						log.Printf("%s The %s email is not valid: %v.", prependErrorString, currentFieldName, fieldData)
						return scenarios
					}
				}

				// Check credentialsLocation has a valid value.
				if currentFieldName == "credentialLocation" {
					validCredentialLocationStrings := []string{"file", "database", "secretstore"}
					if !slices.Contains(validCredentialLocationStrings, fieldData) {
						log.Printf("%s Invalid credentialLocation string on line/row %v of scenarios file. Should be either one of the following values: [%v].", prependErrorString, lineNumber, strings.Join(validCredentialLocationStrings, ","))
						return scenarios
					}
				}
			}
		}
		isValidScenariosFile = true
	}

	// Organize csv data into struct
	if isValidScenariosFile {
		for lineIndex, line := range data {
			if lineIndex > 0 {
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
