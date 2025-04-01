package utils

import (
	"reflect"
	"testing"
)

func TestParseScenariosCSV(t *testing.T) {
	scenarios := ParseScenariosCSV("../test/scenarios_test.csv")
	varLength := len(scenarios)
	varType := reflect.ValueOf(scenarios).Kind().String()
	if varType != "slice" {
		errorString := FormatTestFailureString("Var Type", varType, "slices")
		t.Error(errorString)
	}
	if varLength != 2 {
		errorString := FormatTestFailureString("Slice Length", varLength, 2)
		t.Error(errorString)
	}
}
