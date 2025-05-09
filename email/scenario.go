package email

import (
	"errors"
	"fmt"
	"log"
	"mail-telemetry/utils"
	"strings"
	"time"
)

func ConvertEmailToPlusAddress(emailStringSlice []string, appName, messageId string) []string {
	var modifiedEmails []string

	for _, emailString := range emailStringSlice {
		// Split the email into local part and domain
		parts := strings.Split(emailString, "@")
		if len(parts) == 2 {
			// Add the messageId to the local part using a "+" address
			modifiedEmail := fmt.Sprintf("<%s: %s>%s+%s@%s", appName, messageId, parts[0], messageId, parts[1])
			modifiedEmails = append(modifiedEmails, modifiedEmail)
		} else {
			// If the email is invalid, keep it unchanged
			modifiedEmails = append(modifiedEmails, emailString)
		}
	}

	return modifiedEmails
}

// GenerateCustomTimestampString generates a string in the format HH:MM:SS_MM-DD-YYYY.
func GenerateCustomTimestampString() string {
	/* Help:
	- https://www.golinuxcloud.com/golang-time-format/
	- https://gosamples.dev/date-time-format-cheatsheet/
	*/
	currentTime := time.Now().Local() // Ensure the time is in the system's local time zone
	timeZone, _ := currentTime.Zone() // Get the time zone designation
	return fmt.Sprintf("%s(%s)", currentTime.Format("15:04:05_01-02-2006"), timeZone)
}

func GenerateScenarioSubjectString(messageId string) string {
	timeStamp := GenerateCustomTimestampString()
	subject := fmt.Sprintf("%s_%s_%s", "Telemetry", timeStamp, messageId)
	return subject
}

func GenerateMessageBodies(scenario *utils.Scenario, scenarioHostInstance, messageId string) (string, string) {
	var messageBodyTextPlain, messageBodyHtml string

	messageBodyTextPlainTemplate := `Scenario Name: %s
	Type: %s
	Virtru Encrypt: %s
	DLP: %s
	Description: %s
	Attachment: %s
	Host Instance URI: %s
	Origin Hostname: %s
	Origin Host IP: %s
	Message ID: %s`

	messageBodyHtmlTemplate := `<html><body>
	<b>Scenario Name:</b> %s</br>
	<b>Type:</b> %s</br>
	<b>Virtru Encrypt:</b> %s</br>
	<b>DLP:</b> %s</br>
	<b>Description:</b> %s</br>
	<b>Attachment:</b> %s</br>
	<b>Host Instance URI:</b> %s</br>
	<b>Origin Hostname:</b> %s</br>
	<b>Origin Host IP:</b> %s</br>
	<b>Message ID:</b> %s</br>
	</body></html>`

	messageBodyTextPlain = fmt.Sprintf(messageBodyTextPlainTemplate, scenario.Name,
		scenario.Type, scenario.EnableTestVirtruEncrypt, scenario.EnableTestDLP,
		scenario.Description, scenario.AttachmentFilePath, scenarioHostInstance, utils.SystemHostName, utils.SystemLocalIpAddress, messageId)
	messageBodyHtml = fmt.Sprintf(messageBodyHtmlTemplate, scenario.Name,
		scenario.Type, scenario.EnableTestVirtruEncrypt, scenario.EnableTestDLP,
		scenario.Description, scenario.AttachmentFilePath, scenarioHostInstance, utils.SystemHostName, utils.SystemLocalIpAddress, messageId)

	return messageBodyTextPlain, messageBodyHtml
}

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

func GenerateScenarioHost(scenario *utils.Scenario) ([]interface{}, error) {
	var scenarioHostInstances []interface{}

	// Create slices for hosts,ports, and endpoints.
	hosts := utils.ConvertCommaSeparatedStringToSlice(scenario.Hosts)
	ports := utils.ConvertCommaSeparatedStringToSlice(scenario.Ports)
	endpoints := utils.ConvertCommaSeparatedStringToSlice(scenario.Endpoints)

	// Combine slices for host addresses.
	hostAddressesWithPort := utils.CombineTwoStringSlices(hosts, ports, ":")

	switch scenario.Type {
	case "O365":
		hostFullUris := utils.CombineTwoStringSlices(hostAddressesWithPort, endpoints, "")

		for _, uri := range hostFullUris {
			var scenarioHostInstanceSingle = utils.ScenarioHostO365{
				InstanceURI: uri,
			}
			scenarioHostInstances = append(scenarioHostInstances, scenarioHostInstanceSingle)
		}
		return scenarioHostInstances, nil
	case "SMTP":
		for _, uri := range hostAddressesWithPort {
			var scenarioHostInstanceSingle = utils.ScenarioHostSMTP{
				InstanceURI: uri,
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

func GenerateScenarioMessage(scenario *utils.Scenario, scenarioHostInstance string) utils.ScenarioMessage {
	messageId := utils.RandomAplhaNumericString(20)
	// Convert email list to slices
	fromEmailsToSlice := utils.ConvertCommaSeparatedStringToSlice(scenario.FromEmails)
	toEmailsToSlice := utils.ConvertCommaSeparatedStringToSlice(scenario.ToEmails)
	// Modify email list slice for use. <App Name: Message ID(20 chars)>email+<MessageID>@example.com
	fromEmailList := ConvertEmailToPlusAddress(fromEmailsToSlice, utils.AppName, messageId)
	toEmailList := ConvertEmailToPlusAddress(toEmailsToSlice, utils.AppName, messageId)
	subject := GenerateScenarioSubjectString(messageId)
	bodyPlainText, bodyHtml := GenerateMessageBodies(scenario, scenarioHostInstance, messageId)

	return utils.ScenarioMessage{
		ID:            messageId,
		FromEmails:    fromEmailList,
		ToEmails:      toEmailList,
		Subject:       subject,
		BodyPlainText: bodyPlainText,
		BodyHTML:      bodyHtml,
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

		for _, instanceHostDetails := range scenarioHostInstances {
			scenarioInstanceDetails := utils.ScenarioDetailsO365{
				Scenario: *scenario,
				Auth:     scenarioAuth.(utils.ScenarioAuthO365),
				Host:     instanceHostDetails.(utils.ScenarioHostO365),
				Message:  GenerateScenarioMessage(scenario, instanceHostDetails.(utils.ScenarioHostO365).InstanceURI),
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

		for _, instanceHostDetails := range scenarioHostInstances {
			scenarioInstanceDetails := utils.ScenarioDetailsSMTP{
				Scenario: *scenario,
				Auth:     scenarioAuth.(utils.ScenarioAuthSMTP),
				Host:     instanceHostDetails.(utils.ScenarioHostSMTP),
				Message:  GenerateScenarioMessage(scenario, instanceHostDetails.(utils.ScenarioHostSMTP).InstanceURI),
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
