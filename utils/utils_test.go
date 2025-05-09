package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestRandomAplhaNumericString(t *testing.T) {
	randomAplhaNumericStringOne := RandomAplhaNumericString(10)
	ransSliceOne := strings.Split(randomAplhaNumericStringOne, "")
	if len(ransSliceOne) != 10 {
		errorString := FormatTestFailureString("String Length", ransSliceOne, 10)
		t.Error(errorString)
	}

	randomAplhaNumericStringTwo := RandomAplhaNumericString(0)
	ransSliceTwo := strings.Split(randomAplhaNumericStringTwo, "")
	if len(ransSliceTwo) != 1 {
		errorString := FormatTestFailureString("String Length", ransSliceTwo, 1)
		t.Error(errorString)
	}
}

func TestConvertCommaSeparatedStringToSlice(t *testing.T) {
	testStringOne := "string1, string2"
	parsedTestStringOne := ConvertCommaSeparatedStringToSlice(testStringOne)
	parsedTestStringLengthOne := len(parsedTestStringOne)
	if parsedTestStringLengthOne != 2 {
		errorString := FormatTestFailureString("Slice Length", parsedTestStringLengthOne, 2)
		t.Error(errorString)
	}
	testStringTwo := ""
	parsedTestStringTwo := ConvertCommaSeparatedStringToSlice(testStringTwo)
	fmt.Println(parsedTestStringTwo)
	parsedTestStringLengthTwo := len(parsedTestStringTwo)
	if parsedTestStringLengthTwo != 0 {
		errorString := FormatTestFailureString("Slice Length", parsedTestStringLengthTwo, 0)
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
