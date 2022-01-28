package tests

import (
	"github.com/nekovalue/unpckr/pkg/tools"
	"testing"
)

func TestRemoveSliceElementByIndex(t *testing.T) {
	// Arrange
	testTable := []struct {
		fakeSlice     []string
		indexToRemove int
		expectedSlice []string
	}{
		{
			fakeSlice:     []string{"first", "second", "third"},
			indexToRemove: 0,
			expectedSlice: []string{"second", "third"},
		},
		{
			fakeSlice:     []string{"foxy", "kitten", "george", "summary"},
			indexToRemove: 1,
			expectedSlice: []string{"foxy", "george", "summary"},
		},
	}

	// Act
	for _, testCase := range testTable {
		value := tools.RemoveSliceElementByIndex(testCase.fakeSlice, testCase.indexToRemove)

		t.Logf("Calling RemoveSliceElementByIndex(%v, %v)", testCase.fakeSlice, testCase.indexToRemove)

		// Assert
		if !tools.EqualSlice(value, testCase.expectedSlice) {
			t.Errorf("Uncorrect slice. Expected %v, got %v", testCase.expectedSlice, value)
		}
	}
}
