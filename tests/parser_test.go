package tests

import (
	"github.com/nekovalue/unpckr/internal/argparser"
	"github.com/nekovalue/unpckr/internal/config"
	"os"
	"testing"
)

func TestParseArguments(t *testing.T) {
	// Prepare
	err := os.Mkdir("temp", 0755)
	err = os.Mkdir("DemonLordDante", 0755)
	err = os.Mkdir("gos", 0755)

	// Arrange
	testTable := []struct {
		fakeArgs       []string
		expectedSource []string
		fakeConfig     config.ConfigurationType
	}{
		{
			fakeArgs: []string{
				"unpckr.exe",
				"-s",
				"DemonLordDante"},
			expectedSource: []string{"DemonLordDante"},
			fakeConfig:     config.ConfigurationType{},
		},
		{
			fakeArgs: []string{
				"unpckr.exe",
				"-s",
				"DemonLordDante",
				"-z",
				"-d",
				"temp"},
			expectedSource: []string{"DemonLordDante"},
			fakeConfig:     config.ConfigurationType{},
		},
		{
			fakeArgs: []string{
				"unpckr.exe",
				"-s",
				"DemonLordDante",
				"-s",
				"gos",
				"-z",
				"-d",
				"temp"},
			expectedSource: []string{
				"DemonLordDante",
				"gos"},
			fakeConfig: config.ConfigurationType{},
		},
	}

	// Act
	for _, testCase := range testTable {
		_, err := argparser.ParseArguments(testCase.fakeArgs, &testCase.fakeConfig)

		t.Logf("Calling ParseArguments(%v, %v), inner error: %v",
			testCase.fakeArgs, "StandartFakeConfig", err)

		// Assert
		if err != nil && testCase.fakeConfig.Sources != &testCase.expectedSource {
			t.Errorf("Invalid argument source string. Expected %v, got %v",
				testCase.expectedSource, testCase.fakeConfig.Sources)
		}
	}

	// Cleaning

	err = os.Remove("temp")
	err = os.Remove("DemonLordDante")
	err = os.Remove("gos")
	if err != nil {
		t.Errorf("Directory remove error %v", err.Error())
	}
}
