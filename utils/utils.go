package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString generates a random string of the specified length.
func RandomAplhaNumericString(length uint) string {
	var stringLength uint

	// Validate parameters
	if length == 0 {
		stringLength = 1
	} else {
		stringLength = length
	}

	b := make([]byte, stringLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func ParseCommaSeparatedStringToSlice(envVarString string) []string {
	var resultSlice []string
	// Validate parameters
	if envVarString == "" {
		return []string{}
	}

	for _, item := range strings.Split(envVarString, ",") {
		// Filter out empty strings.
		if item != "" {
			resultSlice = append(resultSlice, item)
		}
	}
	return resultSlice
}

func FormatTestFailureString(testName string, returnedValue any, expectedValue any) string {
	return fmt.Sprintf("\nxxx FAILED (%s) - Got: %v, Expected: %v", testName, returnedValue, expectedValue)
}

// Check if file exists
func CheckFileExists(filePath string) bool {
	// Validate parameters
	if filePath == "" {
		return false
	}
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// TODO: Expand this to take any type and return the a desired type...Maybe.
// If empty string, return passed defaultValue.
func NullishCoalesceString(initialValue, defaultValue string) string {
	if initialValue == "" {
		return defaultValue
	}
	return initialValue
}

// PrintStructAsPrettyJSON prints a struct as a pretty-formatted JSON string.
func PrintStructAsPrettyJSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("error - Failed to marshal struct to JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

func CombineTwoStringSlices(sliceOne, sliceTwo []string, joiningString string) []string {
	var resultSlce []string

	// If the second slice is empty, just return the first slice.
	if len(sliceTwo) == 0 {
		return sliceOne
	}

	for _, sliceOneItem := range sliceOne {
		for _, sliceTwoItem := range sliceTwo {
			resultSlce = append(resultSlce, sliceOneItem+joiningString+sliceTwoItem)
		}
	}

	return resultSlce
}
