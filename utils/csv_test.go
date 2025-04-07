package utils

import (
	"reflect"
	"testing"
)

func TestParseScenariosCSV(t *testing.T) {
	scenarios, err := ParseScenariosCSV("../test/scenarios_test.csv")
	if err != nil {
		t.Error(err)
		return
	}
	varLength := len(scenarios)
	varType := reflect.ValueOf(scenarios).Kind().String()
	if varType != "slice" {
		errorString := FormatTestFailureString("Var Type", varType, "slice")
		t.Error(errorString)
	}
	if varLength != 2 {
		errorString := FormatTestFailureString("Slice Length", varLength, 2)
		t.Error(errorString)
	}
}
