package email

import (
	"mail-telemetry/utils"
	"os"
)

var GRAPH_USER_SCOPES string
var GRAPH_GRANT_TYPE string

func InitializeEnvValuesOF365() {
	GRAPH_USER_SCOPES = os.Getenv("GRAPH_USER_SCOPES")
	GRAPH_GRANT_TYPE = os.Getenv("GRAPH_GRANT_TYPE")
}

// The credentails should be loaded into the DB. Then this func takes the clientId and retrieves that info
// func GraphApiGenerateToken(scenario utils.ScenarioDetailsO365) (string, error) {
// 	tokenUrl := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", scenario.Auth.TenantId)
// 	tokenQuery := fmt.Sprintf("scope=%s&grant_type=%s&client_id=%s&client_secret=%s", GRAPH_USER_SCOPES, GRAPH_GRANT_TYPE, scenario.Auth.ClientId, scenario.Auth.ClientSecret)
// 	queryPayload := strings.NewReader(tokenQuery)
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", tokenUrl, queryPayload)

// 	if err != nil {
// 		fmt.Println(err)
// 		return "", err
// 	}
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return "", err
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return "", err
// 	}

// 	// Unmarshal JSON into a map
// 	jsonStr := string(body)
// 	var data map[string]interface{}
// 	err = json.Unmarshal([]byte(jsonStr), &data)
// 	if err != nil {
// 		fmt.Println("Error unmarshaling JSON:", err)
// 		return "", err
// 	}

// 	// Extract a property
// 	token, ok := data["access_token"].(string)
// 	if !ok {
// 		fmt.Println("Error extracting 'name' property")
// 		return "", err
// 	}

// 	// For Testing: Print the extracted property
// 	// fmt.Println("Access Token:", token)

// 	return token, nil
// }

func GraphApiSendMail(scenarioSendMailConfig utils.ScenarioDetailsO365) {
	// url := fmt.Sprintf("https://graph.microsoft.com/v1.0/users/%s/sendMail", sendMailConfig.FromEmail)
	// payloadString := fmt.Sprintf(`{
	//       "message": {
	//           "subject": "%s",
	//           "body": {
	//               "contentType": "HTML",
	//               "content": "%s"
	//           },
	//           "toRecipients": [
	//               {
	//                   "emailAddress": {
	//                       "address": "%s"
	//                   }
	//               }
	//           ]
	//       },
	//       saveToSentItems: false
	//   }`, sendMailConfig.EmailSubject, sendMailConfig.EmailSubject, sendMailConfig.ToEmail)
	// payload := strings.NewReader(payloadString)

	// client := &http.Client{}
	// req, err := http.NewRequest("POST", url, payload)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// req.Header.Add("Content-Type", "application/json")
	// authorizationString := fmt.Sprintf("Bearer %s", sendMailConfig.GraphApiToken)
	// req.Header.Add("Authorization", authorizationString)

	// res, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer res.Body.Close()

	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(body))
}
