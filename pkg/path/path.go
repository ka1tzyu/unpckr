package path

import (
	"os"
	"strings"
)

// GetPathWithoutFileName returns folder of given path
func GetPathWithoutFileName(path string) string {
	resultParts := strings.Split(path, string(os.PathSeparator))
	result := strings.Join(resultParts[:len(resultParts)-1], string(os.PathSeparator)) + string(os.PathSeparator)

	return result
}

// GetFileNameWithoutExtensionFromPath returns file`s name without it`s extension
func GetFileNameWithoutExtensionFromPath(path string) string {
	resultParts := strings.Split(path, string(os.PathSeparator))
	resultParts = strings.Split(resultParts[len(resultParts)-1], ".")

	result := strings.Join(resultParts[:len(resultParts)-1], ".")

	return result
}
