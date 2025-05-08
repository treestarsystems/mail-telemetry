package email

import (
	"errors"
	"fmt"
	"log"
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
			ClientId:     scenario.ClientId,
			ClientSecret: scenario.ClientSecret,
			TenantId:     scenario.TenantId,
		}
		return scenarioAuth, nil
	case "SMTP":
		var scenarioAuth = utils.ScenarioAuthSMTP{
			Username: scenario.SmtpUsername,
			Password: scenario.SmtpPassword,
		}
		return scenarioAuth, nil
	default:
		errorString := fmt.Sprintf("error - GenerateScenarioAuth: Unsupported scenario type(%s)", scenario.Type)
		return nil, errors.New(errorString)
	}
}

// func GenerateScenarioHost(scenario *utils.Scenario) (interface{}, error) {
func GenerateScenarioHost(scenario *utils.Scenario) ([]interface{}, error) {
	var scenarioHostInstances []interface{}
	// We need one instance per:
	// O365: combination of endpoint and port.
	// SMTP: per port.

	// Create slices for hosts,ports, and endpoints.
	hosts := utils.ParseCommaSeparatedStringToSlice(scenario.Hosts)
	ports := utils.ParseCommaSeparatedStringToSlice(scenario.Ports)
	endpoints := utils.ParseCommaSeparatedStringToSlice(scenario.Endpoints)

	// Combine slices for host addresses.
	hostAddressesWithPort := utils.CombineTwoStringSlices(hosts, ports, ":")

	switch scenario.Type {
	case "O365":
		// var scenarioHost = utils.ScenarioHostO365{
		// 	Address:  "",
		// 	Port:     0,
		// 	Endpoint: "",
		// }
		// return scenarioHost, nil

		hostFullUris := utils.CombineTwoStringSlices(hostAddressesWithPort, endpoints, "")

		for _, uri := range hostFullUris {
			var scenarioHostInstanceSingle = utils.ScenarioHostO365{
				Instance:  uri,
				Addresses: hosts,
				Ports:     ports,
				Endpoints: endpoints,
			}
			scenarioHostInstances = append(scenarioHostInstances, scenarioHostInstanceSingle)
		}
		return scenarioHostInstances, nil
	case "SMTP":
		// var scenarioHost = utils.ScenarioHostSMTP{
		// 	Addresses: hosts,
		// 	Ports:     ports,
		// }
		// return scenarioHost, nil

		for _, uri := range hostAddressesWithPort {
			var scenarioHostInstanceSingle = utils.ScenarioHostSMTP{
				Instance:  uri,
				Addresses: hosts,
				Ports:     ports,
			}
			scenarioHostInstances = append(scenarioHostInstances, scenarioHostInstanceSingle)
		}
		return scenarioHostInstances, nil
	default:
		errorString := fmt.Sprintf("error - GenerateScenarioHost: Unsupported scenario type(%s)", scenario.Type)
		// TODO: I think this should return an empty struct or something better. Like a struct with defaulted values...Wait that doesnt make sense since there might be two types returned.
		return nil, errors.New(errorString)
	}
}

func GenerateScenarioInstance(scenario *utils.Scenario) []interface{} {
	var scenarioInstances []interface{}

	switch scenario.Type {
	case "O365":
		var errorMessages []string
		scenarioAuth, err := GenerateScenarioAuth(scenario)
		if err != nil {
			errorMessages = append(errorMessages, err.Error())
		}
		scenarioHostInstances, err := GenerateScenarioHost(scenario)
		if err != nil {
			errorMessages = append(errorMessages, err.Error())
		}
		// var scenarioInstanceDetails = utils.ScenarioDetailsO365{
		// 	Scenario: *scenario,
		// 	Auth:     scenarioAuth.(utils.ScenarioAuthO365),
		// 	Host:     scenarioHost.(utils.ScenarioHostO365),
		// 	Errors:   errorMessages,
		// }
		// // return scenarioDetails, nil
		// scenarioInstances = append(scenarioInstances, scenarioInstanceDetails)

		for _, instanceHostDetails := range scenarioHostInstances {
			scenarioInstanceDetails := utils.ScenarioDetailsO365{
				Scenario: *scenario,
				Auth:     scenarioAuth.(utils.ScenarioAuthO365),
				Host:     instanceHostDetails.(utils.ScenarioHostO365),
				Errors:   errorMessages,
			}

			scenarioInstances = append(scenarioInstances, scenarioInstanceDetails)
		}
	case "SMTP":
		var errorMessages []string
		scenarioAuth, err := GenerateScenarioAuth(scenario)
		if err != nil {
			errorMessages = append(errorMessages, err.Error())
		}
		scenarioHostInstances, err := GenerateScenarioHost(scenario)
		if err != nil {
			errorMessages = append(errorMessages, err.Error())
		}
		// var scenarioInstanceDetails = utils.ScenarioDetailsSMTP{
		// 	Scenario: *scenario,
		// 	Auth:     scenarioAuth.(utils.ScenarioAuthSMTP),
		// 	Host:     scenarioHost.(utils.ScenarioHostSMTP),
		// 	Errors:   errorMessages,
		// }
		// return scenarioDetails, nil
		// scenarioInstances = append(scenarioInstances, scenarioInstanceDetails)

		for _, instanceHostDetails := range scenarioHostInstances {
			scenarioInstanceDetails := utils.ScenarioDetailsSMTP{
				Scenario: *scenario,
				Auth:     scenarioAuth.(utils.ScenarioAuthSMTP),
				Host:     instanceHostDetails.(utils.ScenarioHostSMTP),
				Errors:   errorMessages,
			}

			scenarioInstances = append(scenarioInstances, scenarioInstanceDetails)
		}
	default:
		errorString := fmt.Sprintf("error - GenerateScenarioHost: Unsupported scenario type(%s)", scenario.Type)
		// return nil, errors.New(errorString)
		// scenarioInstances = append(scenarioInstances,)
		log.Print(errorString)
	}
	return scenarioInstances
}

// func GenerateScenarioMessageStruct(scenario utils.Scenario) {
// 	fmt.Println(scenario.Type)
// }
