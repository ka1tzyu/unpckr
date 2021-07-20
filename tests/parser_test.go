package tests

import (
	"github.com/nekovalue/unpckr/internal/argparser"
	"github.com/nekovalue/unpckr/internal/config"
	"os"
	"testing"
)

func TestParseArguments(t *testing.T) {
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
				"D:\\dev\\go\\go-data\\unpckr-data\\DemonLordDante"},
			expectedSource: []string{"D:\\dev\\go\\go-data\\unpckr-data\\DemonLordDante"},
			fakeConfig:     config.ConfigurationType{},
		},
		{
			fakeArgs: []string{
				"unpckr.exe",
				"-s",
				"D:\\dev\\go\\go-data\\unpckr-data\\DemonLordDante",
				"-z",
				"-d",
				"temp"},
			expectedSource: []string{"D:\\dev\\go\\go-data\\unpckr-data\\DemonLordDante"},
			fakeConfig:     config.ConfigurationType{},
		},
		{
			fakeArgs: []string{
				"unpckr.exe",
				"-s",
				"D:\\dev\\go\\go-data\\unpckr-data\\DemonLordDante",
				"-s",
				"D:\\dev\\go",
				"-z",
				"-d",
				"temp"},
			expectedSource: []string{
				"D:\\dev\\go\\go-data\\unpckr-data\\DemonLordDante",
				"D:\\dev\\go"},
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

	err := os.Remove("temp")
	if err != nil {
		t.Errorf("Directory remove error %v", err.Error())
	}
}
