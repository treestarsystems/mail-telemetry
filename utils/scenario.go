package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net/mail"
	"os"
	"slices"
	"strings"
)

func ParseScenariosCSV(csvFilePath string) ([]Scenario, error) {
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
	var headersMap map[int]string

	// Set default value to determine futher processing.
	isValidScenariosFile := false

	// This validation loop will invalidate the WHOLE file if one row is malformed/incorrect.
	for lineIndex, line := range data {
		lineNumber := lineIndex + 1

		// Generate map of headers and their associated index to use later in if-statments.
		if lineIndex == 0 {
			headersMap, err = CreateScenariosHeadersMap(line)
			if err != nil {
				return scenarios, err
			}
		}

		// Validate field data if not header (first row).
		if lineIndex != 0 {
			// Check that all fields contain the correct data.
			err := ValidateScenarioLine(headersMap, line, lineNumber)
			if err != nil {
				return scenarios, err
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

	return scenarios, nil
}

func CreateScenariosHeadersMap(headersLine []string) (map[int]string, error) {
	// Store headers as a map of field names and it's associated index.
	headersMap := make(map[int]string)

	// Generate map of headers and their associated index to use later in if-statments.
	for fieldIndex, field := range headersLine {
		headersMap[fieldIndex] = field
	}

	if len(headersMap) == 0 {
		return headersMap, errors.New("error - Empty map")
	}
	return headersMap, nil

}

// This validation line will invalidate the WHOLE file if one row is malformed/incorrect.
func ValidateScenarioLine(headersMap map[int]string, scenarioLine []string, scenarioLineNumber int) error {
	// Standard error string.
	prependErrorString := "malformed scenarios file."
	var errorResponse error
	// Validate field data if not header (first row).
	// Check that all fields contain the correct data.
	for fieldIndex, fieldData := range scenarioLine {
		currentFieldName := headersMap[fieldIndex]

		// Check that required fields are not empty strings.
		if currentFieldName != "description" {
			if fieldData == "" {
				errorString := fmt.Sprintf("%s The \"%s\" field on line/row %v can not be empty", prependErrorString, currentFieldName, scenarioLineNumber)
				log.Print(errorString)
				errorResponse = errors.New(errorString)
			}
		}

		// Check email fields have a valid email address.
		if currentFieldName == "from" || currentFieldName == "to" {
			_, err := mail.ParseAddress(fieldData)
			if err != nil {
				errorString := fmt.Sprintf("%s The %s email on line/row %v is not valid: %v", prependErrorString, currentFieldName, scenarioLineNumber, fieldData)
				log.Print(errorString)
				errorResponse = errors.New(errorString)
			}
		}

		// Check credentialsLocation has a valid value.
		if currentFieldName == "credentialLocation" {
			validCredentialLocationStrings := []string{"file", "database", "secretstore"}
			if !slices.Contains(validCredentialLocationStrings, fieldData) {
				errorString := fmt.Sprintf("%s Invalid credentialLocation string on line/row %v of scenarios file. Should be either one of the following values: [%v]", prependErrorString, scenarioLineNumber, strings.Join(validCredentialLocationStrings, ","))
				log.Print(errorString)
				errorResponse = errors.New(errorString)
			}
		}
	}
	return errorResponse
}
