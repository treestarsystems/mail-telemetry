package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestRandomAplhaNumericString(t *testing.T) {
	randomAplhaNumericString := RandomAplhaNumericString(10)
	// ransSlice := []rune(randomAplhaNumericString)
	ransSlice := strings.Split(randomAplhaNumericString, "")
	fmt.Print(ransSlice)
	if len(ransSlice) != 10 {
		t.Errorf("String Length - Got: %v, wanted: 10", len(ransSlice))
	}
}

func TestParseEnvVarStringToArray(t *testing.T) {
	testString := "string1, string2"
	parsedTestString := ParseEnvVarStringToArray(testString)
	if len(parsedTestString) != 2 {
		t.Errorf("Slice Length - Got: %v, Wanted: 10", len(parsedTestString))
	}

}
