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
		errorString := fmt.Sprintf("error - GenerateScenarioAuthO365: Unsupported scenario auth type(%s)", scenario.Type)
		return nil, errors.New(errorString)
	}
}

func GenerateScenarioDetailsStruct(scenario *utils.Scenario) (interface{}, error) {
	switch scenario.Type {
	case "O365":
		scenarioAuth, err := GenerateScenarioAuth(scenario)
		if err != nil {
			return nil, err
		}
		// scenarioHost, err := GenerateScenarioHost(scenario)
		// if err != nil {
		// 	return nil, err
		// }
		var scenarioDetails = utils.ScenarioDetailsO365{
			Scenario: *scenario,
			Auth:     scenarioAuth.(utils.ScenarioAuthO365),
			// Host: scenarioHost.(utils.ScenarioHostO365),
		}
		return scenarioDetails, nil
	case "SMTP":
		scenarioAuth, err := GenerateScenarioAuth(scenario)
		if err != nil {
			return nil, err
		}
		var scenarioDetails = utils.ScenarioDetailsSMTP{
			Scenario: *scenario,
			Auth:     scenarioAuth.(utils.ScenarioAuthSMTP),
		}
		return scenarioDetails, nil
	default:
		return nil, errors.New("error - GenerateScenarioDetailsStruct: Unsupported scenario type")
	}
}

// func GenerateScenarioMessageStruct(scenario utils.Scenario) {
// 	fmt.Println(scenario.Type)
// }
