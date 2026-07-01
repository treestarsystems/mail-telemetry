package email

import (
	"encoding/json"
	"fmt"
	"io"
	"mail-telemetry/db"
	"mail-telemetry/utils"
	"net/http"
	"os"
	"strings"
)

var GRAPH_USER_SCOPES string
var GRAPH_GRANT_TYPE string

func InitializeEnvValuesOF365() {
	GRAPH_USER_SCOPES = os.Getenv("GRAPH_USER_SCOPES")
	GRAPH_GRANT_TYPE = os.Getenv("GRAPH_GRANT_TYPE")
}

// The credentials should be loaded into the DB. Then this func takes the clientId and retrieves that info
func GraphApiGenerateToken(scenarioAuth utils.ScenarioAuth) (string, error) {
	tokenUrl := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", scenarioAuth.TenantId)
	tokenQuery := fmt.Sprintf("scope=%s&grant_type=%s&client_id=%s&client_secret=%s", GRAPH_USER_SCOPES, GRAPH_GRANT_TYPE, scenarioAuth.ClientId, scenarioAuth.ClientSecret)
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
		fmt.Println("error - Graph Api Token Generation: unmarshaling JSON:", err)
		return "", err
	}

	// Extract a property
	token, ok := data["access_token"].(string)
	if !ok {
		fmt.Println("error - Graph Api Token Generation: extracting 'name' property")
		return "", err
	}

	// For Testing: Print the extracted property
	// fmt.Println("Access Token:", token)

	return token, nil
}

// Function handles the interaction of retrieving the token from the db, checking expiration, refreshing token, then returning token string
func GraphApiHandleCacheToken(scenarioDetails utils.ScenarioDetails) (string, error) {
	// Retrieve scenario credentials from DB.
	scenarioCredentials, err := db.RetrieveCredentialFromSqliteByColumnName("credentials", "credential_name", scenarioDetails.Details.Name)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Check if nothing is returned
	if len(scenarioCredentials) == 0 {
		graphApiToken, err := GraphApiGenerateToken(scenarioCredentials[0])
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		return graphApiToken, nil

	}
	// Check if existing token is empty
	if scenarioCredentials[0].GraphApiToken == "" {
		graphApiToken, err := GraphApiGenerateToken(scenarioCredentials[0])
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		return graphApiToken, nil
	}

	return "", nil
}
