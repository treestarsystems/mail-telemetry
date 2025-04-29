package email

import (
	"errors"
	"fmt"
	"mail-telemetry/utils"
)

/*
 - Generate message struct.
 - Wrapper to retrieve scenario from db and route to correct logic based on type.
 - Logic to get credentials from storage points.
*/

func GenerateScenarioAuth(scenario *utils.Scenario) (interface{}, error) {
	switch scenario.Type {
	case "O365":
		var scenarioAuth = utils.ScenarioAuthO365{
			ClientId:     "",
			ClientSecret: "",
			TenantId:     "",
		}
		return scenarioAuth, nil
	case "SMTP":
		var scenarioAuth = utils.ScenarioAuthSMTP{
			Username: "",
			Password: "",
		}
		return scenarioAuth, nil
	default:
		errorString := fmt.Sprintf("error - GenerateScenarioAuth: Unsupported scenario type(%s)", scenario.Type)
		return nil, errors.New(errorString)
	}
}

func GenerateScenarioHost(scenario *utils.Scenario) (interface{}, error) {
	switch scenario.Type {
	case "O365":
		var scenarioHost = utils.ScenarioHostO365{
			Address:  "",
			Port:     0,
			Endpoint: "",
		}
		return scenarioHost, nil
	case "SMTP":
		var scenarioHost = utils.ScenarioHostSMTP{
			Address: "",
			Port:    0,
		}
		return scenarioHost, nil
	default:
		errorString := fmt.Sprintf("error - GenerateScenarioHost: Unsupported scenario type(%s)", scenario.Type)
		return nil, errors.New(errorString)
	}
}

func GenerateScenarioDetails(scenario *utils.Scenario) (interface{}, error) {
	switch scenario.Type {
	case "O365":
		scenarioAuth, err := GenerateScenarioAuth(scenario)
		if err != nil {
			return nil, err
		}
		scenarioHost, err := GenerateScenarioHost(scenario)
		if err != nil {
			return nil, err
		}
		var scenarioDetails = utils.ScenarioDetailsO365{
			Scenario: *scenario,
			Auth:     scenarioAuth.(utils.ScenarioAuthO365),
			Host:     scenarioHost.(utils.ScenarioHostO365),
		}
		return scenarioDetails, nil
	case "SMTP":
		scenarioAuth, err := GenerateScenarioAuth(scenario)
		if err != nil {
			return nil, err
		}
		scenarioHost, err := GenerateScenarioHost(scenario)
		if err != nil {
			return nil, err
		}
		var scenarioDetails = utils.ScenarioDetailsSMTP{
			Scenario: *scenario,
			Auth:     scenarioAuth.(utils.ScenarioAuthSMTP),
			Host:     scenarioHost.(utils.ScenarioHostSMTP),
		}
		return scenarioDetails, nil
	default:
		errorString := fmt.Sprintf("error - GenerateScenarioHost: Unsupported scenario type(%s)", scenario.Type)
		return nil, errors.New(errorString)
	}
}

// func GenerateScenarioMessageStruct(scenario utils.Scenario) {
// 	fmt.Println(scenario.Type)
// }
