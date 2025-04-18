package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString generates a random string of the specified length.
func RandomAplhaNumericString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func ParseEnvVarStringToArray(envVarString string) []string {
	return strings.Split(envVarString, ",")
}

func FormatTestFailureString(testName string, returnedValue any, expectedValue any) string {
	return fmt.Sprintf("\nxxx FAILED (%s) - Got: %v, Expected: %v", testName, returnedValue, expectedValue)
}

// Check if file exists
func CheckFileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
