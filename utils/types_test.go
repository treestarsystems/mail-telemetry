package utils

import (
	"encoding/json"
	"testing"
)

func TestCredentialJSONMarshalling(t *testing.T) {
	// Create a Credential instance
	credential := Credential{
		Name:         "Test Credential",
		Username:     "testuser",
		Password:     "testpassword",
		ClientId:     "test-client-id",
		ClientSecret: "test-client-secret",
	}

	// Marshal the Credential to JSON
	jsonData, err := json.Marshal(credential)
	if err == nil {
		errorString := FormatTestFailureString("Failed to marshal Credential", "byte array", err)
		t.Error(errorString)
	}

	// Pick up where we left off here.

	// Unmarshal the JSON back to a Credential
	var unmarshalledCredential Credential
	err = json.Unmarshal(jsonData, &unmarshalledCredential)
	if err != nil {
		t.Fatalf("Failed to unmarshal Credential: %v", err)
	}

	// Verify the unmarshalled data matches the original
	if credential != unmarshalledCredential {
		t.Errorf("Expected %v, got %v", credential, unmarshalledCredential)
	}
}

func TestScenarioJSONMarshalling(t *testing.T) {
	// Create a Scenario instance
	scenario := Scenario{
		Name:               "Test Scenario",
		Type:               "Test Type",
		CredentialLocation: "/path/to/credential",
		FromEmail:          "from@example.com",
		ToEmail:            "to@example.com",
		Description:        "This is a test scenario",
	}

	// Marshal the Scenario to JSON
	jsonData, err := json.Marshal(scenario)
	if err != nil {
		t.Fatalf("Failed to marshal Scenario: %v", err)
	}

	// Unmarshal the JSON back to a Scenario
	var unmarshalledScenario Scenario
	err = json.Unmarshal(jsonData, &unmarshalledScenario)
	if err != nil {
		t.Fatalf("Failed to unmarshal Scenario: %v", err)
	}

	// Verify the unmarshalled data matches the original
	if scenario != unmarshalledScenario {
		t.Errorf("Expected %v, got %v", scenario, unmarshalledScenario)
	}
}
