package email

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
 Send/Receive emails via GraphQL: https://learn.microsoft.com/en-us/graph/outlook-create-send-messages
 Endpoint
  - Message: https://learn.microsoft.com/en-us/graph/api/resources/message?view=graph-rest-1.0
  - AA:
    - https://learn.microsoft.com/en-us/graph/api/resources/message?view=graph-rest-1.0
	- https://learn.microsoft.com/en-us/exchange/client-developer/legacy-protocols/how-to-authenticate-an-imap-pop-smtp-application-by-using-oauth
 Example
  - Create a MS Graph client: https://learn.microsoft.com/en-us/graph/sdks/create-client?tabs=go
  - https://learn.microsoft.com/en-us/answers/questions/936444/how-to-send-an-email-with-multiple-contents-using
  - https://stackoverflow.com/questions/64481592/sending-office-365-email
 Packages
  - https://pkg.go.dev/github.com/recolabs/go-office365
*/

var CLIENT_ID string
var TENANT_ID string
var CLIENT_SECRET string
var GRAPH_USER_SCOPES string
var GRAPH_GRANT_TYPE string

func InitializeEnvValuesOF365() {
	CLIENT_ID = os.Getenv("CLIENT_ID")
	TENANT_ID = os.Getenv("TENANT_ID")
	CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
	GRAPH_USER_SCOPES = os.Getenv("GRAPH_USER_SCOPES")
	GRAPH_GRANT_TYPE = os.Getenv("GRAPH_GRANT_TYPE")
}

func ApiCallGetToken() (string, error) {
	tokenUrl := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", TENANT_ID)
	tokenQuery := fmt.Sprintf("scope=%s&grant_type=%s&client_id=%s&client_secret=%s", GRAPH_USER_SCOPES, GRAPH_GRANT_TYPE, CLIENT_ID, CLIENT_SECRET)
	queryPayload := strings.NewReader(tokenQuery)
	client := &http.Client{}
	req, err := http.NewRequest("GET", tokenUrl, queryPayload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Unmarshal JSON into a map
	jsonStr := string(body)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return "", err
	}

	// Extract a property
	token, ok := data["access_token"].(string)
	if !ok {
		fmt.Println("Error extracting 'name' property")
		return "", err
	}

	// For Testing: Print the extracted property
	// fmt.Println("Access Token:", token)

	return token, nil
}

func ApiCallSendMail(sendMailConfig map[string]interface{}) {
	apiToken, ok := sendMailConfig["apiToken"].(string)
	if !ok || apiToken == "" {
		fmt.Printf("\nInvalid/Empty token - Received: %v\n", sendMailConfig["apiToken"])
		return
	}

	url := fmt.Sprintf("https://graph.microsoft.com/v1.0/users/%s/sendMail", sendMailConfig["senderEmail"])
	payloadString := fmt.Sprintf(`{
        "message": {
            "subject": "%s",
            "body": {
                "contentType": "HTML",
                "content": "%s"
            },
            "toRecipients": [
                {
                    "emailAddress": {
                        "address": "%s"
                    }
                }
            ]
        }
    }`, sendMailConfig["subject"], sendMailConfig["subject"], sendMailConfig["recipientEmail"])
	payload := strings.NewReader(payloadString)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	authorizationString := fmt.Sprintf("Bearer %s", apiToken)
	req.Header.Add("Authorization", authorizationString)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
