package utils

import (
	"strings"
	"testing"
)

func TestRandomAplhaNumericString(t *testing.T) {
	randomAplhaNumericString := RandomAplhaNumericString(10)
	ransSlice := strings.Split(randomAplhaNumericString, "")
	if len(ransSlice) != 10 {
		errorString := FormatTestFailureString("String Length", ransSlice, 10)
		t.Error(errorString)
	}
}

func TestParseEnvVarStringToArray(t *testing.T) {
	testString := "string1, string2"
	parsedTestString := ParseEnvVarStringToArray(testString)
	parsedTestStringLength := len(parsedTestString)
	if parsedTestStringLength != 2 {
		errorString := FormatTestFailureString("Slice Length", parsedTestStringLength, 2)
		t.Error(errorString)
	}
}

func TestFormatTestFailureString(t *testing.T) {
	input := FormatTestFailureString("Test failure string", "correct", "incorrect")
	expectedOutput := "\nxxx FAILED (Test failure string) - Got: correct, Expected: incorrect"
	if input != expectedOutput {
		errorString := FormatTestFailureString("Test failre string", input, expectedOutput)
		t.Error(errorString)
	}
}
