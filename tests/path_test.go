package tests

import (
	"github.com/nekovalue/unpckr/pkg/path"
	"os"
	"testing"
)

func TestGetPathWithoutFileName(t *testing.T) {
	// Arrange
	testTable := []struct {
		fakePath     string
		expectedPath string
	}{
		{
			fakePath: string(os.PathSeparator) + "folder" + string(os.PathSeparator) + "sub_folder" +
				string(os.PathSeparator) + "filename.jpg",
			expectedPath: string(os.PathSeparator) + "folder" +
				string(os.PathSeparator) + "sub_folder" + string(os.PathSeparator),
		},
	}

	// Act
	for _, testCase := range testTable {
		value := path.GetPathWithoutFileName(testCase.fakePath)

		t.Logf("Calling GetPathWithoutFileName(%v)", testCase.fakePath)

		// Assert
		if value != testCase.expectedPath {
			t.Errorf("Uncorrect out path. Expected %v, got %v", testCase.expectedPath, value)
		}
	}
}

func TestGetFileNameWithoutExtensionFromPath(t *testing.T) {
	// Arrange
	testTable := []struct {
		fakePath         string
		expectedFileName string
	}{
		{
			fakePath: string(os.PathSeparator) + "folder" + string(os.PathSeparator) + "sub_folder" +
				string(os.PathSeparator) + "filename.jpg",
			expectedFileName: "filename",
		},
	}

	// Act
	for _, testCase := range testTable {
		value := path.GetFileNameWithoutExtensionFromPath(testCase.fakePath)

		t.Logf("Calling GetFileNameWithoutExtensionFromPath(%v)", testCase.fakePath)

		// Assert
		if value != testCase.expectedFileName {
			t.Errorf("Uncorrect out filename. Expected %v, got %v", testCase.expectedFileName, value)
		}
	}
}
