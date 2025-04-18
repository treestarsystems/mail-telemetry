package utils

import (
	"log"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
)

func TestParseScenariosCSV(t *testing.T) {
	// Load environment variables
	err := godotenv.Load(EnvFilePath)
	if err != nil {
		log.Fatalf("error - Error loading .env file: %s", err)
	}
	// err := os.Setenv("SCENARIOS_EXPECTED_FIELD_COUNT", "7")
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	scenarios, err := ParseScenariosCSV("../test/scenarios_test.csv")
	if err != nil {
		t.Error(err)
		return
	}
	varLength := len(scenarios)
	varType := reflect.ValueOf(scenarios).Kind().String()

	// Slice created
	if varType != "slice" {
		errorString := FormatTestFailureString("Var Type", varType, "slice")
		t.Error(errorString)
	}

	// Slice should have only two records.
	if varLength != 2 {
		errorString := FormatTestFailureString("Slice Length", varLength, 2)
		t.Error(errorString)
	}

	// Check generate headers map is correct
	scenariosHeaderLine := []string{"name", "type", "credentialLocation", "from", "to", "description", "attachmentFilePath"}
	createScenariosHeadersMap, err := CreateScenariosHeadersMap(scenariosHeaderLine)
	if err != nil {
		errorString := FormatTestFailureString("Create Header Map", err, "map[0:name 1:type 2:credentialLocation 3:from 4:to 5:description 6:attachmentFilePath]")
		t.Error(errorString)
	}

	// Happy Path: Validate scenario line with empty description and attachmentFilePath fields.
	scenariosTestLineOne := []string{"internal-external", "O365", "database", "a@a.com", "b@b.com", "", ""}
	validationErrorResultOne := ValidateScenarioLine(createScenariosHeadersMap, scenariosTestLineOne, 1)
	if validationErrorResultOne != nil {
		errorString := FormatTestFailureString("Validate Scenario Line: Happy Path", err, nil)
		t.Error(errorString)
	}

	// Check that all fields are not empty except description.
	scenariosTestLineTwo := []string{"internal-external", "", "database", "a@a.com", "b@b.com", "Mail from platform users", "./test.pdf"}
	validationErrorResultTwo := ValidateScenarioLine(createScenariosHeadersMap, scenariosTestLineTwo, 1)
	expectedErrorStringTwo := "malformed scenarios file. The \"type\" field on line/row 1 can not be empty"
	// We expect an error.
	if validationErrorResultTwo.Error() != expectedErrorStringTwo {
		errorString := FormatTestFailureString("Validate Scenario Line: Empty Field", validationErrorResultTwo, expectedErrorStringTwo)
		t.Error(errorString)
	}

	// Check that one of the email fields have a valid email address.
	scenariosTestLineThree := []string{"internal-external", "OF365", "database", "a", "b@b.com", "Mail from platform users", "./test.pdf"}
	validationErrorResultThree := ValidateScenarioLine(createScenariosHeadersMap, scenariosTestLineThree, 1)
	expectedErrorStringThree := "malformed scenarios file. The from email on line/row 1 is not valid: a"
	// We expect an error.
	if validationErrorResultThree.Error() != expectedErrorStringThree {
		errorString := FormatTestFailureString("Validate Scenario Line: Invalid Email", validationErrorResultThree, expectedErrorStringThree)
		t.Error(errorString)
	}

	// Check that credentialLocation field has a valid string. ("file","database","secretstore")
	scenariosTestLineFour := []string{"internal-external", "OF365", "wrong location keyword", "a@a.com", "b@b.com", "Mail from platform users", "./test.pdf"}
	validationErrorResultFour := ValidateScenarioLine(createScenariosHeadersMap, scenariosTestLineFour, 1)
	expectedErrorStringFour := "malformed scenarios file. Invalid credentialLocation string on line/row 1 of scenarios file. Should be either one of the following values: [file,database,secretstore]"
	// We expect an error.
	if validationErrorResultFour.Error() != expectedErrorStringFour {
		errorString := FormatTestFailureString("Validate Scenario Line: Valid credentialLocation", validationErrorResultFour, expectedErrorStringFour)
		t.Error(errorString)
	}
}
