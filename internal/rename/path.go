package rename

import (
	"os"
	"strings"
)

func getPathWithoutFileName(path string) string {
	resultParts := strings.Split(path, string(os.PathSeparator))
	result := strings.Join(resultParts[:len(resultParts)-1], string(os.PathSeparator)) + string(os.PathSeparator)

	return result
}
