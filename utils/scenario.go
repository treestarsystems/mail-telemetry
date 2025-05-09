package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net/mail"
	"os"
	"strconv"
	"time"
)

var expectedNumberOfFieldsPerLine *uint64

func ParseScenariosCSV(csvFilePath string) ([]Scenario, error) {
	var scenarios []Scenario

	// Validate parameters
	if csvFilePath == "" {
		return scenarios, errors.New("error - ParseScenariosCSV: \"csvFilePath\" parameter can not be empty")
	}

	// Expected number of fields per line/row
	count, err := strconv.ParseUint(os.Getenv("SCENARIOS_EXPECTED_FIELD_COUNT"), 10, 32)
	if err != nil {
		errorString := fmt.Sprintf("Unable to convert ENV Variable to int: %s", err)
		return scenarios, errors.New(errorString)
	}

	expectedNumberOfFieldsPerLine = &count

	// Standard error string.
	prependErrorString := "malformed scenarios file."

	// Check if file exists
	fileExists := CheckFileExists(csvFilePath)
	if !fileExists {
		log.Fatalf("-- The necessary scenarios.csv file does not exists at the provided path: %s", csvFilePath)
	}

	// Process file
	csvFileData, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	csvFileInfo, err := csvFileData.Stat()
	if err != nil {
		log.Fatalf("%s %s", prependErrorString, err)
	}

	defer csvFileData.Close()

	// Read csv values using csv.Reader
	data, err := csv.NewReader(csvFileData).ReadAll()
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
					Name:                    line[0],
					Type:                    line[1],
					EnableTestVirtruEncrypt: NullishCoalesceString(line[2], "false"),
					EnableTestDLP:           NullishCoalesceString(line[3], "false"),
					FromEmails:              line[4],
					ToEmails:                line[5],
					Description:             line[6],
					AttachmentFilePath:      line[7],
					Hosts:                   line[8],
					Ports:                   line[9],
					Endpoints:               line[10],
					ClientId:                line[11],
					ClientSecret:            line[12],
					TenantId:                line[13],
					SmtpUsername:            line[14],
					SmtpPassword:            line[15],
					FileLastModified:        csvFileInfo.ModTime().Format(time.RFC3339),
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

func SearchHeadersMapValue(headersMap map[int]string, headerString string) (int, bool) {
	for key, value := range headersMap {
		if value == headerString {
			return key, true
		}
	}
	return -1, false
}

// This validation line will invalidate the WHOLE file if one row is malformed/incorrect.
func ValidateScenarioLine(headersMap map[int]string, scenarioLine []string, scenarioLineNumber int) error {
	var errorResponse error

	// Standard error string.
	prependErrorString := "malformed scenarios file."

	// Check the number of fields.
	if uint64(len(scenarioLine)) != *expectedNumberOfFieldsPerLine {
		errorString := fmt.Sprintf("%s Incorrect number of fields on line/row %v", prependErrorString, scenarioLineNumber)
		errorResponse = errors.New(errorString)
		return errorResponse
	}

	// Check that all fields contain the correct data.
	for fieldIndex, fieldData := range scenarioLine {
		currentFieldName := headersMap[fieldIndex]

		// Check that required fields are not empty strings.
		if currentFieldName != "enableTestVirtruEncrypt" && currentFieldName != "enableTestDLP" &&
			currentFieldName != "description" && currentFieldName != "attachmentFilePath" &&
			currentFieldName != "ports" && currentFieldName != "endpoints" && currentFieldName != "clientId" &&
			currentFieldName != "clientSecret" && currentFieldName != "tenantId" && currentFieldName != "smtpUsername" &&
			currentFieldName != "smtpPassword" {
			if fieldData == "" {
				errorString := fmt.Sprintf("%s The \"%s\" field on line/row %v can not be empty", prependErrorString, currentFieldName, scenarioLineNumber)
				errorResponse = errors.New(errorString)
				return errorResponse
			}
		}

		// Check email fields have a valid email address.
		if currentFieldName == "from" || currentFieldName == "to" {
			_, err := mail.ParseAddress(fieldData)
			if err != nil {
				errorString := fmt.Sprintf("%s The %s email on line/row %v is not valid: %v", prependErrorString, currentFieldName, scenarioLineNumber, fieldData)
				errorResponse = errors.New(errorString)
				return errorResponse
			}
		}

		// Check for data in interdependent fields using type field value.
		if currentFieldName == "type" {
			if fieldData == "O365" {
				dependentFieldNames := []string{"clientId", "clientSecret", "tenantId"}
				// Check if dependent fields are empty.
				for _, dpfn := range dependentFieldNames {
					key, found := SearchHeadersMapValue(headersMap, dpfn)
					// If found we need to check if empty, if not we return the errorResponse since the field is required for this scenario type.
					if found {
						if scenarioLine[key] == "" {
							errorString := fmt.Sprintf("%s The \"%s\" field on line/row %v can not be empty when type is \"%s\"", prependErrorString, dpfn, scenarioLineNumber, fieldData)
							errorResponse = errors.New(errorString)
							return errorResponse
						}
					} else {
						errorString := fmt.Sprintf("%s The \"%s\" field is required when type is \"%s\": Not Found", prependErrorString, dpfn, fieldData)
						errorResponse = errors.New(errorString)
						return errorResponse
					}
				}
			}
			if fieldData == "SMTP" {
				dependentFieldNames := []string{"ports"}
				// Check if dependent fields are empty.
				for _, dpfn := range dependentFieldNames {
					key, found := SearchHeadersMapValue(headersMap, dpfn)
					// If found we need to check if empty, if not we return the errorResponse since the field is required for this scenario type.
					if found {
						if scenarioLine[key] == "" {
							errorString := fmt.Sprintf("%s The \"%s\" field on line/row %v can not be empty when type is \"%s\"", prependErrorString, dpfn, scenarioLineNumber, fieldData)
							errorResponse = errors.New(errorString)
							return errorResponse
						}
					} else {
						errorString := fmt.Sprintf("%s The \"%s\" field is required when type is \"%s\": Not Found", prependErrorString, dpfn, fieldData)
						errorResponse = errors.New(errorString)
						return errorResponse
					}
				}
			}
		}
	}
	return errorResponse
}
